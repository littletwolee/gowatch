package main

import (
	"fmt"
)

func main() {
	// m := newMonitor()
	// m.run()
	// m.monitoring()
	p := &paths{m: make(map[string]int64)}
	path, err := getPath()
	if err != nil {
		fmt.Println(err)
	}
	p.root = path
	if err := p.fetch(); err != nil {
		fmt.Println(err)
	}
	for k, v := range p.m {
		fmt.Printf("%s : %d \n", k, v)
	}
}
