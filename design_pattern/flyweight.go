package design_pattern

import "fmt"

const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

const (
	TerroristPlayerType        = "terrorist"
	CounterTerroristPlayerType = "counterTerrorist"
)

type DressFactory struct {
	createTime int
	dresses    map[string]Dress
}

var dressFactory *DressFactory

func getDressFactoryInstance() *DressFactory {
	if dressFactory == nil {
		fmt.Println("create dress factory")
		dressFactory = &DressFactory{
			createTime: 0,
			dresses:    make(map[string]Dress),
		}
		return dressFactory
	} else {
		return dressFactory
	}
}

func (d *DressFactory) getOrCreateFlyweightObject(dressType string) Dress {
	if dress, exist := d.dresses[dressType]; exist {
		return dress
	} else {
		// 创建享元
		d.accCreateTime()
		if dressType == TerroristDressType {
			newDress := &TerroristDress{color: "red"}
			d.dresses[TerroristDressType] = newDress
			return newDress
		} else if dressType == CounterTerroristDressType {
			newDress := &CounterTerroristDress{color: "green"}
			d.dresses[CounterTerroristDressType] = newDress
			return newDress
		} else {
			panic("unknown type")
		}
	}
}

func (d *DressFactory) accCreateTime() {
	d.createTime += 1
	fmt.Println("create dress time:", d.createTime)
}

type Dress interface {
	getColor() string
}

type CounterTerroristDress struct {
	color string
}

func (c *CounterTerroristDress) getColor() string {
	return c.color
}

type TerroristDress struct {
	color string
}

func (t *TerroristDress) getColor() string {
	return t.color
}

type PlayerFactory struct {
	totalPlayer int
}

func (p *PlayerFactory) CreatePlayer(playerType string) Player {
	p.totalPlayer += 1

	if playerType == TerroristPlayerType {
		return &Terrorist{
			dress:      getDressFactoryInstance().getOrCreateFlyweightObject(TerroristDressType),
			number:     p.totalPlayer,
			playerType: TerroristPlayerType,
		}
	} else if playerType == CounterTerroristPlayerType {
		return &CounterTerrorist{
			dress:      getDressFactoryInstance().getOrCreateFlyweightObject(CounterTerroristDressType),
			number:     p.totalPlayer,
			playerType: CounterTerroristPlayerType,
		}
	} else {
		panic("unknown player type")
	}
}

type Player interface {
	Info()
	setNumber(int)
}

type Terrorist struct {
	number     int
	dress      Dress
	playerType string
}

func (t *Terrorist) setNumber(number int) {
	t.number = number
}
func (t *Terrorist) Info() {
	fmt.Printf("Player %d: %s, dress %s\n", t.number, t.playerType, t.dress.getColor())
}

type CounterTerrorist struct {
	number     int
	dress      Dress
	playerType string
}

func (c *CounterTerrorist) setNumber(number int) {
	c.number = number
}
func (c *CounterTerrorist) Info() {
	fmt.Printf("Player %d: %s, dress %s\n", c.number, c.playerType, c.dress.getColor())
}
