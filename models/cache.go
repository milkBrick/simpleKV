package models

import "time"

type Cache interface {
	//size是一个字符串，支持以下参数： 1kb 100kb 1mb 2mb 1gb等
	SetMaxMemory(size string) bool
	//设置一个缓存项，并且在exprie时间之后过期
	Set(key string, val interface{}, expire time.Duration)
	//获取一个值
	Get(key string) (interface{}, bool)
	//删除一个值
	Del(key string) bool
	//检测一个值，是否存在
	Exists(key string) bool
	//清除所有值
	Flush() bool
	//返回所有的key多少
	Keys() int64
	//清除过期缓存
	ClearExpireNode()
}
