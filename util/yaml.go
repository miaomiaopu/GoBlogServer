package util

import (
	"os"
)

// 从 YAML 文件加载。返回 byte 内容。
func LoadYaml(path string) ([]byte, error) {
	return os.ReadFile(path)
}
