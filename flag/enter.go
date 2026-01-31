package flag

import "flag"

type Option struct {
	ConfigPath string
}

// Parse parses command-line flags and returns an Option struct. -c is used to specify the config file path.
// 解析命令行参数并返回一个 Option 结构体。-c 用于指定配置文件路径。
func Parse() Option {
	var opt Option
	flag.StringVar(&opt.ConfigPath, "c", "conf/config.yaml", "path to config yaml")
	flag.Parse()
	return opt
}
