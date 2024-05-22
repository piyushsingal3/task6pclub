package main

import (
	"attendance-app/server"

	"attendance-app/store"

	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	mongoStore := store.NewMongoStore()

	go func() {
		defer wg.Done()
		server.Performserver(mongoStore)
	}()
	wg.Wait()
	mongoStore.Close()

}
