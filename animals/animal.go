package animals

import (
	"fmt"
)

type Animal interface {
	Sound()
	Move()
}


type CustomAnimal struct {
	name  string
	sound string
	move  string
	cs 	  string
}


func (ca CustomAnimal) Sound() {
	fmt.Printf("%s: %s\n", ca.name, ca.sound)
}

func (ca CustomAnimal) Move() {
	fmt.Printf("%s: %s\n", ca.name, ca.move)
}

func (ca CustomAnimal) Cs() {
	fmt.Printf("%s: %s\n", ca.name, ca.cs)
}



type Lion struct {
	sound string
	move  string
}

func NewLion() Lion {
	return Lion{
		sound: "Рычание",
		move:  "Бегает",
	}
}

func (l Lion) Sound() {
	fmt.Println("Лев: ", l.sound)
}

func (l Lion) Move() {
	fmt.Println("Лев: ", l.move)
}

func (l Lion) Hunt() {
	fmt.Println("Лев: Охотится на добычу")
}


type Penguin struct {
	sound string
	move  string
}

func NewPenguin() Penguin {
	return Penguin{
		sound: "Крякает",
		move:  "Плавает",
	}
}

func (p Penguin) Sound() {
	fmt.Println("Пингвин: ", p.sound)
}

func (p Penguin) Move() {
	fmt.Println("Пингвин: ", p.move)
}

func (p Penguin) Swim() {
	fmt.Println("Пингвин: Плавает в воде")
}


type Sparrow struct {
	sound string
	move  string
}

func NewSparrow() Sparrow {
	return Sparrow{
		sound: "Писк",
		move:  "Летает",
	}
}

func (s Sparrow) Sound() {
	fmt.Println("Воробей: ", s.sound)
}

func (s Sparrow) Move() {
	fmt.Println("Воробей: ", s.move)
}

func (s Sparrow) Fly() {
	fmt.Println("Воробей: Летает высоко")
}


type Shark struct {
	sound string
	move  string
}

func NewShark() Shark {
	return Shark{
		sound: "Рычание",
		move:  "Плавает быстро",
	}
}

func (sh Shark) Sound() {
	fmt.Println("Акула: ", sh.sound)
}

func (sh Shark) Move() {
	fmt.Println("Акула: ", sh.move)
}

func (sh Shark) Swim() {
	fmt.Println("Акула: Плавает в океане")
}


type Elephant struct {
	sound string
	move  string
}

func NewElephant() Elephant {
	return Elephant{
		sound: "Трубит",
		move:  "Идёт медленно",
	}
}

func (e Elephant) Sound() {
	fmt.Println("Слон: ", e.sound)
}

func (e Elephant) Move() {
	fmt.Println("Слон: ", e.move)
}

func (e Elephant) Splash() {
	fmt.Println("Слон: Обрызгивает водой")
}


func (ca CustomAnimal) Sound() {
	fmt.Printf("%s: %s\n", ca.name, ca.sound)
}

func (ca CustomAnimal) Move() {
	fmt.Printf("%s: %s\n", ca.name, ca.move)
}

func (ca CustomAnimal) Cs() {
	fmt.Printf("%s: %s\n", ca.name, ca.cs)
}