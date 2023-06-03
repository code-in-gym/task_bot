package config

import (
	"fmt"
	"log"
	"os"

	"code-in-gym/task_bot/task"
	"code-in-gym/task_bot/notify"

	"github.com/BurntSushi/toml"
)

type Config[T any] struct {
	Task   map[string]task.Task[T]
	Notify notify.Config
}

func NewConfig[T any]() *Config[T] {
	return &Config[T]{}
}

// Load load config from file,
// DO NOT put any sensitive key or secret in it,
// use GitHub Actions secrets.
func (c *Config[T]) Load(file string) {
	if _, err := os.Stat(file); err != nil {
		log.Fatalf("Config file `%s` DOES NOT exist.", file)
	}

	_, err := toml.DecodeFile(file, &c)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
