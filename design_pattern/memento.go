package design_pattern

type Memento struct {
	state string
}

func (m *Memento) getSavedState() string {
	return m.state
}

type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) CreateMemo() *Memento {
	return &Memento{
		state: o.state,
	}
}

func (o *Originator) RestoreMemo(m *Memento) {
	o.state = m.getSavedState()
}

type MemoRecords struct {
	memos []*Memento
}

func (mr *MemoRecords) AddMemo(m *Memento) {
	mr.memos = append(mr.memos, m)
}

func (mr *MemoRecords) FindMemo(idx int) *Memento {
	if mr.memos[idx] != nil {
		return mr.memos[idx]
	} else {
		panic("memo doesn't exist")
	}
}
