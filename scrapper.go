package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type Dictionary map[string]string

// type RecipeSpecs struct {
// 	difficulty, prepTime, cookTime, totalTime, servings string
// }

type Recipe struct {
	url, name      string
	ingredients    []string
	// specifications RecipeSpecs
}

func main() {
	args := os.Args[1:]
	url := args[0]
	collector := colly.NewCollector()

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Response", r.StatusCode)
	})

	collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error", err)
	})

	var recipes []Recipe

	collector.OnHTML("main", func(main *colly.HTMLElement) {
		recipe := Recipe{}
		ingredients_dict := Dictionary{}

		recipe.url = url

		recipe.name = main.ChildText("h1")

		ingredients_dict["ingredients"] = main.ChildText("div.ingredients")

		recipe.ingredients = append(recipe.ingredients, ingredients_dict["ingredients"])

		recipes = append(recipes, recipe)
	})

	err := collector.Visit(url)

	if err != nil {
		fmt.Println("Collector Error", err)
	}

}
