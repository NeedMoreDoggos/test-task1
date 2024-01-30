package main

import (
	"fmt"

	"github.com/NeedMoreDoggos/test-task1/internal/config"
)

func main() {
	//config load
	cfg := config.MustConfig()
	fmt.Printf("config: %+v\n", cfg)
	fmt.Printf("meta: %+v\n", cfg.GetMeta())

	//logger

	//app init

	//app start

	//sys signals handling
}
