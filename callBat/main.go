package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func runBat(f string) {
	fmt.Println("f =>\t", &f)
	out, err := exec.Command(f).CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func runComm() {
	out, err := exec.Command("aws", "s3", "ls", "--profile", "taiwan-prod", "s3://prod-fusion-s3-projects-ap-southeast-2/taiwan/inputs/CHT/WS_Tx3/").CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	fmt.Println(strings.Contains(string(out), "20220427.tx3"))
}

func main() {
	var f string = "run.bat"
	fmt.Println("f =>\t", &f)
	runBat(f)
	runComm()
}
