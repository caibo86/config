// -------------------------------------------
// @file      : main.go
// @author    : bo cai
// @contact   : caibo923@gmail.com
// @time      : 2024/11/21 下午3:25
// -------------------------------------------

package main

import (
	"fmt"
	"github.com/caibo86/config"
	"os"
)

type Example struct {
	Level   int32  `yaml:"level"`
	Size    int32  `yaml:"size"`
	Async   bool   `yaml:"async"`
	LogFile string `yaml:"logFile"`
}

// GetType implements IConfig
func (e *Example) GetType() string {
	return "example"
}

func main() {
	_ = os.Setenv("SERVER_ID", "10001")
	config.Load("example.yaml", &Example{})
	example := config.Get("example").(*Example)
	fmt.Printf("%+v\n", *example)
}
