package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Parse(cap string) (int, error) {
	if len(cap) < 3 {
		return 0, errors.New("cap is error")
	}
	upper := strings.ToUpper(cap)
	size := upper[0 : len(upper)-2]
	level := upper[len(upper)-2:]
	fmt.Println("size, level: ", size, level)
	atoi, err := strconv.Atoi(size)
	if err != nil {
		return 0, errors.New("size is error")
	}

	switch level {
	case "KB":
		atoi = atoi * 1024
	case "MB":
		atoi = atoi * 1024 * 1024
	case "GB":
		atoi = atoi * 1024 * 1024 * 1024
	default:
		return 0, errors.New("level is error")
	}
	return atoi, nil
}

//func GetLong(i interface{}) int {
//	return int(reflect.TypeOf(i).Size())
//}
//
//func GetFieldMaxSize(i interface{}) int {
//	var maxSize int
//	itype := reflect.TypeOf(i)
//	for i := 0; i < itype.NumField(); i++ {
//		field := itype.Field(i)
//		if int(field.Type.Size()) > maxSize {
//			maxSize = int(field.Type.Size())
//		}
//	}
//	return maxSize
//}
//
//func GetStructAlign(i interface{}) int {
//	var align int
//	itype := reflect.TypeOf(i)
//	for i := 0; i < itype.NumField(); i++ {
//		field := itype.Field(i)
//		if field.Type.Align() > align {
//			align = field.Type.Align()
//		}
//	}
//	return align
//}
//
//func GetStructSize(long, size, align int) (cap int) {
//	var c int
//	if size > align {
//		c = align
//	} else {
//		c = size
//	}
//	remainder := long % c
//	if remainder == 0 {
//		cap = long
//	} else {
//		cap = c * (long/c + 1)
//	}
//	return cap
//}
