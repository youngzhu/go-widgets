package main

import (
	"bufio"
	"os"
)

/*
实现：有一组名单，轮流陪餐，遇到节假日顺延
*/

const (
	startFrom = "2023-09-11" // 陪餐首次开始的时间

	subject = "陪餐 11:20-12:20" // 邮件标题
)

func main() {

}

func turn() int {
	return 0
}

func readBabies() ([]string, error) {
	f, err := os.Open("babies.txt")
	if err != nil {
		return nil, err
	}

	babies := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		babies = append(babies, scanner.Text())
	}
	return babies, nil
}
