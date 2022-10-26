package main

func (p *plugin) Scheme() string {
	return `https://raw.githubusercontent.com/pangum/drone/master/scheme.json`
}

func (p *plugin) Card() (card any, err error) {
	return
}
