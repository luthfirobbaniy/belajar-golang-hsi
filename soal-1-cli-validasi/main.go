package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getName() (string, error) {
	fmt.Print("Nama: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	return input, err
}

func getAge() (int, error) {
	fmt.Print("Umur: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	inputInt, _ := strconv.Atoi(input)

	return inputInt, err
}

func checkAgeValidity(age int) bool {
	return age >= 18
}

func main() {
	name, _ := getName()
	age, _ := getAge()

	isAgeValid := checkAgeValidity(age)

	if isAgeValid {
		fmt.Printf("Selamat datang %s", name)
	} else {
		err := errors.New("umur tidak valid (minimal 18 tahun)")
		err = fmt.Errorf("Error: %w", err)
		fmt.Println(err)
		// log.WithError(err).Error(err)
	}
}
