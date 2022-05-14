package models

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSimpleKV_Set(t *testing.T) {
	cap := "256MB"
	expireTime := time.Duration(5)
	cache := NewCache(cap, expireTime)
	for i := 0; i < 10; i++ {
		cache.Set("hello"+fmt.Sprintf("%d", i+1), "world"+fmt.Sprintf("%d", i+1), time.Duration(rand.Int63n(10000000000)))
	}
	fmt.Println(cache.Keys())
}

func TestSimpleKV_Get(t *testing.T) {
	cap := "256MB"
	expireTime := time.Duration(5)
	cache := NewCache(cap, expireTime)
	for i := 0; i < 10; i++ {
		cache.Set("hello"+fmt.Sprintf("%d", i+1), "world"+fmt.Sprintf("%d", i+1), time.Duration(rand.Int63n(10000000000)))
	}
	for i := 0; i < 12; i++ {
		fmt.Println(cache.Get("hello" + fmt.Sprintf("%d", i)))
	}
}

func TestSimpleKV_Exists(t *testing.T) {
	cap := "256MB"
	expireTime := time.Duration(5)
	cache := NewCache(cap, expireTime)
	for i := 0; i < 10; i++ {
		cache.Set("hello"+fmt.Sprintf("%d", i+1), "world"+fmt.Sprintf("%d", i+1), time.Duration(rand.Int63n(10000000000)))
	}
	for i := 0; i < 12; i++ {
		fmt.Println(cache.Exists("hello" + fmt.Sprintf("%d", i)))
	}
}

func TestSimpleKV_Del(t *testing.T) {
	cap := "256MB"
	expireTime := time.Duration(5)
	cache := NewCache(cap, expireTime)
	for i := 0; i < 10; i++ {
		cache.Set("hello"+fmt.Sprintf("%d", i+1), "world"+fmt.Sprintf("%d", i+1), time.Duration(rand.Int63n(10000000000)))
	}
	fmt.Println(cache.Exists("hello1"))
	fmt.Println(cache.Del("hello1"))
	fmt.Println(cache.Exists("hello1"))
}

func TestSimpleKV_Flush(t *testing.T) {
	cap := "256MB"
	expireTime := time.Duration(5)
	cache := NewCache(cap, expireTime)
	for i := 0; i < 10; i++ {
		cache.Set("hello"+fmt.Sprintf("%d", i+1), "world"+fmt.Sprintf("%d", i+1), time.Duration(rand.Int63n(10000000000)))
	}
	for i := 0; i < 12; i++ {
		fmt.Println(cache.Exists("hello" + fmt.Sprintf("%d", i)))
	}
	fmt.Println("cache.Flush(): ", cache.Flush())
	for i := 0; i < 12; i++ {
		fmt.Println(cache.Exists("hello" + fmt.Sprintf("%d", i)))
	}
}

func TestSimpleKV_Keys(t *testing.T) {
	cap := "256MB"
	expireTime := time.Duration(5)
	cache := NewCache(cap, expireTime)
	for i := 0; i < 10; i++ {
		cache.Set("hello"+fmt.Sprintf("%d", i+1), "world"+fmt.Sprintf("%d", i+1), time.Duration(rand.Int63n(10000000000)))
	}
	fmt.Println(cache.Keys())
	fmt.Println("cache.Flush(): ", cache.Flush())
	fmt.Println(cache.Keys())
}
