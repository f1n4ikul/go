package animals

import "fmt"

// Структура Shark
type Shark struct {
    sound string
    move  string
}

func NewShark() *Shark {
    return &Shark{
        sound: "Рычание",
        move:  "Плавает быстро",
    }
}

func (sh *Shark) Sound() {
    fmt.Println("Акула: ", sh.sound)
}

func (sh *Shark) Move() {
    fmt.Println("Акула: ", sh.move)
}

func (sh *Shark) Swim() {
    fmt.Println("Акула: Плавает в океане")
}
