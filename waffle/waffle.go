package waffle

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Conor-Moran/waffle/utils"
)

func ReadJson(filePath string) (jsonAsMap map[string]any) {
	jsonAsMap = make(map[string]any)

	data, err := os.ReadFile(filePath)
	utils.IfErrLogFatal(err)

	err = json.Unmarshal(data, &jsonAsMap)
	utils.IfErrLogFatal(err)

	return
}

func ReadCutter(filePath string) map[string]string {
	data, err := os.ReadFile(filePath)
	utils.IfErrLogFatal(err)

	return map[string]string{filePath: string(data)}
}

func Read[T any](dirPath string, parser func(string) T) (parsed []T) {
	parsed = make([]T, 0)

	entries, err := os.ReadDir(dirPath)
	utils.IfErrLogFatal(err)

	for _, entry := range entries {
		if !entry.IsDir() {
			parsed = append(parsed, parser(utils.FilePath(dirPath, entry)))
		}
	}

	return
}

func ReadJsonFiles(dirPath string) []map[string]any {
	return Read(dirPath, ReadJson)
}

func ReadCutterFiles(dirPath string) []map[string]string {
	return Read(dirPath, ReadCutter)
}

func Run() {
	recipes := ReadJsonFiles("./ins/recipes")
	ingredients := ReadJsonFiles("./ins/ingredients")
	cutters := ReadCutterFiles("./ins/cutters")

	fmt.Println(recipes, ingredients, cutters)
}
