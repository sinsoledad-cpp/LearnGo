package redis

// redis key

// redis key注意使用命名空间的方式,方便查询和拆分

const (
	KeyPrefix          = "bluebell:"
	KeyPostTimeZSet    = "post:time"   // zset;贴子及发帖时间
	KeyPostScoreZSet   = "post:score"  // zset;贴子及投票的分数
	KeyPostVotedZSetPF = "post:voted:" // zset;记录用户及投票类型;参数是post id
	KeyCommunitySetPF  = "community:"  // set;保存每个分区下帖子的id
)

// getRedisKey 获取redis key,需要拼接前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
