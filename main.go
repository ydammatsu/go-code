package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jaswdr/faker"
)

func main() {
	faker := faker.New()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := faker.Person().Name()
		fmt.Fprint(w, name)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		os.Exit(1)
	}
}
