package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jaswdr/faker"
)

type Todo struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func request() string {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return "HTTP Error"
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "READ Error"
	}

	var t Todo
	if err := json.Unmarshal(body, &t); err != nil {
		return "JSON Error"
	}
	return t.Title
}

func main() {
	faker := faker.New()

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		name := faker.Person().Name()
		resp := request()
		fmt.Fprint(w, name+": "+resp)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		os.Exit(1)
	}
}
