package main

import (
	"github.com/LuizEduardo-service/go_oportunidade/config"
	"github.com/LuizEduardo-service/go_oportunidade/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// inicialize configs
	err := config.Init()

	if err != nil {
		logger.Errorf("Erro na configuração de iniciaização: %v", err)
		return

	}
	router.Initialize()
}
