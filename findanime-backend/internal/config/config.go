package config

import (
	"io"
	"os"

	"github.com/goccy/go-json"
)

var (
	Configuration Config
)

func init() {
	file, err := os.Open("config.json")

	if err != nil {
		panic(err)
	}

	bytesValue, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(bytesValue, &Configuration)
	defer file.Close()
}
