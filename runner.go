package main

import (
	"bufio"
	"fmt"
	"os"
	"errors"
)

type runner struct {
	firstname string
	lastname  string
	country   string
	id        int
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func readNthLineFromFile(filepath string, line int) (string, error) {
	f, err := os.Open(filepath)
	checkError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for l := 1; scanner.Scan(); l++ {
		if l == line {
			return scanner.Text(), nil
		}
	}
	checkError(scanner.Err())

	return "", errors.New("Could not find line in file")
}

func getRandomFirstName() string {
	return ""
}

func newRunner() *runner {
	r := runner{
		firstname: "John",
		lastname:  "Doe",
		country:   "Canada",
		id:        1,
	}
	return &r
}

func main() {
	line, err := readNthLineFromFile("./countries.dat", 1000)
	checkError(err)
	fmt.Println(line)
}
