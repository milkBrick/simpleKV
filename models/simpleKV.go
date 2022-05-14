package models

import (
	"fmt"
	"reflect"
	"simplekv/utils"
	"sync"
	"time"
)

type simpleKV struct {
	cap        int
	tateup     int
	l          *linkList
	rwMutex    sync.RWMutex
	expireTime time.Duration
}

func (s *simpleKV) SetMaxMemory(size string) bool {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	cap, err := utils.Parse(size)
	if err != nil {
		return false
	}
	s.cap = cap
	return true
}

func (s *simpleKV) Set(key string, val interface{}, expire time.Duration) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	if s.l.isExist(key) {
		if !s.l.deleteNode(key) {
			fmt.Printf("key %v is exist and delete failed\n", key)
			return
		}
	}
	if s.l.insertToTail(key, val, expire) {
		s.setTakeUp()
		fmt.Printf("simpleKV tateup: %d\n", s.tateup)
	}
	if s.checkSizeOver() {
		fmt.Printf("wraning!!! simpleKV was overSize!, cap is %d, take up %d\n", s.cap, s.tateup)
		s.clearNode()
	}
}

func (s *simpleKV) Get(key string) (interface{}, bool) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	node := s.l.get(key)
	if node != nil {
		return node.value, true
	}
	return nil, false
}

func (s *simpleKV) Del(key string) bool {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	if s.l.isExist(key) {
		s.l.deleteNode(key)
		return true
	}
	return false
}

func (s *simpleKV) Exists(key string) bool {
	return s.l.isExist(key)
}

func (s *simpleKV) Flush() bool {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	s.l.cleanAll()
	return true
}

func (s *simpleKV) Keys() int64 {
	return int64(s.l.length)
}

func (s *simpleKV) ClearExpireNode() {
	for {
		select {
		case <-time.After(s.expireTime * time.Second):
			s.clearNode()
		}
	}
}

func (s *simpleKV) clearNode() {
	if keys := s.l.getExpireKeys(); len(keys) > 1 {
		for _, key := range keys {
			s.Del(key)
		}
	}
}

func NewCache(cap string, expireTime time.Duration) Cache {
	size, err := utils.Parse(cap)
	if err != nil {
		panic(err)
	}
	return &simpleKV{
		cap:        size,
		l:          newlinkList(),
		expireTime: expireTime,
	}
}

func (s simpleKV) getLong() int {
	return int(reflect.TypeOf(s).Size())
}

func (s simpleKV) getFieldMaxSize() int {
	var maxSize int
	itype := reflect.TypeOf(s)
	for i := 0; i < itype.NumField(); i++ {
		field := itype.Field(i)
		if int(field.Type.Size()) > maxSize {
			maxSize = int(field.Type.Size())
		}
	}
	return maxSize
}

func (s simpleKV) getStructAlign() int {
	var align int
	itype := reflect.TypeOf(s)
	for i := 0; i < itype.NumField(); i++ {
		field := itype.Field(i)
		if field.Type.Align() > align {
			align = field.Type.Align()
		}
	}
	return align
}

func (s simpleKV) getStructSize(long, size, align int) (cap int) {
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

func (s *simpleKV) checkSizeOver() bool {
	return s.tateup > s.cap
}

func (s *simpleKV) setTakeUp() bool {
	if s.l == nil {
		return false
	}
	var totalcap int
	var memory []imemory
	imemories := append(memory, *s, *s.l, *s.l.head)
	for i, v := range imemories {
		long := v.getLong()
		size := v.getFieldMaxSize()
		align := v.getStructAlign()
		if i == len(imemories)-1 {
			cap := v.getStructSize(long, size, align)
			totalcap += cap * int(s.l.length)
		} else {
			totalcap += long
		}
	}
	s.tateup = totalcap
	return true
}
