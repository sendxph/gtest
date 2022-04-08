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
	var f string = "run.bat"
	fmt.Println("f =>\t", &f)
	runBat(f)
}
