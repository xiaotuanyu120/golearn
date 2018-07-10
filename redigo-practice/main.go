package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	/*
		参照文档：
		https://godoc.org/github.com/gomodule/redigo/redis
	*/
	c, err := redis.Dial("tcp", "192.168.33.10:6379", redis.DialPassword("123456"))
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// prepare test data
	c.Send("MULTI")
	c.Send("FLUSHALL")
	for i := 1; i < 100; i++ {
		c.Send("SET", i, "testForSession")
		c.Send("EXPIRE", i, "60")
	}
	if _, err := c.Do("EXEC"); err != nil {
		panic(err)
	}

	// scan all keys
	iter := 0
	var keys []string
	for {
		reply, err := redis.Values(c.Do("SCAN", iter))
		if err != nil {
			panic(err)
		}

		/*
			目的：要获取[iterator, [key1, key2, ...]]
			要点：直接reply[0]获取不到iterator，但是通过redis.Int转换一下reply[0]就可以，同理需要redis.String转换reply[1]

			参照文档：
			- https://godoc.org/github.com/gomodule/redigo/redis#hdr-Reply_Helpers
			- https://stackoverflow.com/questions/31498296/how-to-scan-keys-from-redis-using-golang-using-scan-not-keys
		*/
		iter, _ = redis.Int(reply[0], nil)
		keysIter, _ := redis.Strings(reply[1], nil)

		/*
			slice append by item：
			```
			var slice1, slice2 []string
			slice1 = []string{"1", "2"}
			slice2 = []string{"3", "4"}
			fmt.Println(append(slice1, "1", "2"))
			```
			结果是[1 2 1 2]

			slice extend another slice:
			```
			var slice1, slice2 []string
			slice1 = []string{"1", "2"}
			slice2 = []string{"3", "4"}
			fmt.Println(append(slice1, slice2...))
			```
			结果是[1 2 3 4]
		*/
		keys = append(keys, keysIter...)

		if iter == 0 {
			break
		}
	}
	fmt.Println(keys)
}
