package waffle

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/Conor-Moran/waffle/utils"
)

type Recipe map[string]string

type JsonAsMap map[string]any

type Cutter string

func parseJson(filePath string) (jsonAsMap JsonAsMap) {
	jsonAsMap = linkedhashmap.new(type)

	parseJsonX(filePath, &jsonAsMap)
	return
}

func parseRecipe(filePath string) (recipe Recipe) {
	recipe = make(Recipe)
	parseJsonX(filePath, &recipe)
	return
}

func parseJsonX[T Recipe | JsonAsMap](filePath string, jsonMap *T) {
	data, err := os.ReadFile(filePath)
	utils.IfErrLogFatal(err)

	err = json.Unmarshal(data, jsonMap)
	utils.IfErrLogFatal(err)
}

func parseCutter(filePath string) Cutter {
	data, err := os.ReadFile(filePath)
	utils.IfErrLogFatal(err)
	return Cutter(data)
}

func readFiles[T any](dirPath string, parse func(string) T) (parsed map[string]T) {
	parsed = map[string]T{}

	entries, err := os.ReadDir(dirPath)
	utils.IfErrLogFatal(err)

	for _, entry := range entries {
		if !entry.IsDir() {
			filePath := utils.FilePath(dirPath, entry)
			parsed[filePath] = parse(filePath)
		}
	}

	return
}

func readRecipes(dirPath string) map[string]Recipe {
	return readFiles(dirPath, parseRecipe)
}

func readJsonFiles(dirPath string) map[string]JsonAsMap {
	return readFiles(dirPath, parseJson)
}

func readCutterFiles(dirPath string) map[string]Cutter {
	return readFiles(dirPath, parseCutter)
}

// Run the Waffle
func Run() {
	recipes := readRecipes("./ins/recipes")
	ingredients := readJsonFiles("./ins/ingredients")
	cutters := readCutterFiles("./ins/cutters")

	for _, x := range recipes {
		ingredientsKey := x["ingredients"]
		context := ingredients[ingredientsKey]

		for key, value := range context {
			fmt.Printf("\n%s -- %v", key, reflect.TypeOf(value))
		}

		fmt.Printf("\nXXX: %v\n", context)
	}

	fmt.Println(cutters)
}
