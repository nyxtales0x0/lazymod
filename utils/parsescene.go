package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Scene struct {
	Crc    int            `json:"crc"`
	Lines  []string       `json:"lines"`
	Labels map[string]int `json:"labels"`
}

func ParseSceneFile(filePath string) Scene {
	fileContent, _ := os.Open(filePath)
	defer fileContent.Close()
	byteData, _ := ioutil.ReadAll(fileContent)
	var scene Scene
	json.Unmarshal(byteData, &scene)
	return scene
}
