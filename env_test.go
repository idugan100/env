package env

import (
	"testing"
)

func TestDefaultFileUsage(t *testing.T) {
	e := NewEnv()
	e.ParseEnv()
	if e.GetEnvValue("Hi") != "earth" {
		t.Errorf("unexpected env value for Hi. Expected 'earth' recieved: %s", e.GetEnvValue("Hi"))
	}

	if e.GetEnvValue("Hello") != "world" {
		t.Errorf("unexpected env value for Hello. Expected 'world' recieved: %s", e.GetEnvValue("Hello"))
	}
}

func TestDefaultFileMissingValueUsage(t *testing.T) {
	e := NewEnv()
	e.ParseEnv()
	if e.GetEnvValue("Missing") != "" {
		t.Errorf("unexpected env value for Missing. Expected '' recieved: %s", e.GetEnvValue("Hi"))
	}
}
func TestCustomFileUsage(t *testing.T) {
	e := NewEnv()
	e.SetEnvPath("./.customenv")
	e.ParseEnv()
	if e.GetEnvValue("Hi") != "earth" {
		t.Errorf("unexpected env value for Hi. Expected 'earth' recieved: %s", e.GetEnvValue("Hi"))
	}

	if e.GetEnvValue("Hello") != "world" {
		t.Errorf("unexpected env value for Hello. Expected 'world' recieved: %s", e.GetEnvValue("Hello"))
	}

	if e.GetEnvValue("Test") != "=" {
		t.Errorf("unexpected env value for Hello. Expected 'world' recieved: %s", e.GetEnvValue("Hello"))
	}
}

func TestCustomFileMissingValueUsage(t *testing.T) {
	e := NewEnv()
	e.SetEnvPath("./.customenv")
	e.ParseEnv()
	if e.GetEnvValue("Missing") != "" {
		t.Errorf("unexpected env value for Missing. Expected '' recieved: %s", e.GetEnvValue("Hi"))
	}
}
