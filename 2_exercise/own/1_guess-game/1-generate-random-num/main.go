package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano()) // 设定好随机种子，tgtNum才会变化，否则会输出一个固定值
	tgtNum := rand.Intn(maxNum)
	fmt.Printf("The tgt Number is %v \n", tgtNum)
}
