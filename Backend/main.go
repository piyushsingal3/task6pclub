package main

import (
	"attendance-app/server"

	"attendance-app/store"

	"sync"
)

func main() {
	//A wait group is created to wait for goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1)

	//creates a new mongostore
	mongoStore := store.NewMongoStore()

	//starts the server in  a goroutine
	go func() {
		defer wg.Done()                  //decrease the waitgroup counter when it completes
		server.Performserver(mongoStore) //runs the server with mongodb store
	}()

	//it waites for server go routine to finish
	wg.Wait()
	//it closes mongodb store
	mongoStore.Close()

}
