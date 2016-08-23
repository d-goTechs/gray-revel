package main

import (
	"os/exec"
	"os"
)

func main() {
	var c = exec.Command("ping","www.baidu.com")
	c.Stdout = os.Stdout
	c.Run()
}
