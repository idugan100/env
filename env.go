/*
This package allows use to parse env files. By default it will parse the file ./.env, but you can also set a custom path to the env file. Values must be seperated by an equals sign and each value must be on a new line.
*/
package env

import (
	"fmt"
	"os"
	"strings"
)

/* This is the env parsing struct. It contains a FilePath variable which is the location of the env file you want to parse and a map of values that were parsed*/
type Env struct {
	filePath string
	values   map[string]string
}

/* The function will return a pointer to a new Env struct. The default path is ./.env and the default Values map is empty*/
func NewEnv() *Env {
	return &Env{filePath: "./.env", values: make(map[string]string)}
}

/* This function allows you to set a path to a custom .env file besides the default ./.env */
func (e *Env) SetEnvPath(path string) {
	e.filePath = path
}

/* This function performs that parsing of your env file. If you want to set the path to the env file please use SetEnvPath before calling this function. */
func (e *Env) ParseEnv() error {

	contents, err := os.ReadFile(e.filePath)
	if err != nil {
		return fmt.Errorf("error reading file %w", err)
	}

	stringContents := string(contents)
	lineArray := strings.Split(stringContents, "\n")

	for _, line := range lineArray {
		if strings.Contains(line, "=") {
			res := strings.SplitN(line, "=", 2)
			if len(res) == 2 {
				e.values[res[0]] = res[1]
			}
		}
	}
	return nil
}

/* This function is how you access the parsed env values. If the value has not been set an empty string will be returned. */
func (e *Env) GetEnvValue(key string) string {
	return e.values[key]
}
