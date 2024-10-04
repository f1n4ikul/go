package main

import (
    "database/sql"
    "fmt"
    "log"
    "interface-go/animalbd"
    _ "github.com/lib/pq"
)





const (
	host = "localhost"2
    port = 5432
    user = "postgres"
    password = "postgres"
    dbname = "Animals"

)

func main() {
    connStr := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    animals, err := animaldb.GetAllAnimals(db)
    if err != nil {
        log.Fatal(err)
    }

    

    var n int
    fmt.Println("Сортировка:\n1 - ID\n2 - Name")
    fmt.Scan(&n)
    switch n {
        case 2:
            animals, err = animaldb.GetSortedAnimalsByName(db)
            if err != nil {
                log.Fatal(err)
            }
            fmt.Println("Вывод всех животных по Name:")
            for _, animal := range animals {
                fmt.Printf("ID: %d, Name: %s, Species: %s, Age: %d, Habitat: %s\n", animal.ID, animal.Name, animal.Species, animal.Age, animal.Habitat)
            }

        case 1:
            var asc string

            
            fmt.Println("Выберите порядок сортировки:\nПо возрастанию - asc\nПо убыванию - desc")
            fmt.Scan(&asc)
            switch asc {
                case "asc":
                    animals, err = animaldb.GetSortedAnimalsByIdASC(db)
                    if err != nil {
                        log.Fatal(err)
                    }
                    fmt.Printf("Вывод всех животных по ID(%s):\n", asc)
                    for _, animal := range animals {
                        fmt.Printf("ID: %d, Name: %s, Species: %s, Age: %d, Habitat: %s\n", animal.ID, animal.Name, animal.Species, animal.Age, animal.Habitat)
                    }

                case "desc":
                    animals, err = animaldb.GetSortedAnimalsByIdDESC(db)
                    if err != nil {
                        log.Fatal(err)
                    }
                    fmt.Printf("Вывод всех животных по ID(%s):\n", asc)
                    for _, animal := range animals {
                        fmt.Printf("ID: %d, Name: %s, Species: %s, Age: %d, Habitat: %s\n", animal.ID, animal.Name, animal.Species, animal.Age, animal.Habitat)
                    }

            }
        // case 3:
        //     var name, species, habitat string
        //     var age int
        //     fmt.Println("Введите имя животного")
        //     fmt.Scan(&name)
        //     fmt.Println("Ваше животное травоядное/плотоядное")
        //     fmt.Scan(&species)
        //     fmt.Println("Введите возраст животного")
        //     fmt.Scan(&age)
        //     fmt.Println("Введите животное в каком месте обитает животное")
        //     fmt.Scan(&habitat)

        //     animals, err = animaldb.AddAnimal(db, name, species, age, habitat)
        //     if err!= nil {
        //         log.Fatal(err)
        //     }

        //     for _, animal := range animals {
        //         fmt.Printf("Добавлено животное: %s, %s, %d, %s\n", animal.Name, animal.Species, animal.Age, animal.Habitat)
        //     }
            

    }

}