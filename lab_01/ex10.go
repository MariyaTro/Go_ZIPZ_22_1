package main

import "fmt"

func main() {
	var chartype int8 = 'R'

	fmt.Printf("Code '%c' - %d\n", chartype, chartype)

	// Завдання.
	// 1. Вивести українську літеру 'Ї'
	// 2. Пояснити призначення типу "rune"

	// Виконання
	// 1. Вивести українську літеру 'Ї'
	var char rune = 'Ї'
	fmt.Printf("Ukrainian letter 'Ї' is: %c\n", char)
	// 2. Пояснити призначення типу "rune"
	/*
		Тип "rune" в Go використовуються для збереження символів Юнікоду, він є синонімом для int32,
		оскільки кожен символ Юнікоду займає 4 байти. Таким чином, тип rune дає змогу працювати
		не лише латинськими алфавітами, але й іншими мовами та символами, що містяться в Юнікоді.
	*/
}
