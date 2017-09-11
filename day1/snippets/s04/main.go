// BEGINS OMIT
package main

import (
	"fmt"
	"net/url"
)

// Project defines a collaborative, time-bounded effort.
type Project struct {
	Name         string
	Homepage     *url.URL
	Contributors []Contributor
}

// Contributor defines a person contributing to a project.
type Contributor struct {
	ID   int
	Mail string
}

// ENDS OMIT

// BEGINC OMIT
func main() {
	hp, _ := url.Parse("http://example.com")
	aproject := Project{
		Name:     "A simple project",
		Homepage: hp,
		Contributors: []Contributor{
			{
				ID:   29,
				Mail: "ann@example.com",
			},
			{
				ID:   77,
				Mail: "joe@example.com",
			},
		},
	}

	fmt.Printf("Project: %s\nHomepage: %s\nContributors:\n", aproject.Name, aproject.Homepage)
	for _, contributor := range aproject.Contributors {
		fmt.Printf("- %s with ID %d\n", contributor.Mail, contributor.ID)
	}
}

// ENDC OMIT
