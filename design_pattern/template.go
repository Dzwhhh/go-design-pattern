package design_pattern

import (
	"fmt"
)

type Skeleton struct {
	T Transformer
}

func (s *Skeleton) Merge() {
	s.T.BuildHead()
	s.T.BuildBody()
	s.T.BuildCore()
}

type Transformer interface {
	BuildHead()
	BuildBody()
	BuildCore()
}

type Human struct{}

func (*Human) BuildHead() {
	fmt.Println("我来组成--人类头部！")
}

func (*Human) BuildBody() {
	fmt.Println("我来组成--人类身体！")
}

func (*Human) BuildCore() {
	fmt.Println("我来组成--扁桃体！")
}

type Rocket struct{}

func (*Rocket) BuildHead() {
	fmt.Println("我来组成--火箭头部！")
}

func (*Rocket) BuildBody() {
	fmt.Println("我来组成--火箭管体！")
}

func (*Rocket) BuildCore() {
	fmt.Println("我来组成--引擎！")
}
