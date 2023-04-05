package scene

import (
	"encoding/json"
	"io"
	"os"
)

type Scene struct {
	Crc    int            `json:"crc"`
	Lines  []string       `json:"lines"`
	Labels map[string]int `json:"labels"`
}

func New() Scene {
	return Scene{}
}

func NewSceneFromFile(path string) Scene {
	fileData, _ := os.Open(path)
	defer fileData.Close()
	bytes, _ := io.ReadAll(fileData)
	var scene Scene
	json.Unmarshal(bytes, &scene)
	return scene
}
