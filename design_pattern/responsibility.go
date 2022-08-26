package design_pattern

import (
	"fmt"
)

type Patient struct {
	Name        string
	Charge      int32
	HasPaid     bool
	HasChecked  bool
	HasMedicine bool
}

type Department interface {
	Execute(*Patient)
	SetNext(Department) Department
}

type Reception struct {
	next Department
}

func (r *Reception) Execute(p *Patient) {
	fmt.Println("进入门诊")
	if !p.HasPaid {
		p.Charge -= 2000
		if p.Charge < 0 {
			p.Charge += 2000
			fmt.Println("余额不足，支付失败")
			return
		}
		p.HasPaid = true
	}
	r.next.Execute(p)
}

func (r *Reception) SetNext(department Department) Department {
	r.next = department
	return r.next
}

type Doctor struct {
	next Department
}

func (d *Doctor) Execute(p *Patient) {
	fmt.Println("看医生")
	Name := p.Name
	if Name == "张三" {
		fmt.Println("医生：哈哈哈哈，张三你死定了")
		p.HasChecked = false // 医生和张三有仇
	} else {
		p.HasChecked = true
	}
	d.next.Execute(p)
}

func (d *Doctor) SetNext(department Department) Department {
	d.next = department
	return d.next
}

type MedicineRoom struct {
	next Department
}

func (m *MedicineRoom) Execute(p *Patient) {
	fmt.Println("拿药")
	if p.HasChecked && p.HasPaid {
		p.HasMedicine = true
		fmt.Println("药拿对了，病人活了")
	} else if p.HasPaid {
		p.HasMedicine = true
		fmt.Println("拿的是毒药，病人死了")
	} else {
		p.HasMedicine = false
		fmt.Println("没有药方，无法开药")
	}
}

func (m *MedicineRoom) SetNext(department Department) Department {
	m.next = department
	return m.next
}
