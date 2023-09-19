package main

import (
	"fmt"

	"chamger.org/firsttry/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	//TO DO: init logger: slog

	//TO DO: init storage: sqlite

	//TO DO: init router: chi, "chi ren"

	//TO DO: run server
}
