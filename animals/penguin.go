package animals

import "fmt"

// Структура Penguin
type Penguin struct {
    sound string
    move  string
}

func NewPenguin() *Penguin {
    return &Penguin{
        sound: "Крякает",
        move:  "Плавает",
    }
}

func (p *Penguin) Sound() {
    fmt.Println("Пингвин: ", p.sound)
}

func (p *Penguin) Move() {
    fmt.Println("Пингвин: ", p.move)
}

func (p *Penguin) Swim() {
    fmt.Println("Пингвин: Плавает в воде")
}
