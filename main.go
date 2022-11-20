package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Conor-Moran/waffle/utils"
)

func main() {
	myMap := make(map[string]any)
	data, err := os.ReadFile("./ins/ingredients/x.json")

	utils.IfErrLogFatal(err)

	err = json.Unmarshal(data, &myMap)
	utils.IfErrLogFatal(err)
	fmt.Printf("%v", myMap)

	entries, err := os.ReadDir("./ins/ingredients")
	utils.IfErrLogFatal(err)
	for _, entry := range entries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
		}
	}
}
