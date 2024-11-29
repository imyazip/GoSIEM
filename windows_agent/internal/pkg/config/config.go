package config

import (
	"log" // Для логирования
	"os"  // Для работы с файловой системой

	"gopkg.in/yaml.v3" // Для работы с YAML-файлами
)

// Структура для конфигурации приложения
type Config struct {
	Server struct {
		Host string `yaml:"host"` // Адрес сервера
		Port int    `yaml:"port"` // Порт сервера
	} `yaml:"server"`
	Api struct {
		Key string `yaml:"key"` //API ключ
	} `yaml:"api"`
}

// Функция для загрузки конфигурации из YAML-файла
func LoadConfig(filePath string) *Config {
	// Открываем конфигурационный файл
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err) // Логируем ошибку, если не удалось открыть файл
	}
	defer file.Close()

	config := &Config{}
	// Декодируем YAML-файл в структуру Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		log.Fatalf("Failed to decode config file: %v", err) // Логируем ошибку, если не удалось декодировать
	}
	return config
}
