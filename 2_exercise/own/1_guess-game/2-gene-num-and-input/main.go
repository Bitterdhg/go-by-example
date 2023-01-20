package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	tgtNum := rand.Intn(maxNum)
	// fmt.Println("The target number is ", tgtNum)

	for {
		fmt.Println("Please input your guess")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n') // 这里是'\n'
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again!", err)
			continue
		}
		input = strings.TrimSuffix(input, "\r\n") // 这里是 "\n" "\r"是回车符 "\n"是换行符
		// strings.TrimSuffix 具体用法 : https://vimsky.com/examples/usage/golang_strings_TrimSuffix.html
		// 为什么用 "\r\n" https://zhidao.baidu.com/question/625107249351749292.html

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an interger value")
			continue
		}
		if tgtNum < guess {
			fmt.Println("Your guess is bigger than target number")
		} else if guess < tgtNum {
			fmt.Println("Your guess is smaller than target number")
		} else {
			fmt.Println("Your ans is correct!")
			break
		}
	}

}
