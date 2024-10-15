package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type userdata struct {
	Name string
	Pin  int
}

func main() {

	userinfo := map[int]string{

		1: "Leon Scott Kennedy",
		2: "Luis Sierra",
		3: "John Silver",
		4: "Luis Topez",
	}

	balanceinfo := map[string]int{

		"Leon Scott Kennedy": 6
		"Luis Sierra":        ,
		"John Silver":        "John Silver",
		"Luis Topez":         "Luis Topez",
	}

	var persons []userdata

	fmt.Println("---------------------")
	fmt.Println("Simple ATM Simulation")
	fmt.Println("---------------------")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">>> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Вхідна помилка:", err)
			return
		}

		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		text = strings.TrimSpace(text)

		switch text {
		case "/testcase":

			fmt.Println("Успішна перевірка команди, продовжуйте")

			if err != nil {
				fmt.Println("кейс /testcase не працює :(")
			}

		case "/Register":
			fmt.Println("Вхід в систему...")
			fmt.Println("---------------------")
			fmt.Println("Введіть ваші дані для входу:")

			var name string
			fmt.Print("Введіть номер картки: ")
			nameText, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Вхідна помилка:", err)
				continue
			}
			nameText = strings.Replace(nameText, "\n", "", -1)
			nameText = strings.TrimSpace(nameText)
			name = nameText

			var pin int
			fmt.Print("Введіть PIN: ")
			pinText, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Вхідна помилка:", err)
				continue
			}

			pinText = strings.Replace(pinText, "\n", "", -1)
			pinText = strings.TrimSpace(pinText)
			pin, err = strconv.Atoi(pinText)
			if err != nil {
				fmt.Println("Введіть корректне значення.")
				continue
			}

			person := userdata{
				Name: name,
				Pin:  pin,
			}

			persons = append(persons, person)

			fmt.Println("Реєстрація виконана успішно!")

		case "@Closing":
			fmt.Println("Закриття программи..")
			os.Exit(0)

		case "@Printing-a-check":
			f, err := os.Create("PrintedCheck_.txt")
			if err != nil {
				fmt.Println(err)
				return
			}

			l, err := f.WriteString("Інформація про чек:\r\n" + "=====================\r\n" + "Унікальний номер:\r\n")
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			fmt.Println(l, "#ЛОГ / Чек був створений успішно")
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}

		case "@Pay":
			fmt.Println("#:")
			fmt.Print("Введіть суму для поповнення: ")
			sumText, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Вхідна помилка:", err)
				continue
			}
			sumText = strings.Replace(sumText, "\n", "", -1)
			sumText = strings.TrimSpace(sumText)

			sum, err := strconv.Atoi(sumText)
			if err != nil {
				fmt.Println("Введіть корректне значення.")
				continue
			}
			if sum <= 0 {
				fmt.Println("Сума повинна бути більше нуля.")
				continue
			}
			fmt.Println("Поповнення виконано успішно!")

		case "@List":
			for i, person := range persons {
				fmt.Printf("Користувач %d:\n", i+1, 1)
				fmt.Printf("Номер карти: %s\n", person.Name)
				fmt.Printf("ПІН: %d\n", person.Pin)
				if err != nil {
					fmt.Println("Введіть правильно команду")

				}
			}

		case "@Help":
			fmt.Println("Команди: \r\n")
			fmt.Println("------------------------------------------")
			fmt.Println("@testcase - демонстрація тестового кейсу")
			fmt.Println("@Register - реєстрація користувача")
			fmt.Println("@Closing - закриття программи")
			fmt.Println("@Printing-a-check - створення чека")
			fmt.Println("@Pay - поповнення балансу")
			fmt.Println("@List - перегляд списку користувачів")
			fmt.Println("@Help - допомога по використанню системи")
			fmt.Println("------------------------------------------")

		case "@Allusers":
			fmt.Printf("Список всіх людей, хто користувався банкоматом:")
			fmt.Println("Користувач №1:\r\n", userinfo)

		default:
			fmt.Println("Впишіть команду для початку роботи.")
		}
	}
}

// Технічне завдання номер 2
