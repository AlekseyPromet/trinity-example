package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"context"

	"github.com/AlekseyPromet/trinity-example/orders/cmd/api"
	"github.com/AlekseyPromet/trinity-example/orders/server"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {

	srv := server.NewService()

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

	monitoring(ctx)

	if err := api.NewHTTPServer(srv.GetHttpAddress(), srv); err != nil {
		log.Fatalln(err)
	}
}

func monitoring(ctx context.Context) {
	go func() {
		v, err := mem.VirtualMemory()
		if err != nil {
			return
		}

		for {
			// almost every return value is a struct
			fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
			time.Sleep(60 * time.Second)

			select {
			case <-ctx.Done():
				return
			}
		}
	}()
}
