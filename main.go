package main

import (
	"context"
	"log"
	"search/pkg/search"
	"time"
)

func main() {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	files := []string{"I am Kavrakov Shavqat \n I  go to want Backent whera Amount\n", " I am ver good\n I want go to America", "I am very want"}
	for i := range search.All(ctx, "want", files) {
		log.Print(i)
	}
	<-time.After(time.Millisecond)
	cancel()

}
