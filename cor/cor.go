// Chain of Responsibility: A chain of components who all get a chance to process a command
// or a query, optionally having default processing implementation and an ability to terminate
// the processing chain.
// It can be implemented as a linked list of pointers or a centralized construct.
// Control object removal from chain.

package main

import "fmt"

type Modifier interface {
	Add(m Modifier)
	Handle()
}

type Creature struct {
	Name string
	Attack, Defense int
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
}

func NewCreature(name string, attack int, defense int) *Creature {
	return &Creature{
		Name:    name,
		Attack:  attack,
		Defense: defense,
	}
}

// Modifier the same type creatures across the linked list
type CreatureModifier struct {
	creature *Creature
	next Modifier
}

func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}

func (c *CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

func NewCreatureModifier(creature *Creature) *CreatureModifier {
	return &CreatureModifier{creature: creature}
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling", d.creature.Name, "attack...")
	d.creature.Attack *= 2
	d.CreatureModifier.Handle()
}

func NewDoubleAttackModifier(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier{creature: c}}
}

type IncreasedDefenseModifier struct {
	CreatureModifier
}

func (i *IncreasedDefenseModifier) Handle() {
	fmt.Println("Increasing", i.creature.Name, "'s defense")
	i.creature.Defense++
	i.CreatureModifier.Handle()
}

func NewIncreasedDefenseModifier(c *Creature) *IncreasedDefenseModifier {
	return &IncreasedDefenseModifier{CreatureModifier{creature: c}}
}

func main() {
	goblin := NewCreature("Goblin", 1, 1)
	fmt.Println(goblin.String())

	root := NewCreatureModifier(goblin)
	root.Add(NewDoubleAttackModifier(goblin))
	root.Add(NewIncreasedDefenseModifier(goblin))
	// Traversal the linked list where every single modifier is applied to the creature
	root.Handle()
	fmt.Println(goblin.String())
}
