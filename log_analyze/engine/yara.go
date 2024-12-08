package engine

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/hillu/go-yara"
)

func loadRules(directory string) (*yara.Rules, error) {
	compiler, err := yara.NewCompiler()
	if err != nil {
		return nil, fmt.Errorf("Ошибка создания компилятора YARA: %v", err)
	}

	err = filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("Ошибка при обходе директории: %v", err)
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(d.Name()) == ".yar" {
			log.Printf("Загрузка правил из файла: %s", path)
			file, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("Ошибка открытия файла %s: %v", path, err)
			}
			defer file.Close()

			if err := compiler.AddFile(file, ""); err != nil {
				return fmt.Errorf("Ошибка добавления правил из файла %s: %v", path, err)
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	rules, err := compiler.GetRules()
	if err != nil {
		return nil, fmt.Errorf("Ошибка компиляции правил: %v", err)
	}

	return rules, nil
}

func scanLogsBySource(rules *yara.Rules, logSource string, data []byte) ([]yara.MatchRule, error) {
	scanner, err := yara.NewScanner(rules)
	if err != nil {
		return nil, fmt.Errorf("Ошибка создания сканера: %v", err)
	}

	return matches, nil
}
