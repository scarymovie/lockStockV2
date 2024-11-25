package logger

import (
	"log"
	"os"
)

func SetupLogging(logFilePath string) error {
	// Открываем файл для записи логов
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// Перенаправляем стандартный лог в файл
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile) // Формат логов: время + имя файла:строка

	return nil
}
