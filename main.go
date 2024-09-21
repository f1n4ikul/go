package main

import (
	"fmt" 
	"strings"
	"bufio"
	"os"


	
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

type Cs interface {
	Cs()
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


type Swimmer interface {
	Swim()
}

// Структура Lion
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

// Структура Penguin
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

// Структура Sparrow
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

// Структура Shark
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

// Структура Elephant
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

func main() {
	var animals []Animal

	animals = append(animals, NewLion(), NewPenguin(), NewSparrow(), NewShark(), NewElephant())

	scanner := bufio.NewScanner(os.Stdin)

	var addMore string

	for {
		fmt.Println("Хотите добавить новое животное? (да/нет)")
		if !scanner.Scan() {
			fmt.Println("Ошибка при чтении ответа")
			continue
		}
		addMore = scanner.Text()

		if strings.ToLower(addMore) == "да" {

			var name, sound, move, custom_skill string
			fmt.Println("Введите имя животного:")

			if !scanner.Scan() {
				fmt.Println("Ошибка при вводе имени животного")
				continue
			}
			name = scanner.Text()

			fmt.Println("Введите звук животного:")
			if !scanner.Scan() {
				fmt.Println("Ошибка при вводе звука животного")
				continue
			}
			sound = scanner.Text()

			fmt.Println("Введите способ передвижения животного:")
			if !scanner.Scan() {
				fmt.Println("Ошибка при вводе способа передвижения животного")
				continue
			}
			move = scanner.Text()

			fmt.Println("Введите введите личное умение животного:")
			if !scanner.Scan() {
				fmt.Println("Ошибка при вводе личного умения")
				continue
			}
			custom_skill = scanner.Text()

			customAnimal := CustomAnimal{name: name, sound: sound, move: move, cs: custom_skill}
			animals = append(animals, customAnimal)
		} else if strings.ToLower(addMore) == "нет" {
			break
		} else {
			fmt.Println("Неверный ответ. Попробуйте ещё раз.")
		}

	}


	for _, animal := range animals {
		animal.Sound()
		animal.Move()

		if swimmer, ok := animal.(Swimmer); ok {
			swimmer.Swim()
		}

		if custom, ok := animal.(Cs); ok {
			custom.Cs()
		}

		switch a := animal.(type) {
		case Lion:
			a.Hunt()
		case Sparrow:
			a.Fly()
		case Elephant:
			a.Splash()
		}
	}
}
