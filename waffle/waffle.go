package waffle

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Conor-Moran/waffle/utils"
)

func parseJson(filePath string) (jsonAsMap map[string]any) {
	jsonAsMap = make(map[string]any)

	data, err := os.ReadFile(filePath)
	utils.IfErrLogFatal(err)

	err = json.Unmarshal(data, &jsonAsMap)
	utils.IfErrLogFatal(err)

	return
}

func parseCutter(filePath string) map[string]string {
	data, err := os.ReadFile(filePath)
	utils.IfErrLogFatal(err)

	return map[string]string{filePath: string(data)}
}

func readFiles[T any](dirPath string, parse func(string) T) (parsed []T) {
	parsed = make([]T, 0)

	entries, err := os.ReadDir(dirPath)
	utils.IfErrLogFatal(err)

	for _, entry := range entries {
		if !entry.IsDir() {
			parsed = append(parsed, parse(utils.FilePath(dirPath, entry)))
		}
	}

	return
}

func readJsonFiles(dirPath string) []map[string]any {
	return readFiles(dirPath, parseJson)
}

func readCutterFiles(dirPath string) []map[string]string {
	return readFiles(dirPath, parseCutter)
}

func Run() {
	recipes := readJsonFiles("./ins/recipes")
	ingredients := readJsonFiles("./ins/ingredients")
	cutters := readCutterFiles("./ins/cutters")

	fmt.Println(recipes, ingredients, cutters)
}
