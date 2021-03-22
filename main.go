package main

import (
	"encoding/json"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

type Person struct {
	FirstName string
}

func main() {

	/**
	p1 := Person{FirstName: "Temp1"}
	p2 := Person{FirstName: "Temp2"}

	persons := []Person{p1, p2}

	bytes, err := json.Marshal(persons)

	fmt.Println(string(bytes))
	fmt.Println(err)
	var personsUnmarshalled []Person
	err = json.Unmarshal(bytes, &personsUnmarshalled)

	fmt.Println(personsUnmarshalled)
	*/

	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)
	// uses default serveMux
	http.ListenAndServe(":8080", nil)
}

func encode(w http.ResponseWriter, r *http.Request) {
	p1 := Person{FirstName: "Bondss"}
	err := json.NewEncoder(w).Encode(p1)

	if err != nil {
		log.Printf("encoded bad data : %s", err)
	}
}

// https://curlbuilder.com/
// curl -H "Content-type: application/json" -d '{"FirstName": "Gaurav"}' 'http://localhost:8080/decode'
func decode(w http.ResponseWriter, r *http.Request) {
	p1 := Person{}
	err := json.NewDecoder(r.Body).Decode(&p1)

	if err != nil {
		log.Printf("decoded bad data : %s", err)
	}
	log.Println("Person:", p1)
}
