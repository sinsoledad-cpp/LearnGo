# Hertz

## 介绍

![介绍](image/介绍.png)

## 中间件

![全局中间件](image/全局中间件.png)

![组中间件](image/组中间件.png)

![client中间件](image/client中间件.png)

## Hertz拓展中间件

![自定义中间件](image/自定义中间件.png)

![next](image/next.png)

![abort](image/abort.png)


# BUG

## consul注册问题


在配置文件中的address不可写为127.0.0.1,否则会出现注册失败(使用docker运行consul)
```yaml
kitex:
  service: "demo_proto"
  address: "192.168.0.104:8080"
```

## 注意类型断言的使用.(类型)

如果类型断言出错,程序会停止