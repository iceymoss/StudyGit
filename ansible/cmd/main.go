package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	cmd := exec.Command("top", "-n", "1")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	// 启动命令
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	// 等待命令执行完成
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting for command:", err)
		return
	}

	// 获取输出内容
	outputStr := stdout.String()

	// 解析输出内容，提取所需字段
	rows := strings.Split(outputStr, "\n")
	var data [][]string
	for _, row := range rows {
		// 使用正则表达式提取字段
		re := regexp.MustCompile(`(\d+)\s+(\S+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\S+)\s+(\S+)\s+(\S+)\s+(\S+)\s+(\S+)\s+(.+)`)
		match := re.FindStringSubmatch(row)
		if len(match) == 13 {
			pid := match[1]
			user := match[2]
			virt := match[5]
			res := match[6]
			cpu := match[9]
			mem := match[10]
			time := match[11]
			command := match[12]

			// 将提取的字段组成一个数据行
			data = append(data, []string{pid, user, virt, res, cpu, mem, time, command})
		}
	}

	// 将数据写入 CSV 文件
	file, err := os.Create("top_output.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入 CSV 文件头部
	header := []string{"PID", "USER", "VIRT", "RES", "%CPU", "%MEM", "TIME+", "COMMAND"}
	writer.Write(header)

	// 写入数据
	writer.WriteAll(data)

	fmt.Println("Data written to top_output.csv")
}
