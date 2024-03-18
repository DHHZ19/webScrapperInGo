package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

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
	fName := "courses.json"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	// Create a channel to signal when the goroutine is done
	done := make(chan struct{})

	c := colly.NewCollector(
		colly.AllowedDomains("frontendmasters.com"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./courses_cache"),
	)

	courses := make([]Course, 0, 200)

	c.OnHTML("div.FM-Course-Item-Content", func(e *colly.HTMLElement) {
		var course Course

		// Iterate over div components and add details to course
		e.ForEach("div > .FM-Heading-3 .FM-Link", func(_ int, el *colly.HTMLElement) {
			course.Title = el.Text
			course.URL = fmt.Sprintf("%s%s", "https:/frontendmasters.com", el.Attr("href"))
		})

		e.ForEach("div > .instructor", func(_ int, el *colly.HTMLElement) {
			course.Instructor = el.Attr("title")
		})

		e.ForEach("div > .description", func(_ int, el *colly.HTMLElement) {
			course.Description = el.Text
		})
		courses = append(courses, course)
	})
	// Start the goroutine
	go func() {
		defer close(done) // Signal that the goroutine is done
		c.Visit("https://frontendmasters.com/courses")
	}()

	// Wait for the goroutine to finish
	<-done

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	enc.Encode(courses)
}
