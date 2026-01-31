package util

import (
	"os"
)

// Load from a YAML file. Return byte content.
// 从 YAML 文件加载。返回 byte 内容。
func LoadYaml(path string) ([]byte, error) {
	return os.ReadFile(path)
}
