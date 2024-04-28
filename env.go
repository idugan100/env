package env

import (
	"fmt"
	"os"
	"strings"
)

type Env struct {
	FilePath string
	Values   map[string]string
}

func NewEnv() *Env {
	return &Env{FilePath: "./.env", Values: make(map[string]string)}
}

func (e *Env) SetEnvPath(path string) {
	e.FilePath = path
}

func (e *Env) ParseEnv() error {

	contents, err := os.ReadFile(e.FilePath)
	if err != nil {
		return fmt.Errorf("error reading file %w", err)
	}

	stringContents := string(contents)
	lineArray := strings.Split(stringContents, "\n")

	for _, line := range lineArray {
		if strings.Contains(line, "=") {
			res := strings.SplitN(line, "=", 2)
			if len(res) == 2 {
				e.Values[res[0]] = res[1]
			}
		}
	}
	return nil
}

func (e *Env) GetEnvValue(key string) string {
	return e.Values[key]
}
