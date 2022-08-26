package design_pattern

type Depart interface {
	GetName() string
	GetStudentNum() int
}

type MathDepartment struct{}

func (m *MathDepartment) GetName() string {
	return "math"
}

func (m *MathDepartment) GetStudentNum() int {
	return 100
}

type CSDepartment struct{}

func (c *CSDepartment) GetName() string {
	return "computer science"
}

func (c *CSDepartment) GetStudentNum() int {
	return 250
}

type NullDepartment struct{}

func (n *NullDepartment) GetName() string {
	return ""
}

func (n *NullDepartment) GetStudentNum() int {
	return 0
}

type College struct {
	departments []Depart
}

func (c *College) GetDepartment(name string) Depart {
	for _, depart := range c.departments {
		if depart.GetName() == name {
			return depart
		}
	}
	return &NullDepartment{}
}

func (c *College) AddDepartment(d Depart) {
	c.departments = append(c.departments, d)
}

func CreateCollege() *College {
	college := &College{}
	mathD := &MathDepartment{}
	csD := &CSDepartment{}

	college.AddDepartment(mathD)
	college.AddDepartment(csD)

	return college
}
