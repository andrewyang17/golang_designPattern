// Singleton: A component which is instantiated only once.
// For some components it only makes sense to have one in the system
// like Database repository, Object factory (stateless). (eg. the construction call is expensive)
// Adhere to DIP: Singleton depends on interfaces, not concrete type.
// Then you can replace Singleton with some sort of test dummy.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var once sync.Once
var instance Database

type Database interface {
	GetPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

func readData(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}
	return result, nil
}

func GetSingletonDatabase() Database {
	once.Do(func() {
		db := singletonDatabase{}
		capitals, err := readData("singleton/capitals.txt")
		if err != nil {
			panic(err)
		}
		db.capitals = capitals
		instance = &db
	})
	return instance
}

func main() {
	db := GetSingletonDatabase()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a city:")
	for scanner.Scan() {
		getCity := scanner.Text()
		population := db.GetPopulation(getCity)
		if population == 0 {
			fmt.Println("Please enter a correct city:")
		} else {
			fmt.Printf("The population in %v is %d \n", getCity, population)
			break
		}
	}
}
