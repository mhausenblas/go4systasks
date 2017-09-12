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

// BEGIND OMIT

func mailof(uid int, p Project) string {
	for _, contributor := range p.Contributors {
		if contributor.ID == uid {
			return contributor.Mail
		}
	}
	return ""
}

// ENDD OMIT

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
	// BEGINC OMIT
	uid := 29
	mail := mailof(uid, aproject)
	fmt.Printf("Mail address of contributor with ID %d is:\n%s", uid, mail)
	// ENDC OMIT
}
