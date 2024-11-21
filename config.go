// -------------------------------------------
// @file      : config.go
// @author    : bo cai
// @contact   : caibo923@gmail.com
// @time      : 2024/11/21 下午1:33
// -------------------------------------------

package config

import (
	"encoding/json"
	"github.com/caibo86/cberrors"
	"gopkg.in/yaml.v2"
	"os"
)

// 全局配置管理器
var manager = &Manager{
	loaded:  make(map[string]bool),
	configs: make(map[string]IConfig),
}

// IConfig 配置接口
type IConfig interface {
	GetType() string
}

// Manager 配置管理器
type Manager struct {
	loaded  map[string]bool
	configs map[string]IConfig
}

// Load 加载全局配置
func Load(filename string, configs ...IConfig) {
	for _, config := range configs {
		manager.addConfig(config)
	}
	loadGlobalConfig(filename)
}

// 添加配置
func (m *Manager) addConfig(c IConfig) {
	if m.configs == nil {
		m.configs = make(map[string]IConfig)
	}
	if _, ok := m.configs[c.GetType()]; ok {
		cberrors.Panic("dup config type:%s", c.GetType())
	}
	m.configs[c.GetType()] = c
	m.loaded[c.GetType()] = false
	return
}

// Get 获取指定配置
func Get(t string) IConfig {
	return manager.configs[t]
}

// String 实现fmt.Stringer接口
func (m *Manager) String() string {
	data, _ := json.Marshal(m.configs)
	return string(data)
}

// 加载全局配置
func loadGlobalConfig(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		cberrors.Panic("load global config file %s error:%v", filename, err)
		return
	}
	content := os.ExpandEnv(string(data))
	if err = parseGlobalConfig(content); err != nil {
		cberrors.Panic("parse global config file %s error:%v", filename, err)
		return
	}
	// 检查配置是否读取完整
	for key, loaded := range manager.loaded {
		if !loaded {
			cberrors.Panic("config not loaded:%v", key)
		}
	}
	return
}

// 解析全局配置
func parseGlobalConfig(content string) error {
	temp := make(map[string]interface{})
	if err := yaml.Unmarshal([]byte(content), temp); err != nil {
		return err
	}
	for _, config := range manager.configs {
		data, err := yaml.Marshal(temp[config.GetType()])
		if err != nil {
			return err
		}
		if err = yaml.Unmarshal(data, config); err != nil {
			return err
		}
		manager.loaded[config.GetType()] = true
	}
	return nil
}
