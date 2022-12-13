package crates

type Crane struct {
	crates       []Stack
	instructions []Instruction
}

func NewCrane(crates []Stack, instructions []Instruction) *Crane {
	return &Crane{
		crates:       crates,
		instructions: instructions,
	}
}

func (c *Crane) Run() {
	for _, i := range c.instructions {
		//c.step(i)
		c.step9001(i)
	}
}

func (c *Crane) step(ins Instruction) {
	for i := 0; i < ins.Crates(); i++ {
		var crate string
		crate, c.crates[ins.FromIndex()] = c.crates[ins.FromIndex()].Pop()
		c.crates[ins.ToIndex()] = c.crates[ins.ToIndex()].Push(crate)
	}
}

func (c *Crane) step9001(ins Instruction) {
	var poppedCrates Stack
	to := ins.ToIndex()
	from := ins.FromIndex()
	poppedCrates, c.crates[from] = c.crates[from].MultiplePop(ins.Crates())
	c.crates[to] = c.crates[to].MultiplePush(poppedCrates)

}

func (c *Crane) Tops() string {
	tops := ""
	for _, column := range c.crates {
		tops += column.Top()
	}

	return tops
}
