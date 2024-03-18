	```go 
    // Extract details of the course
	detailCollector.OnHTML(`div[id=rendered-content]`, func(e *colly.HTMLElement) {
		log.Println("Course found", e.Request.URL)
		title := e.ChildText(".banner-title")
		if title == "" {
			log.Println("No title found", e.Request.URL)
		}
		course := Course{
			Title:       title,
			URL:         e.Request.URL.String(),
			Description: e.ChildText("div.content"),
			Creator:     e.ChildText("li.banner-instructor-info > a > div > div > span"),
			Rating:      e.ChildText("span.number-rating"),
		}
		// Iterate over div components and add details to course
		e.ForEach(".AboutCourse .ProductGlance > div", func(_ int, el *colly.HTMLElement) {
			svgTitle := strings.Split(el.ChildText("div:nth-child(1) svg title"), " ")
			lastWord := svgTitle[len(svgTitle)-1]
			switch lastWord {
			// svg Title: Available Languages
			case "languages":
				course.Language = el.ChildText("div:nth-child(2) > div:nth-child(1)")
			// svg Title: Mixed/Beginner/Intermediate/Advanced Level
			case "Level":
				course.Level = el.ChildText("div:nth-child(2) > div:nth-child(1)")
			// svg Title: Hours to complete
			case "complete":
				course.Commitment = el.ChildText("div:nth-child(2) > div:nth-child(1)")
			}
		})
		courses = append(courses, course)
	})
