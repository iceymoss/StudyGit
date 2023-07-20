package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	//cmd := exec.Command("docker", "run", "-d", "-p", "8030:8030", "mgr9525/gokins:latest")
	//out, err := cmd.CombinedOutput()
	//if err != nil {
	//	log.Fatal("fail:", err)
	//	return
	//}
	res := make([]string, 0)
	//res = append(res, string(out))

	cmd := exec.Command("docker", "ps")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("fail:", err)
		return
	}
	res = append(res, string(out))

	cmd = exec.Command("docker", "ps")
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal("fail:", err)
		return
	}
	res = append(res, string(out))

	for _, v := range res {
		fmt.Println("log:", v)
	}
}
