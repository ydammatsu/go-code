package main

import (
	"fmt"

	"github.com/jaswdr/faker"
)

func main() {
	faker := faker.New()
	name := faker.Person().Name()
	fmt.Println(name)
}
