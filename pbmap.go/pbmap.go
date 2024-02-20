package main

import "fmt"

type Person struct {
	Name  string
	Phone string
}

var phoneBook []Person

func main() {
	for {
		fmt.Println("1. Kişi Ekle")
		fmt.Println("2. Kişileri Listele")
		fmt.Println("3. Kişi sil")

		var secim int
		fmt.Print("Seçiminizi giriniz: ")
		fmt.Scanln(&secim)

		switch secim {
		case 1:
			addPerson()
		case 2:
			listPeople()

		}
	}
}

func addPerson() {
	var phone, name string

	fmt.Print("Ad Giriniz: ")
	fmt.Scanln(&name)

	fmt.Print("Telefon giriniz: ")
	fmt.Scanln(&phone)

	fmt.Println("Kişi Eklendi :)")

	person := Person{Name: name, Phone: phone}
	phoneBook = append(phoneBook, person)

	return
}

func listPeople() {
	fmt.Println("\n Telefon Rehberi:")
	for i, person := range phoneBook {
		fmt.Printf("%d. Adı %s, Telefon: %s\n", i+1, person.Name, person.Phone)
	}

	return
}
