package models

import (
	"fmt"
	"time"
)

var cache Cache

func InitCache() {
	cache = NewCache("256mb", 5)
	go cache.ClearExpireNode()
	fmt.Printf("start init cache cap: %v, expireTime: %d\n", "256mb", 5)
}

func SetMaxMemory(size string) bool {
	b := cache.SetMaxMemory(size)
	if b {
		fmt.Printf("cache setmaxmemory %v successfully\n", size)
	} else {
		fmt.Printf("cache setmaxmemory %v failed\n", size)
	}
	return b
}

func Set(key string, val interface{}, expire time.Duration) {
	cache.Set(key, val, expire)
	fmt.Printf("cache set {{key:%v},{value:%v},{expire:%v}}\n", key, val, expire)
}

func Get(key string) (interface{}, bool) {
	v, b := cache.Get(key)
	fmt.Printf("cache get key:%v, status:%v, value:%v\n", key, b, v)
	return cache.Get(key)
}

func Del(key string) bool {
	b := cache.Del(key)
	fmt.Printf("cache delete key: %v, status: %v\n", key, b)
	return b
}

func Exists(key string) bool {
	exists := cache.Exists(key)
	if exists {
		fmt.Printf("cache exist key: %v\n", key)
	} else {
		fmt.Printf("cache do not exist key: %v\n", key)
	}
	return exists
}

func Flush() bool {
	flush := cache.Flush()
	if flush {
		fmt.Println("cache flush successfully.")
	} else {
		fmt.Println("cache flush failed.")
	}
	return flush
}

func Keys() int64 {
	keys := cache.Keys()
	fmt.Printf("cache total exists %v\n", keys)
	return keys
}
