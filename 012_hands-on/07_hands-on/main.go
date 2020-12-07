package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type meal struct {
	Name, Description string
	Price             int
}

type menu struct {
	Breakfast, Lunch, Dinner []meal
}

type restaurant struct {
	Name     string
	MainMenu menu
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	menu1 := menu{
		Breakfast: []meal{
			meal{
				Name:        "English breakfast",
				Description: "Eggs, tomato, cheese, sosages",
				Price:       5,
			},
			meal{
				Name:        "French breakfast",
				Description: "Chocolate, kruassan, jem, tosts",
				Price:       7,
			},
		},
		Lunch: []meal{
			meal{
				Name:        "Soup",
				Description: "Chicken, potatos, ...",
				Price:       15,
			},
			meal{
				Name:        "Kebab",
				Description: "Beaf, vegetables, ...",
				Price:       20,
			},
		},
		Dinner: []meal{
			meal{
				Name:        "Pasta",
				Description: "Chicken, noods, ...",
				Price:       30,
			},
			meal{
				Name:        "Mussaka",
				Description: "Beaf, rice, ...",
				Price:       40,
			},
		},
	}

	restaurant1 := restaurant{
		Name:     "The Best",
		MainMenu: menu1,
	}
	listOfRestaurant := []restaurant{restaurant1}
	err := tpl.Execute(os.Stdout, listOfRestaurant)
	if err != nil {
		log.Fatalln(err)
	}
}
