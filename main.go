package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func importFileToString(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	resultToString := string(content)
	return resultToString
}

func main() {
	jsonString := importFileToString("index.json")
	type Fruit struct {
		Apple  string `json:"apple"`
		Orange string `json:"orange"`
		Banana string `json:"banana"`
	}
	type Person struct {
		Title string `json:"title"`
		Name  string `json:"name"`
		Fruit Fruit  `json:"fruit"`
	}

	var person Person

	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf(person.Fruit.Apple)
	fmt.Printf("%+v\n", person)

	// m := map[string]int{"foo": 42}
	// log.Print(m["foo"])
}
