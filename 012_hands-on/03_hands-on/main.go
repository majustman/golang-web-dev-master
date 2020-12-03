package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type hotel struct {
	Name, Address, City, Zip, Region string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sb := hotel{
		Name:    "South Beach",
		Address: "Mexican street, 20",
		City:    "LA",
		Zip:     "000100",
		Region:  "California",
	}
	sh := hotel{
		Name:    "Sky Hotel",
		Address: "Lopez street, 2",
		City:    "LA",
		Zip:     "000101",
		Region:  "California",
	}

	hotels := []hotel{sb, sh}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
