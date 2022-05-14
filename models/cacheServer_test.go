package models

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestInitCache(t *testing.T) {
	InitCache()
}

func TestSetMaxMemory(t *testing.T) {
	InitCache()
	SetMaxMemory("1GB")
}

func TestSet(t *testing.T) {
	InitCache()
	for i := 0; i < 10; i++ {
		Set("hello"+fmt.Sprintf("%d", i+1), "world"+fmt.Sprintf("%d", i+1), time.Duration(rand.Int63n(10000000000)))
	}
	Keys()
}

func TestGet(t *testing.T) {
	InitCache()
	fmt.Println(cache.Keys())
	for i := 0; i < 10; i++ {
		Set("hello"+fmt.Sprintf("%d", i+1), "world"+fmt.Sprintf("%d", i+1), time.Duration(rand.Int63n(10000000000)))
	}
	for i := 0; i < 12; i++ {
		Get("hello" + fmt.Sprintf("%d", i))
	}
}

func TestDel(t *testing.T) {
	InitCache()
	for i := 0; i < 10; i++ {
		Set("hello"+fmt.Sprintf("%d", i+1), "world"+fmt.Sprintf("%d", i+1), time.Duration(rand.Int63n(10000000000)))
	}
	Exists("hello1")
	Del("hello1")
	Exists("hello1")
}

func TestFlush(t *testing.T) {
	InitCache()
	for i := 0; i < 10; i++ {
		Set("hello"+fmt.Sprintf("%d", i+1), "world"+fmt.Sprintf("%d", i+1), time.Duration(rand.Int63n(10000000000)))
	}
	for i := 0; i < 12; i++ {
		Exists("hello" + fmt.Sprintf("%d", i))
	}
	Flush()
	for i := 0; i < 12; i++ {
		Exists("hello" + fmt.Sprintf("%d", i))
	}
}

func TestKeys(t *testing.T) {
	InitCache()
	Keys()
	for i := 0; i < 10; i++ {
		Set("hello"+fmt.Sprintf("%d", i+1), "world"+fmt.Sprintf("%d", i+1), time.Duration(rand.Int63n(10000000000)))
	}
	Keys()
}
