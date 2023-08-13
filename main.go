package main

import ( 
	"database/sql"
	"math/rand"
	"bufio"
	"fmt"
	"os"
	"errors"
)

const NUMBER_OF_COUNTRIES int = 138
const NUMBER_OF_FIRSTNAMES int = 100 
const NUMBER_OF_LASTNAMES int = 100

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
	firstname, err := readNthLineFromFile("./firstnames.dat", rand.Intn(NUMBER_OF_FIRSTNAMES+1))
	checkError(err)
	return firstname
}

func getRandomLastName() string {
	lastname, err := readNthLineFromFile("./lastnames.dat", rand.Intn(NUMBER_OF_LASTNAMES+1))
	checkError(err)
	return lastname
}

func getRandomCountry() string {
	country, err := readNthLineFromFile("./countries.dat", rand.Intn(NUMBER_OF_COUNTRIES+1))
	checkError(err)
	return country 
}

func newRandomRunner() *runner {
	r := runner{
		firstname: getRandomFirstName(),
		lastname:  getRandomLastName(),
		country:   getRandomCountry(),
		id:        1,
	}
	return &r
}

func main() {
	db, err := sql.Open("sqlite3", "./runners.db")
	checkError(err)
	defer db.Close()

	r := newRandomRunner()
	//fmt.Println(r)
	//fmt.Println(*r)
	addRunner(db, *r)

	runners := getAllRunners(db)
	for _, r := range(runners) { 
		fmt.Println(r)
	}
	
	deleteRunner(db, 1)
}
