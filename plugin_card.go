package main

func (p *core.Plugin) Scheme() string {
	return `https://raw.githubusercontent.com/pangum/drone/master/scheme.json`
}

func (p *core.Plugin) Card() (card any, err error) {
	return
}
