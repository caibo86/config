# config
___
一个简易的yaml文件读取及配置管理工具

## 主要功能
___
- yaml文件读取
- 支持自定义配置结构体，只需要实现IConfig接口
- 支持多配置结构体按需注册
- 支持yaml文件引用环境变量

## Example
___
```go
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
```

**Output**:
```shell
{Level:1 Size:100 Async:true LogFile:example_10001.log}
```