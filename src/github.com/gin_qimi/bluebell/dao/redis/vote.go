package redis

import (
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
	scorePerVote      = 432.0 // 每一票的分数
)

func CreatePost(postID, communityID int64) error {
	// // 帖子时间
	// _, err := client.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{ // ZSet中的分数是时间戳
	// 	Score:  float64(time.Now().Unix()),
	// 	Member: postID,
	// }).Result()
	// if err != nil {
	// 	return err
	// }
	// // 帖子分数
	// _, err = client.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
	// 	Score:  float64(time.Now().Unix()),
	// 	Member: postID,
	// }).Result()
	pipeline := client.TxPipeline() // 事务
	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{ // ZSet中的分数是时间戳
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 把帖子id加入到社区的set
	// 更新: 把帖子id加到社区的set 
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))
	pipeline.SAdd(cKey, postID)
	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userID, postID string, value float64) (err error) {
	// 1. 判断投票限制
	// 去redis取帖子发布时间
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	zap.L().Debug("VoteForPost", zap.Float64("postTime", postTime))
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 2. 更新帖子的分数
	// 先查当前用户给当前帖子的投票记录
	ov := client.ZScore(getRedisKey(KeyPostVotedZSetPF+postID), userID).Val() // 之前的投票分数
	if value == ov {
		return ErrVoteRepeated
	}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value)    // 计算两次投票的差值
	pipeline := client.TxPipeline() // 事务
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)
	// 3. 记录用户为该帖子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPF+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  value, // 赞成票还是反对票
			Member: userID,
		})
	}
	_, err = pipeline.Exec()
	// _, err = client.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID).Result()
	// if err != nil {
	// 	return err
	// }
	// // 3. 记录用户为该帖子投票的数据
	// if value == 0 {
	// 	_, err = client.ZRem(getRedisKey(KeyPostVotedZSetPF+postID), userID).Result()
	// } else {
	// 	_, err = client.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
	// 		Score:  value, // 赞成票还是反对票
	// 		Member: userID,
	// 	}).Result()
	// }
	return err
}
