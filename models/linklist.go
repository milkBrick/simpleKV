package models

import (
	"fmt"
	"reflect"
	"time"
)

type node struct {
	key        string
	value      interface{}
	insertTime time.Time
	expireTime time.Duration
	next       *node
}

func newNode(k string, v interface{}, expire time.Duration) *node {
	return &node{
		key:        k,
		value:      v,
		insertTime: time.Now(),
		expireTime: expire,
	}
}

type linkList struct {
	head   *node
	length uint
}

func newlinkList() *linkList {
	return &linkList{
		head:   newNode("", nil, 0),
		length: 0,
	}
}

func (l *linkList) isEmpty() bool {
	return l.length == 0
}

//在某个节点后面插入
func (l *linkList) insertAfter(p *node, k string, v interface{}, expire time.Duration) bool {
	if p == nil {
		return false
	}
	newNode := newNode(k, v, expire)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	l.length++
	return true
}

//在链表尾部插入节点
func (l *linkList) insertToTail(k string, v interface{}, expire time.Duration) bool {
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	return l.insertAfter(cur, k, v, expire)
}

//获取某个节点
func (l *linkList) get(k string) *node {
	cur := l.head
	for cur != nil {
		if cur.key == k {
			return cur
		}
		cur = cur.next
	}
	return nil
}

//获取某个节点的值
func (l *linkList) getValue(k string) interface{} {
	node := l.get(k)
	if node != nil {
		return node.value
	}
	return nil
}

//判断链表中是否存在这个节点
func (l *linkList) isExist(k string) bool {
	return l.get(k) != nil
}

//删除某个节点
func (l *linkList) deleteNode(k string) bool {
	cur := l.head.next
	pre := l.head
	for cur != nil {
		if cur.key == k {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	pre.next = cur.next
	cur = nil
	l.length--
	return true
}

//清除所有节点
func (l *linkList) cleanAll() bool {
	cur := l.head
	for cur != nil {
		l.deleteNode(cur.key)
		cur = cur.next
	}
	return true
}

//获取过期的缓存
func (l *linkList) getExpireKeys() (keys []string) {
	cur := l.head
	for cur != nil {
		if time.Now().Sub(cur.insertTime) >= cur.expireTime {
			keys = append(keys, cur.key)
		}
		cur = cur.next
	}
	return keys
}

//打印链表
func (l *linkList) print() {
	cur := l.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("{{%v}:{%+v}}", cur.key, cur.value)
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}

func (l linkList) getLong() int {
	return int(reflect.TypeOf(l).Size())
}

func (l linkList) getFieldMaxSize() int {
	var maxSize int
	itype := reflect.TypeOf(l)
	for i := 0; i < itype.NumField(); i++ {
		field := itype.Field(i)
		if int(field.Type.Size()) > maxSize {
			maxSize = int(field.Type.Size())
		}
	}
	return maxSize
}

func (l linkList) getStructAlign() int {
	var align int
	itype := reflect.TypeOf(l)
	for i := 0; i < itype.NumField(); i++ {
		field := itype.Field(i)
		if field.Type.Align() > align {
			align = field.Type.Align()
		}
	}
	return align
}

func (l linkList) getStructSize(long, size, align int) (cap int) {
	var c int
	if size > align {
		c = align
	} else {
		c = size
	}
	remainder := long % c
	if remainder == 0 {
		cap = long
	} else {
		cap = c * (long/c + 1)
	}
	return cap
}

func (n node) getLong() int {
	return int(reflect.TypeOf(n).Size())
}

func (n node) getFieldMaxSize() int {
	var maxSize int
	itype := reflect.TypeOf(n)
	for i := 0; i < itype.NumField(); i++ {
		field := itype.Field(i)
		if int(field.Type.Size()) > maxSize {
			maxSize = int(field.Type.Size())
		}
	}
	return maxSize
}

func (n node) getStructAlign() int {
	var align int
	itype := reflect.TypeOf(n)
	for i := 0; i < itype.NumField(); i++ {
		field := itype.Field(i)
		if field.Type.Align() > align {
			align = field.Type.Align()
		}
	}
	return align
}

func (n node) getStructSize(long, size, align int) (cap int) {
	var c int
	if size > align {
		c = align
	} else {
		c = size
	}
	remainder := long % c
	if remainder == 0 {
		cap = long
	} else {
		cap = c * (long/c + 1)
	}
	return cap
}
