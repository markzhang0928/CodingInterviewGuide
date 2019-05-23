package main

import "fmt"

var sli = []int{1, 2, 3, 2, 3, 4, 5, 5}

func GetLength(arr1 []int) int {

	if len(arr1)<=0 {
		return 0
	}
	n := len(arr1)
	var length, lengthB = 0, 0
	for i:=0; i<n;i++{
		// 全等数列
		// 1 2 3 3
		for j:=i+1; j<=n-1; j++{

			if arr1[i]==arr1[j] {
				length += 1
			} else if arr1[i] != arr1[j]{
				length = 1
				break
			}
		}
		length = length + 1
		fmt.Println("get 全等数列", length)

	}

	// 相等子串
	// 1 2 3 1 2 3
	tmp := make(map[int]int)
	for i:=0; i<n; {
		// 1 2 3 2 3 4 5 5
		// i = 0 arr1[0]=1
		// j =
		for j:=i+1; j<n-1;{
			if arr1[i]==arr1[j] && len(tmp)<=2 {
				i++
				j++
				lengthB += 1
				tmp[arr1[i]] = 1
			} else if arr1[i] != arr1[j]{
				j++
				lengthB = lengthB
				break
			} else if len(tmp) > 2 {
				for k := range tmp {
					delete(tmp, k)
				}
				lengthB = 0
				break
			}
		}
		lengthB =lengthB
		fmt.Println("get相等子串", lengthB)
	}

	if length >= lengthB{
		return length
	} else {
		return lengthB
	}
}

func main(){

	res := GetLength(sli)
	fmt.Println(res)
}
