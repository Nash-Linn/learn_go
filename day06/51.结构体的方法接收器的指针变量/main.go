package main

import "fmt"

type Player struct {
	Name        string
	HealthPoint int
	Level       int
	NowPosition []int
	Prop        []string
}

func NewPlayer(name string, hp int, level int, np []int, prop []string) Player {
	return Player{
		Name:        name,
		HealthPoint: hp,
		Level:       level,
		NowPosition: np,
		Prop:        prop,
	}
}

func (p *Player) attack() {
	fmt.Printf("%s正在攻击\n", p.Name)
}

func (p *Player) attacked() {
	fmt.Printf("%s被攻击\n", p.Name)
	p.HealthPoint = p.HealthPoint - 10
}

func (p *Player) buyProp(propName string) {
	fmt.Printf("%s购买了%s\n", p.Name, propName)
	p.Prop = append(p.Prop, propName)
}

func main() {
	p1 := NewPlayer("Nash", 100, 1, []int{1, 1}, []string{"木剑"})
	p1.attack()
	p1.attacked()

	fmt.Println("p1的血槽值", p1.HealthPoint)

	p1.buyProp("药水")
	fmt.Println(p1)
}
