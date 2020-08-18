package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"testing"
)

var w = &sync.WaitGroup{}

func TestRedisClusterFactory(t *testing.T) {
	redisClusterFactory(0)
}
func TestSet(t *testing.T) {
	s, err := Set(0, "test1", "test1", 0)
	fmt.Println("rs:", s, err, ":end")
}

//func BenchmarkSet(b *testing.B) {
//	w.Add(b.N)
//	for i := 0; i < b.N; i++ {
//		//go Set(0, "test"+strconv.Itoa(i), 0, 60, w)
//	}
//	w.Wait()
//}
func TestGet(t *testing.T) {
	s, err := Get(0, "test1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}

func TestDel(t *testing.T) {
	s, err := Del(0, "test1")
	fmt.Println(s, err)
}
func TestHSet(t *testing.T) {
	b, err := HSet(0, "htest1", "name1", "test1")
	fmt.Println(b, err)
}

func TestHGet(t *testing.T) {
	rs, err := HGet(0, "htest", "name")
	fmt.Println(rs, err)
}
func TestHMSet(t *testing.T) {
	s, err := HMSet(0, "htest", map[string]interface{}{"t1": "v1", "t2": "v2"})
	fmt.Println(s, err)
}

func TestHMGet(t *testing.T) {
	fmt.Println(HMGet(0, "htest", []string{}))
}
func TestHGetAll(t *testing.T) {
	rs, _ := HGetAll(0, "htest1")
	fmt.Println(rs)
}

func TestHDel(t *testing.T) {
	rs, err := HDel(0, "htest", []string{"t1", "t3", "name"})
	fmt.Println(rs, err)
}

func TestZAdd(t *testing.T) {
	rs, err := ZAdd(0, "zaddtest", []float64{}, []interface{}{})
	fmt.Println(rs, err)
}

func TestZRem(t *testing.T) {
	rs, err := ZRem(0, "zaddtest", []interface{}{"aa", "bb"})
	fmt.Println(rs, err)
}
func TestZCard(t *testing.T) {
	rs, err := ZCard(0, "zaddtest")
	fmt.Println(rs, err)
}
func TestZRange(t *testing.T) {
	rs, err := ZRange(0, "zaddtest", 0, -1)
	fmt.Println(rs, err)
}

func TestZRangeWithScores(t *testing.T) {
	rs, err := ZRangeWithScores(0, "zaddtest", 0, -1)
	fmt.Println(rs, err)
}

func TestSAdd(t *testing.T) {
	rs, err := SAdd(0, "saddtest", []interface{}{"cc"})
	fmt.Println(rs, err)
}

func TestSRem(t *testing.T) {
	rs, err := SRem(0, "saddtest", []interface{}{""})
	fmt.Println(rs, err)
}

func TestSMembers(t *testing.T) {
	rs, err := SMembers(0, "saddtest")
	fmt.Println(rs, err)
}
func TestExpire(t *testing.T) {
	//Set(0, "test1", "test", 0)
	//fmt.Println(Get(0, "test1"))
	//fmt.Println(TTL(0, "test1"))
	//fmt.Println(Expire(0, "test1", 100))
	fmt.Println(TTL(0, "test1"))
}

func test(pipe redis.Pipeliner) error {
	pipe.HGetAll("htest").Result()
	pipe.HGetAll("htest1").Result()
	return nil
}
func TestPipelined(t *testing.T) {
	cmdEr, _ := Pipelined(0, test)
	for _, cmder := range cmdEr {
		cmd := cmder.(*redis.StringStringMapCmd)
		strMap, _ := cmd.Result()

		fmt.Println("strMap", strMap["name"])
	}
}
