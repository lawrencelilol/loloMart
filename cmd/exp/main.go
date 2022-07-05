package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")

	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Jon Snow",
		Age:  23,
	}

	err = t.Execute(os.Stdout, user)

	if err != nil {
		panic(err)
	}

}
