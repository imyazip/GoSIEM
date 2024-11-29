package utils

import (
	"strconv"

	"github.com/imyazip/GoSIEM/windows_agent/internal/pkg/parser"
)

func StructToSlice(event parser.Win32_NTLogEvent) []string {
	// Создаем слайс с фиксированными полями структуры
	result := []string{
		strconv.Itoa(int(event.EventID)), // Преобразуем int32 в строку
		event.Logfile,                    // Logfile уже строка
		event.Message,                    // Message уже строка
		event.Category,                   // Category уже строка
		event.Source,                     // Source уже строка
	}

	// Добавляем каждый элемент InsertionStrings отдельно
	result = append(result, event.InsertionStrings...)

	return result
}
