package main

import (
	"context"
	"log"

	"github.com/bozkurtemre/go-service-weaver-template/internal/frontend"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	if err := weaver.Run(context.Background(), frontend.Serve); err != nil {
		log.Fatal(err)
	}
}
