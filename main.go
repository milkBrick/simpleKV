package main

import (
	"simplekv/commands"
	"simplekv/models"
)

func init() {
	models.InitCache()
}

func main() {
	//for i := 0; i < 20; i++ {
	//	models.Set("hello"+fmt.Sprintf("%d", i), "world"+fmt.Sprintf("%d", i), time.Duration(rand.Int63n(10000000000)))
	//}
	//
	//fmt.Printf("get cache size: %d\n", models.Keys())
	////查询name100 key的值
	//val, ok := models.Get("hello10")
	//fmt.Printf("get hello100 value:%v, status:%v\n", val, ok)
	////删除name100 key
	//models.Del("hello10")
	////查询name100 key是否删除成功
	//fmt.Printf("get hello10 key:%v\n", models.Exists("hello10"))

	commands.Run()
	select {}
}
