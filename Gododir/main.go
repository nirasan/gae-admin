package main

import (
	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	do.Env = `GOPATH=.vendor::$GOPATH`

	p.Task("bindata", nil, func(c *do.Context) {
		c.Run(`go-bindata -o bindata/bindata.go -pkg "bindata" -prefix "templates/" templates/...`)
	}).Src("templates/**/*")
}

func main() {
	do.Godo(tasks)
}