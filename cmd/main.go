package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// @title sahtian API
// @version 1.0
// @description This  is a sahtian API server.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host api.sahtian.com
// @BasePath /api/v1
func main() {
	if err := Execute(); err != nil {
		outputInitError(err)
	}
}

func outputInitError(err error) {
	fmt.Println(`{"error":"` + err.Error() + `"}`)

	os.Exit(1)
}

func waitSignalExit(cancel func()) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-ch

	cancel()
}
