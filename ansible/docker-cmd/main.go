package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	//ansible -i hosts.ini all --become-user ubuntu -m shell -a 'ls'
	cmd := exec.Command("ansible", "-i", "/Users/iceymoss/ansible/hosts.ini", "all", "--become-user", "ubuntu", "-m", "shell", "-a", "go run /home/ubuntu/iceymoss/ansible/main.go difidhfidf")
	cmd.Env = append(os.Environ(), "PATH=/opt/homebrew/bin/ansible")
	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	fmt.Println(cmd.String())
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing docker ps command:", err)
		fmt.Println("Stderr:", stderr.String())
		return
	}

	fmt.Println("返回：", stdout.String())

	//// 获取输出内容
	//outputStr, err := cmd.CombinedOutput()
	//if err != nil {
	//	log.Println("err:", err)
	//	return
	//}

	// 解析输出内容，按行分隔
	//rows := strings.Split(string(outputStr), "\n")
	//
	//for _, r := range rows {
	//	fmt.Println(r)
	//}
	// 写入 CSV 文件头部
	//header := []string{"CONTAINER ID", "IMAGE", "COMMAND", "CREATED", "STATUS", "PORTS", "NAMES"}
	//writer.Write(header)
}
