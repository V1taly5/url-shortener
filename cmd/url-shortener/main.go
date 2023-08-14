package main

import (
	"url-shortener/internal/config"
)

func main() {
	// TODO: init config: cleanenv
	config := config.MustLoad()

	// TODO: init logger: slog

	// TODO: init storage: psql

	// TODO: init router: chi

	// TODO: run server
}
