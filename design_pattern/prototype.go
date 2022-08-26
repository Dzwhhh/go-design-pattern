package design_pattern

import "fmt"

type absProto interface {
	Print()
	Clone() absProto
}

type ProtoA struct {
	nodes []string
	kv    map[int]string
}

func (a *ProtoA) AppendNodes(n []string) {
	a.nodes = append(a.nodes, n...)
}

func (a *ProtoA) UpdateKey(key int, value string) {
	if a.kv == nil {
		a.kv = make(map[int]string)
	}
	a.kv[key] = value
}

func (a *ProtoA) Print() {
	fmt.Printf("nodes: %v \n kv:%v \n", a.nodes, a.kv)
}

func (a *ProtoA) Clone() *ProtoA {
	b := new(ProtoA)
	b.kv = make(map[int]string)

	for _, v := range a.nodes {
		b.nodes = append(b.nodes, v)
	}
	for k, v := range a.kv {
		b.kv[k] = v
	}
	return b
}
