package main

import "fmt"

type I interface {
	Get() int
	Put(int)
}

type S struct {
	i int
}

func (p *S) Get() int {
	return p.i
}

func (p *S) Put(i int) {
	p.i = i
}

func fn(p I) {
	p.Put(10)
	fmt.Println(p.Get())
}

func main() {
	obj := &S{i: 5}
	fn(obj)
}
