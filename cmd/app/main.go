package main

import (
	"lockStockV2/external/logger"
	config2 "lockStockV2/internal/config"
	"lockStockV2/internal/container" // Инициализация зависимостей
	"log"
)

func main() {
	if err := logger.SetupLogging("./logs/app.log"); err != nil {
		log.Fatalf("could not set up logging: %v", err)
	}

	log.Println("Application started")

	// 1. Загрузка конфигурации
	config, err := config2.LoadConfig("./internal/config/config.dev.yaml")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// 2. Инициализация контейнера зависимостей
	c := container.Initialize(config)
	if c == nil {
		log.Fatalf("could not initialize container")
	}

	// 3. Настройка серверов
	//httpServer := network.NewHTTPServer(c, config.AppPort)
	//
	//// 4. Управление жизненным циклом приложения
	//runServers(httpServer)
}
