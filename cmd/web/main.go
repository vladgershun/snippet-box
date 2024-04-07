package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {

	port := flag.String("port", ":2020", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	app := &application{logger: logger}

	logger.Info("starting server", slog.String("port", *port))

	err := http.ListenAndServe(*port, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
