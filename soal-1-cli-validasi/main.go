package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var errInvalidName = errors.New("nama harus diisi")
var errInvalidAge = errors.New("umur tidak valid (minimal 18 tahun)")

func getInput(question string) string {
	fmt.Print(question)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

func checkInputValidity(name string, age int) (bool, error) {
	var errType error

	if name == "" {
		errType = errInvalidName
	} else if age < 18 {
		errType = errInvalidAge
	}

	var err error

	if errType != nil {
		err = fmt.Errorf("Error: %w", errType)
	}

	return errType == nil, err
}

func main() {
	name := getInput("Nama: ")
	age, _ := strconv.Atoi(getInput("Umur: "))

	isValid, err := checkInputValidity(name, age)

	if !isValid {
		fmt.Println(err)
		return
	}

	fmt.Printf("Selamat datang %s", name)
}
