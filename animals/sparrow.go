package animals

import "fmt"

// Структура Sparrow
type Sparrow struct {
    sound string
    move  string
}

func NewSparrow() *Sparrow {
    return &Sparrow{
        sound: "Писк",
        move:  "Летает",
    }
}

func (s *Sparrow) Sound() {
    fmt.Println("Воробей: ", s.sound)
}

func (s *Sparrow) Move() {
    fmt.Println("Воробей: ", s.move)
}

func (s *Sparrow) Fly() {
    fmt.Println("Воробей: Летает высоко")
}
