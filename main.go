package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

// Course stores information about a frontendmasters course
type Course struct {
	Title       string
	URL         string
	Description string
	Instructor  string
}

func main() {
	// Where I will put the json file of frontend masters courses
	// fName := "courses.json"
	// file, err := os.Create(fName)
	// if err != nil {
	// 	log.Fatalf("Cannot create file %q: %s\n", fName, err)
	// 	return
	// }
	// defer file.Close()

	c := colly.NewCollector(colly.AllowedDomains("frontendmasters.com"))

	courses := make([]Course, 0, 200)

	c.OnHTML("div.FM-Course-Item-Content", func(e *colly.HTMLElement) {
		course := Course{
			Title: e.Text,
		}

		// Iterate over div components and add details to course
		e.ForEach("div > .FM-Heading-3 .FM-Link", func(_ int, el *colly.HTMLElement) {
			course.URL = fmt.Sprintf("%s%s", "https:/frontendmasters.com", el.Attr("href"))
		})

		e.ForEach("div > .instructor", func(_ int, el *colly.HTMLElement) {
			course.Instructor = el.Attr("title")
		})

		e.ForEach("div > .description", func(_ int, el *colly.HTMLElement) {
			course.Description = el.Text
		})
		// fmt.Println(course)
		courses = append(courses, course)
		// fmt.Println(courses)
	})
	jsonBytes, err := json.Marshal(courses)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(jsonBytes))
	// // Find and visit all links
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.Attr("href"))
	// })

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })

	c.Visit("https://frontendmasters.com/courses")
}
