package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
)

var (
	redisClusterOnce sync.Once
	client           *redis.ClusterClient
)

/**
*	获取redis cluster 实例
*	@param redisInstance redis类型 0 redis实例
 */
func redisClusterFactory(redisInstance int) (clusterClient *redis.ClusterClient) {
	switch redisInstance {
	case 0:
		clusterClient = getRedisClusterInstance()
	default:
		e := fmt.Sprintf("redisInstance err:%s the redisInstance %d", "请输入正确的redis类型", redisInstance)
		panic(e)
	}
	return clusterClient
}

func getRedisClusterInstance() (clusterClient *redis.ClusterClient) {
	redisClusterOnce.Do(func() {
		conf := &redis.ClusterOptions{
			Addrs:    []string{"ip1:port","ip2:port"},
			ReadOnly: true,
			PoolSize: 100,
		}
		client = redis.NewClusterClient(conf)
	})
	return client
}
