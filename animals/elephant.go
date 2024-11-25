package animals

import "fmt"

// Структура Elephant
type Elephant struct {
    sound string
    move  string
}

func NewElephant() *Elephant {
    return &Elephant{
        sound: "Трубит",
        move:  "Идёт медленно",
    }
}

func (e *Elephant) Sound() {
    fmt.Println("Слон: ", e.sound)
}

func (e *Elephant) Move() {
    fmt.Println("Слон: ", e.move)
}

func (e *Elephant) Splash() {
    fmt.Println("Слон: Обрызгивает водой")
}
