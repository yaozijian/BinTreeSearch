package main

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	zone struct {
		val   int
		min   int
		max   int
		left  *zone
		right *zone
	}

	tree struct {
		root *zone
	}

	item struct {
		a int
		b int
	}
)

func main() {

	t := &tree{}

	list := random_list()
	for _, v := range list {
		t.insert(v)
	}
	fmt.Println("\n\n")

	total := 0
	for i := 0; i < 100; i++ {
		total += t.find(i)
	}

	avg := float32(total) / 100
	fmt.Printf("len=%v avg=%.3f\n", len(list), avg)
}

func (this *tree) insert(x *item) {

	// 生成新节点
	p := &zone{
		val: (x.a + x.b) / 2,
		min: x.a,
		max: x.b,
	}

	// 找插入位置
	prv := (*zone)(nil)
	cur := this.root

	for cur != nil {
		prv = cur
		if p.val <= cur.val {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}

	// 插入到树中
	if prv == nil {
		this.root = p
	} else if p.val <= prv.val {
		prv.left = p
	} else {
		prv.right = p
	}
}

func (this *tree) find(val int) (cnt int) {

	cur := this.root

	for cur != nil {
		cnt++
		if val < cur.min {
			cur = cur.left
		} else if val <= cur.max {
			//fmt.Printf("%v ∈[%v,%v]\n", val, l.val, r.val)
			return
		} else {
			cur = cur.right
		}
	}

	//fmt.Printf("%v not found\n", val)

	return
}

func random_list() []*item {

	rand.Seed(time.Now().UnixNano())

	list := make([]*item, 0, 100)

	a := 1

	// 生成随机范围
	for a < 1000 {
		b := a + 1 + rand.Intn(24)
		n := &item{a: a, b: b}
		list = append(list, n)
		a = b + 1 + rand.Intn(100)%2*rand.Intn(10) // 随机产生裂缝
	}

	// 次序打乱
	newlist := make([]*item, len(list))

	for idx := range newlist {

		x := rand.Intn(len(list))

		for list[x] == nil {
			x = rand.Intn(len(list))
		}

		newlist[idx] = list[x]
		list[x] = nil
	}

	return newlist
}
