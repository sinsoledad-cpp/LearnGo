package logger

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin_qimi/bluebell/setting"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Init 初始化Logger
func Init(cfg *setting.LogConfig, mode string) (err error) {
	writeSyncer := getLogWriter(
		cfg.Filename,
		cfg.MaxSize,
		cfg.MaxBackups,
		cfg.MaxAge,
	)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	var core zapcore.Core
	if mode == "dev" {
		// 进入开发模式，日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}

	lg := zap.New(core, zap.AddCaller())
	// 替换zap库中全局的logger
	zap.ReplaceGlobals(lg)
	return
}

func Close() {
	err := zap.L().Sync()
	if err != nil {
		zap.L().Error("zap close failed", zap.Error(err))
		fmt.Println("zap close failed", zap.Error(err))
	}
}

// func Init() (err error) {
// 	writerSyncer := getLogWriter( // 获取到文件句柄
// 		viper.GetString("log.filename"), // 日志文件名
// 		viper.GetInt("log.max_size"),    // 每个日志文件保存的最大尺寸
// 		viper.GetInt("log.max_backups"), // 日志文件最多保存多少个备份
// 		viper.GetInt("log.max_age"),     // 日志文件最多保存多少天
// 	)
// 	encoder := getEncoder()    // 获取编码器
// 	var l = new(zapcore.Level) // 获取日志级别
// 	err = l.UnmarshalText([]byte(viper.GetString("log.level")))
// 	if err != nil {
// 		return
// 	}
// 	core := zapcore.NewCore(encoder, writerSyncer, l) // 获取日志的core
// 	logger := zap.New(core, zap.AddCaller()) // 获取日志,zap.AddCaller() 添加文件和行号
// 	zap.ReplaceGlobals(logger) // 替换zap库中全局的logger
// 	return nil
// }

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()             //生产环境
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder         //时间格式
	encoderConfig.TimeKey = "time"                                //时间key
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder       //日志级别
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder //日志持续时间
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder       //日志文件
	return zapcore.NewJSONEncoder(encoderConfig)                  //生产环境
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{ //日志切割
		Filename:   filename,  //日志文件名
		MaxSize:    maxSize,   //每个日志文件保存的最大尺寸
		MaxBackups: maxBackup, //日志文件最多保存多少个备份
		MaxAge:     maxAge,    //日志文件最多保存多少天
	}
	return zapcore.AddSync(lumberJackLogger) //日志切割
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
