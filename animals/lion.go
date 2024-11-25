package animals

import "fmt"

// Структура Lion
type Lion struct {
    sound string
    move  string
}

func NewLion() *Lion {
    return &Lion{
        sound: "Рычание",
        move:  "Бегает",
    }
}

func (l *Lion) Sound() {
    fmt.Println("Лев: ", l.sound)
}

func (l *Lion) Move() {
    fmt.Println("Лев: ", l.move)
}

func (l *Lion) Hunt() {
    fmt.Println("Лев: Охотится на добычу")
}
