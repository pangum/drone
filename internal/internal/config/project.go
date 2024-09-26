package config

type Project struct {
	// 源文件目录
	Source string `default:"${SOURCE=.}" json:"source,omitempty"`
	// 输出目录
	Dir string `default:"${DIR=.}" json:"dir,omitempty"`
	// 私有库
	Privates []string `default:"${PRIVATES}" json:"privates,omitempty"`
	// 环境变量
	Environments map[string]string `default:"${ENVIRONMENTS=${ENVS}}" json:"environments,omitempty"`
}
