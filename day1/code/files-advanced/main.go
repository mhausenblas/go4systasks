package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
)

// BEGIND OMIT

// Project defines a collaborative, time-bounded effort.
type Project struct {
	Name         string        `json:"name"`
	Homepage     *url.URL      `json:"homepage"`
	Contributors []Contributor `json:"members"`
}

// Contributor defines a person contributing to a project.
type Contributor struct {
	ID   int    `json:"id"`
	Mail string `json:"mail"`
}

// ENDD OMIT

func main() {

	// BEGINC OMIT
	p := Project{}
	raw, err := ioutil.ReadFile("../../data/project.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_ = json.Unmarshal(raw, &p)
	// ENDC OMIT
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
}
