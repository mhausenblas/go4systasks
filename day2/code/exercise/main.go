package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

func main() {
	w, err := fsnotify.NewWatcher()
	watcher = w
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = watcher.Close()
	}()
	limit := 3
	done := make(chan bool)
	fileq := make(chan string, limit)
	go queue(fileq)
	go warn(limit, fileq)
	err = watcher.Add(".")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func queue(fileq chan string) {
	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Create == fsnotify.Create {
				fileq <- event.Name
			}
		case errw := <-watcher.Errors:
			log.Println("error:", errw)
		}
	}
}

func warn(limit int, fileq chan string) {
	overall := 0
	for {
		<-fileq
		overall++
		if overall >= limit {
			log.Printf("Warning, there are more than %d files in this directory", limit)
		}
	}
}
