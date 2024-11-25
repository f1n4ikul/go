package animals

import "fmt"

// Структура для кастомного животного
type CustomAnimal struct {
    Name        string
    SoundText   string
    MoveText    string
    CustomSkill string
}

// Метод Sound для CustomAnimal
func (ca *CustomAnimal) Sound() {
    fmt.Println(ca.Name,": ", ca.SoundText)
}

// Метод Move для CustomAnimal
func (ca *CustomAnimal) Move() {
    fmt.Println(ca.Name,": ", ca.MoveText)
}

// Функция для создания нового CustomAnimal
func NewCustomAnimal(name, sound, move, customSkill string) *CustomAnimal {
    return &CustomAnimal{
        Name:        name,
        SoundText:   sound,
        MoveText:    move,
        CustomSkill: customSkill,
    }
}