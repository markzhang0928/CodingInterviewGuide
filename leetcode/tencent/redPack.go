package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//10个人 抢10000分  也就是10个人抢100块钱
	count, amount := int64(10), int64(10000)
	remain := amount
	sum := int64(0)
	for i := int64(0); i < count; i++ {
		x := DoubleAverage(count-i, remain)
		remain -= x
		sum += x
		fmt.Println(i+1, "=", float64(x)/float64(100), ", ")
	}
	fmt.Println()
	fmt.Println("总和是:", sum)
}

// 提前定义能抢到的最小金额1分
var min int64 = 1

// 二倍均值算法
func DoubleAverage(count, amount int64) int64 {
	if count == 1 {
		return amount
	}
	// 计算出最大可用金额
	maxNum := amount - min*count
	// 计算出最大可用平均值
	avg := maxNum / count
	avg2 := 2*avg + min
	// 随机红包金额序列元素，把二倍均值作为随机的最大数
	rand.Seed(time.Now().UnixNano())
	x := rand.Int63n(avg2) + min
	return x
}
