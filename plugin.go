package main

import (
	`github.com/dronestock/drone`
)

type plugin struct {
	config *config
}

func newPlugin() drone.Plugin {
	return &plugin{
		config: &config{},
	}
}

func (p *plugin) Configuration() drone.Configuration {
	return p.config
}

func (p *plugin) Steps() []*drone.Step {
	return []*drone.Step{
		drone.NewStep(p.tidy, drone.Name(`清理依赖`)),
		drone.NewStep(p.lint, drone.Name(`代码静态检查`)),
		drone.NewStep(p.test, drone.Name(`测试`)),
		drone.NewStep(p.build, drone.Name(`编译`)),
		drone.NewStep(p.upx, drone.Name(`压缩`)),
	}
}
