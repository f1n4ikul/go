package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"goapp/animals"
	"goapp/interfaces"
	"github.com/gen2brain/beeep"
)

func main() {
	var animalsList []interfaces.Animal
	var wg sync.WaitGroup

	// Инициализируем животных
	animalsList = append(animalsList, animals.NewLion(), animals.NewPenguin(), animals.NewSparrow(), animals.NewShark(), animals.NewElephant())

	scanner := bufio.NewScanner(os.Stdin)
	var addMore string

	// Обработка ввода пользователя для добавления новых животных
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

			fmt.Println("Введите личное умение животного:")
			if !scanner.Scan() {
				fmt.Println("Ошибка при вводе личного умения")
				continue
			}
			custom_skill = scanner.Text()

			// Создаем экземпляр CustomAnimal и добавляем его в список
			customAnimal := animals.NewCustomAnimal(name, sound, move, custom_skill)
			animalsList = append(animalsList, customAnimal)

		} else if strings.ToLower(addMore) == "нет" {
			break
		} else {
			fmt.Println("Неверный ответ. Попробуйте ещё раз.")
		}
	}

	// Обрабатываем каждого животного в горутине
	for _, animal := range animalsList {
		wg.Add(1)
		go func(animal interfaces.Animal) {
			defer wg.Done()

			// Выводим звук и движение
			animal.Sound()
			animal.Move()

			// Обрабатываем дополнительные интерфейсы, такие как Swimmer и Cs
			if swimmer, ok := animal.(interfaces.Swimmer); ok {
				swimmer.Swim()
			}

			if custom, ok := animal.(interfaces.Cs); ok {
				custom.CustomSkill()
			}

			// Отправляем уведомление с использованием beeep.Notify в горутине
			go func() {
				beeep.Notify("Завершено", fmt.Sprintf("Животное %T обработано.", animal), "")
			}()

			// Специфичные действия для каждого типа животного
			switch a := animal.(type) {
			case *animals.Lion:
				a.Hunt()
			case *animals.Sparrow:
				a.Fly()
			case *animals.Elephant:
				a.Splash()
			}
		}(animal)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
	fmt.Println("Все животные обработаны!")
}
