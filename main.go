package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event: ", event)
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					log.Println("rename : ", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					log.Println("error: ", err)
					return
				}
			}
		}
	}()
	err = watcher.Add("/tmp/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
