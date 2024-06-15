package main

import (
	"flag"
	"log/slog"
	"os"
)

type config struct {
	addr string
}

var Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	Level:     slog.LevelDebug,
	AddSource: true,
}))

var cfg config

func init() {
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
}

// Run - responsible for the instantiation and startup
func Run() error {
	Logger.Info("Starting up API on", "addr", cfg.addr)
	return nil
}

func main() {
	Logger.Info("Welcome!")
	if err := Run(); err != nil {
		Logger.Error("Error", err)
	}
}
