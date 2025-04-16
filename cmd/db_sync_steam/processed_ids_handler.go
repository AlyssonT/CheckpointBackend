package main

import (
	"fmt"
	"os"
	"strings"
)

const filePath = "processed_ids.txt"

func ReadProcessedIds() map[string]bool {
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]bool)
		}
		fmt.Println("Error on reading file", err)
		return nil
	}

	ids := strings.Split(string(data), ",")
	idMap := make(map[string]bool)
	for _, id := range ids {
		if id != "" {
			idMap[id] = true
		}
	}
	return idMap
}

func SaveProcessedIds(idMap map[string]bool) {
	var idList []string
	for id := range idMap {
		idList = append(idList, id)
	}
	data := strings.Join(idList, ",")

	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		fmt.Println("Error on saving file", err)
	}
}
