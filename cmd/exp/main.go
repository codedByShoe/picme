package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.html")
	if err != nil {
		panic(err)
	}

	user := struct {
		Name string
	}{
		Name: "Andrew Shoemaker",
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
