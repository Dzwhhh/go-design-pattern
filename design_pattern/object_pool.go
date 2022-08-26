package design_pattern

import (
	"fmt"
	"strconv"
	"sync"
)

/*对象池
  InitPool 初始化对象池
  Loan 获取可用对象
  Destroy 将对象放回对象池
*/

type PoolObject interface {
	GetId() string
}

type Connection struct{ id string }

func (c Connection) GetId() string {
	return c.id
}

type Pool struct {
	idle     []PoolObject
	active   []PoolObject
	capacity int
	muLock   *sync.Mutex
}

var pool *Pool

func InitPool(capacity int) *Pool {
	if pool == nil {
		fmt.Println("pool is nil, create...")
		pool = new(Pool)
		pool.idle = make([]PoolObject, 0)
		pool.active = make([]PoolObject, 0)
		pool.capacity = capacity
		pool.muLock = new(sync.Mutex)

		for i := 0; i < capacity; i++ {
			conn := Connection{strconv.Itoa(i)}
			pool.idle = append(pool.idle, conn)
		}
	}
	return pool
}

func (p *Pool) Loan() PoolObject {
	p.muLock.Lock()
	defer p.muLock.Unlock()

	if len(p.idle) == 0 {
		panic("idle is empty")
	}
	poolObject := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, poolObject)
	return poolObject
}

func (p *Pool) Destroy(target PoolObject) {
	p.muLock.Lock()
	defer p.muLock.Unlock()

	if len(p.active) == 0 {
		panic("active is empty")
	}
	for i, activeItem := range p.active {
		if activeItem.GetId() == target.GetId() {
			p.active = append(p.active[0:i], p.active[i+1:]...)
			p.idle = append(p.idle, activeItem)
			return
		}
	}
	panic("active id not find")
}
