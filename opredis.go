package main


import (
	"errors"
	"github.com/go-redis/redis"
	"time"
)

/**
*	设置字符串值
*	@param: redisInstance redis集群实例
*	@param: key 键
*	@param: data 值
*	@param: expiration 过期时间 单位秒
 */
func Set(redisInstance int, k string, v interface{}, expiration int64) (rs string, err error) {
	return redisClusterFactory(redisInstance).Set(k, v, time.Duration(expiration)*time.Second).Result()
}

/**
*	获取字符串值
*	@param: redisInstance redis实例
*	@param: key 键
 */
func Get(redisInstance int, k string) (rs string, err error) {
	return redisClusterFactory(redisInstance).Get(k).Result()
}

/**
*	删除指定键的值
*	@param: redisInstance redis实例
*	@param: key 键
*	@return :v 1 成功 0 没有该键
 */
func Del(redisInstance int, k string) (rs int64, err error) {
	return redisClusterFactory(redisInstance).Del(k).Result()
}

/**
*	hash设置 键值
*	@param: k key
*	@param: field字段
*	@param: v 值
 */
func HSet(redisInstance int, k, field string, v interface{}) (rs bool, err error) {
	return redisClusterFactory(redisInstance).HSet(k, field, v).Result()
}

/**
*	hash获取值
*	@param:
*	@param:k key
 */
func HGet(redisInstance int, k, field string) (rs string, err error) {
	return redisClusterFactory(redisInstance).HGet(k, field).Result()
}

/**
*	hash 设置多个值
*	@param:redisInstance redis实例
*	@param: k 键
*	@param: v hash键值对
 */
func HMSet(redisInstance int, k string, v map[string]interface{}) (rs string, err error) {
	return redisClusterFactory(redisInstance).HMSet(k, v).Result()
}

/**
*	获取hash多个值
*	@param: redisInstance redis实例
*	@param: k 键
*	@param: field 要获取的字段
 */
func HMGet(redisInstance int, k string, fields []string) (rs []interface{}, err error) {
	return redisClusterFactory(redisInstance).HMGet(k, fields...).Result()
}

/**
*	获取hash所有值
*	@param: redisInstance redis实例
*	@param: k 键
 */
func HGetAll(redisInstance int, k string) (rs map[string]string, err error) {
	return redisClusterFactory(redisInstance).HGetAll(k).Result()
}

/**
*	删除hash指定field,支持删除多个field
*	@param: redisInstance redis实例
*	@param: k 键
 */
func HDel(redisInstance int, k string, fields []string) (rs int64, err error) {
	return redisClusterFactory(redisInstance).HDel(k, fields...).Result()
}

/**
*	有序集合中添加元素
*	@param: redisInstance redis实例
*	@param: k 键
*	@param: scores[]float64 分数
*	@param: members[]interface 元素
 */
func ZAdd(redisInstance int, k string, scores []float64, members []interface{}) (rs int64, err error) {
	if len(scores) != len(members) {
		return 0, errors.New("zadd scores != members")
	}
	zs := make([]redis.Z, 0)
	for i, v := range scores {
		zs = append(zs, redis.Z{
			Score:  v,
			Member: members[i],
		})
	}
	return redisClusterFactory(redisInstance).ZAdd(k, zs...).Result()
}

/**
*	删除有序集合中的元素
*	@param: redisInstance redis实例
*	@param: k 键
*	@param: members 要删除的元素
 */
func ZRem(redisInstance int, k string, members []interface{}) (rs int64, err error) {
	return redisClusterFactory(redisInstance).ZRem(k, members...).Result()
}

/**
*	计算有序集合中元素个数
*	@param: redisInstance redis实例
*	@param: k 键
 */
func ZCard(redisInstance int, k string) (rs int64, err error) {
	return redisClusterFactory(redisInstance).ZCard(k).Result()
}

/**
*	遍历有序集合
*	@param: redisInstance redis实例
*	@param: k 键
*	@param: s 开始位置
*	@param: s 结束位置
 */
func ZRange(redisInstance int, k string, s, e int64) (rs []string, err error) {
	return redisClusterFactory(redisInstance).ZRange(k, s, e).Result()
}

/**
*	按照分数遍历有序集合
*	@param: redisInstance
*	@param: k 键
*	@param: s开始位置
*	@param: e结束位置
*	@return [值 分数 值 分数....]
 */
func ZRangeWithScores(redisInstance int, k string, s, e int64) (rs []interface{}, err error) {
	z, err := redisClusterFactory(redisInstance).ZRangeWithScores(k, s, e).Result()
	if err != nil {
		return rs, err
	}
	for _, v := range z {
		rs = append(rs, v.Member, v.Score)
	}
	return rs, nil
}

/**
*	添加无序集合
*	@param: redisInstance redis实例
*	@param: k 键
*	@param: members 添加的元素
*	@return: rs 添加成功的元素个数 err 错误信息
 */
func SAdd(redisInstance int, k string, members []interface{}) (rs int64, err error) {
	return redisClusterFactory(redisInstance).SAdd(k, members...).Result()
}

/**
*	删除无序集合元素
*	@param: redisInstance redis实例
*	@param: k 键
*	@param: members 删除的元素
*	@return: rs 删除成功的元素个数 err 错误信息
 */
func SRem(redisInstance int, k string, members []interface{}) (rs int64, err error) {
	return redisClusterFactory(redisInstance).SRem(k, members...).Result()
}

/**
*	获取无序集合中的元素
*	@param: redisInstance redis实例
*	@param: k 键
*	@return: rs 返回的字符串slice,err 错误信息
 */
func SMembers(redisInstance int, k string) (rs []string, err error) {
	return redisClusterFactory(redisInstance).SMembers(k).Result()
}

/**
*	设置过期时间
*	@param: redisInstance redis实例
*	@param: k 键
*	@param: expiration 过期时间 单位秒
*	@return: rs bool 设置成功失败 err 错误信息
 */
func Expire(redisInstance int, k string, expiration int) (rs bool, err error) {
	return redisClusterFactory(redisInstance).Expire(k, time.Duration(expiration)*time.Second).Result()
}

/**
*	查看过期时间
*	@param: redisInstance redis实例
*	@param: k 键
*	@return: rs 过期时间 err 错误信息
 */
func TTL(redisInstance int, k string) (time.Duration, error) {
	return redisClusterFactory(redisInstance).TTL(k).Result()
}

func Pipelined(redisInstance int, f func(pipe redis.Pipeliner) error) (cmdEr []redis.Cmder, err error) {
	return redisClusterFactory(redisInstance).Pipelined(f)
}
