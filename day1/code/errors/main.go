package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"
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

func mailof(uid int, p Project) (string, error) {
	for _, contributor := range p.Contributors {
		if contributor.ID == uid {
			return contributor.Mail, nil
		}
	}
	return "", errors.New("No contributor found")
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
	uid := 1
	mail, err := mailof(uid, aproject)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Mail address of contributor with ID %d is:\n%s", uid, mail)
	// ENDC OMIT
}
