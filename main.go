package main

import (
	"context"
	"log"
	"search/pkg/search"
	"time"
)

func main() {
	root := context.Background()
	ctx, _ := context.WithTimeout(root, time.Second)
	files := []string{"I am Kavrakov Shavqat \n I  go to Backent whera Amount\n I am ver good\n I want go to America", "I am very want"}
	go func() {
		for i := range search.All(ctx, "want", files) {
			log.Print(i)
		}
	}()
	<-time.After(time.Second * 2)
	//cancel()

}