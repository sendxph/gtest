package main

import (
	"fmt"
	"os/exec"
)

func runBat(f string) {
	fmt.Println("f =>\t", &f)
	out, err := exec.Command(f).CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func main() {
	var bat string = "run.bat"
	fmt.Println("bat =>\t", &bat)
	runBat(bat)
}
