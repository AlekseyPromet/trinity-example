package main

import (
	"log"
	"os"
	"os/signal"

	"context"

	"github.com/AlekseyPromet/trinity-example/orders/server"
)

func main() {
	app := server.NewServer()

	ctx, cancel := context.WithCancel(context.Background())
	sigchan := make(chan os.Signal)
	defer close(sigchan)

	go func() {
		signal.Notify(sigchan, os.Interrupt)
		sigint := <-sigchan
		if sigint != nil {
			cancel()
		}
	}()

	if err := app.Run(ctx); err != nil {
		log.Fatalln(err)
	}
}
