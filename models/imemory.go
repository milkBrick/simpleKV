package models

type imemory interface {
	//获取结构体占用的内存大小
	getLong() int
	//获取结构体内的字段最大类型长度
	getFieldMaxSize() int
	//获取结构体内字段的最大对齐数值
	getStructAlign() int
	////获取结构体内存对齐之后所占用的内存大小
	getStructSize(long, size, align int) (cap int)
}
