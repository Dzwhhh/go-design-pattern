package design_pattern

import (
	"fmt"
	"sync"
	"time"
)

type Train interface {
	RequestArrival()
	PermitArrival()
	Departure()
}

type PassengerTrain struct {
	Mediator mediator
}

type OilTrain struct {
	Mediator mediator
}

type GoodsTrain struct {
	Mediator mediator
}

type mediator interface {
	CanArrive(Train) bool
	NotifyFree()
	ArrangeDepart(duration time.Duration)
}

type PlatformMediator struct {
	isPlatformFree bool
	trainQueue     []Train
	muLuck         *sync.Mutex
}

func NewPlatformMediator() *PlatformMediator {
	return &PlatformMediator{
		isPlatformFree: true,
		trainQueue:     make([]Train, 0),
		muLuck:         new(sync.Mutex),
	}
}

func (p *PlatformMediator) CanArrive(train1 Train) bool {
	p.muLuck.Lock()
	defer p.muLuck.Unlock()
	p.trainQueue = append(p.trainQueue, train1)
	if len(p.trainQueue) > 1 {
		p.isPlatformFree = false
	} else {
		p.isPlatformFree = true
	}
	return p.isPlatformFree
}

func (p *PlatformMediator) NotifyFree() {
	p.muLuck.Lock()
	defer p.muLuck.Unlock()
	p.trainQueue = p.trainQueue[1:]
	if len(p.trainQueue) > 0 {
		arrivedTrain := p.trainQueue[0]
		arrivedTrain.PermitArrival()
	} else {
		fmt.Println("platform is empty")
	}
}

func (p *PlatformMediator) ArrangeDepart(duration time.Duration) {
	for {
		if len(p.trainQueue) > 0 {
			leavedTrain := p.trainQueue[0]
			leavedTrain.Departure()
		} else {
			fmt.Println("platform is empty")
		}
		time.Sleep(duration)
	}
}

func (p *PassengerTrain) RequestArrival() {
	if p.Mediator.CanArrive(p) {
		p.PermitArrival()
	} else {
		fmt.Println("passenger Train waiting", time.Now())
	}
}

func (p *PassengerTrain) PermitArrival() {
	fmt.Println("passenger Train has arrived", time.Now())
}

func (p *PassengerTrain) Departure() {
	fmt.Println("passenger Train leaving", time.Now())
	p.Mediator.NotifyFree()
}

func (g *GoodsTrain) RequestArrival() {
	if g.Mediator.CanArrive(g) {
		g.PermitArrival()
	} else {
		fmt.Println("goods Train waiting", time.Now())
	}
}

func (g *GoodsTrain) PermitArrival() {
	fmt.Println("goods Train has arrived", time.Now())
}

func (g *GoodsTrain) Departure() {
	fmt.Println("goods Train leaving", time.Now())
	g.Mediator.NotifyFree()
}

func (o *OilTrain) RequestArrival() {
	if o.Mediator.CanArrive(o) {
		o.PermitArrival()
	} else {
		fmt.Println("oil Train waiting", time.Now())
	}
}

func (o *OilTrain) PermitArrival() {
	fmt.Println("oil Train has arrived", time.Now())
}

func (o *OilTrain) Departure() {
	fmt.Println("oil Train leaving", time.Now())
	o.Mediator.NotifyFree()
}
