package commands

import (
	"bufio"
	"fmt"
	"os"
	"simplekv/models"
	"strconv"
	"strings"
	"time"
)

func Run() {
	input := bufio.NewScanner(os.Stdin)
	hello()
	help()
	for {
		follow()
		input.Scan()         //扫描输入内容
		line := input.Text() //把输入内容转换为字符串
		args := strings.Split(line, " ")
		if checkArgsBefore(args) {
			switch args[0] {
			case "SET":
				setCache(args)
			case "GET":
				getCache(args)
			case "DEL":
				delCache(args)
			case "EXISTS":
				queryCache(args)
			case "FLUSH":
				flushCache()
			case "KEYS":
				queryAllCache()
			case "SMM":
				setMaxMemory(args)
			default:
				fmt.Println("unknow opertaion")
				help()
			}
		}

	}
}

func setCache(args []string) {
	if checkArgsFour(args) {
		expire, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Println("expireTime is error, it must be number")
			return
		}
		models.Set(args[1], args[2], time.Duration(expire*1000*1000*1000))
	}
}

func getCache(args []string) {
	if checkArgsTwo(args) {
		models.Get(args[1])
	}
}

func delCache(args []string) {
	if checkArgsTwo(args) {
		models.Del(args[1])
	}
}

func queryCache(args []string) {
	if checkArgsTwo(args) {
		models.Exists(args[1])
	}
}

func flushCache() {
	models.Flush()
}

func queryAllCache() {
	models.Keys()
}

func setMaxMemory(args []string) {
	models.SetMaxMemory(args[1])
}

func hello() {
	fmt.Println("欢迎使用simpleKV,请按照提示操作: ")
}

func help() {
	fmt.Println("如想新增缓存请输入 SET key value expireTime，中间以空隔隔开")
	fmt.Println("如想获取缓存请输入 GET key，中间以空隔隔开")
	fmt.Println("如想删除缓存请输入 DEL key，中间以空隔隔开")
	fmt.Println("如想查询缓存是否存在请输入 EXISTS key，中间以空隔隔开")
	fmt.Println("如想清空缓存请输入 FLUSH")
	fmt.Println("如想查询缓存数量请输入 KEYS")
}

func follow() {
	fmt.Print("请输入: ")
}

func remind(requires, received int) {
	fmt.Printf("Error: requires at least %d args, only received %d\n", requires, received)
}

func checkArgsBefore(args []string) bool {
	if len(args) < 1 {
		help()
		follow()
		return false
	}
	return true
}

func checkArgsFour(args []string) bool {
	if len(args) < 4 {
		//fmt.Println("args: ", args)
		remind(4, len(args))
		help()
		follow()
		return false
	}
	return true
}

func checkArgsTwo(args []string) bool {
	if len(args) < 2 {
		remind(2, len(args))
		help()
		follow()
		return false
	}
	return true
}
