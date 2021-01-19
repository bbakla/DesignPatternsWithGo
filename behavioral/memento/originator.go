package memento

type Memento struct {
	state string
}

//Originator
type Originator struct {
	state string
}

func (o *Originator) createMemento() Memento {
	return Memento{
		state: o.state,
	}
}

func (o *Originator) setState(state string) {
	o.state = state
}

type Caretaker struct {
	history    []Memento
	Originator *Originator
}

func (c *Caretaker) undo(index int) {
	c.Originator.setState(c.history[index].state)
}

func (c *Caretaker) createBackup() {
	c.history = append(c.history, c.Originator.createMemento())
}
