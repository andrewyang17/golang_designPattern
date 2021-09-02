package main

type CompositeAccountCommand struct {
	commands []Command
}

func (c *CompositeAccountCommand) Call() {
	for _, cmd := range c.commands {
		cmd.Call()
	}
}

func (c *CompositeAccountCommand) Undo() {
	for i := range c.commands {
		c.commands[len(c.commands)-i-1].Undo()
	}
}

func (c *CompositeAccountCommand) Succeeded() bool {
	for _, cmd := range c.commands {
		if !cmd.Succeeded() {
			return false
		}
	}
	return true
}

func (c *CompositeAccountCommand) SetSucceeded(value bool) {
	for _, cmd := range c.commands {
		cmd.SetSucceeded(value)
	}
}
