package models

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestInsertToTail(t *testing.T) {
	l := newlinkList()
	for i := 0; i < 10; i++ {
		l.insertToTail("hello"+fmt.Sprintf("%d", i+1), "world"+fmt.Sprintf("%d", i+1), time.Duration(rand.Int63n(10000)))
	}
	l.print()
}

func TestLinkList_DeleteNode(t *testing.T) {
	l := newlinkList()
	for i := 0; i < 10; i++ {
		l.insertToTail("hello"+fmt.Sprintf("%d", i+1), "world"+fmt.Sprintf("%d", i+1), time.Duration(rand.Int63n(10000)))
	}
	l.print()

	t.Log(l.deleteNode("hello"))
	l.print()

	t.Log(l.deleteNode("hello1"))
	l.print()

	t.Log(l.deleteNode("hello2"))
	l.print()

	t.Log(l.deleteNode("hello5"))
	l.print()

	t.Log(l.deleteNode("hello10"))
	l.print()

	l.insertToTail("hello"+fmt.Sprintf("%d", 11), "world"+fmt.Sprintf("%d", 11), time.Duration(rand.Int63n(10000)))
	l.print()
}

func TestCleanAll(t *testing.T) {
	l := newlinkList()
	for i := 0; i < 10; i++ {
		l.insertToTail("hello"+fmt.Sprintf("%d", i+1), "world"+fmt.Sprintf("%d", i+1), time.Duration(rand.Int63n(10000)))
	}
	l.print()
	l.cleanAll()
	l.print()
}
