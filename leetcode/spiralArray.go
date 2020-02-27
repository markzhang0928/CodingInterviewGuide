package main

// pdd third
import "fmt"

func main() {
	//s := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	//s := [][]int{{1}, {2}, {3}, {4}}
	//s := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	s := [][]int{{1, 2}, {5, 6}, {9, 10}, {13, 14}}

	printMatrix(s)
}
func printMatrix(s [][]int) {
	if s == nil {
		fmt.Println("切片为空，无法打印")
	}
	bex, bey := 0, 0
	hang := len(s) - 1
	lie := len(s[0]) - 1
	if hang == 0 {
		for _, v := range s[0] {
			fmt.Println(v)
		}
		return
	}
	if lie == 0 {
		for _, v := range s {
			fmt.Println(v[0])
		}
		return
	}
	for bex <= hang && bey <= lie {
		ax, ay := bex, bey
		for ay < lie {
			fmt.Println(s[ax][ay])
			ay++
		}
		for ax < hang {
			fmt.Println(s[ax][ay])
			ax++
		}
		for ay > bey {
			fmt.Println(s[ax][ay])
			ay--
		}
		for ax > bex {
			fmt.Println(s[ax][ay])
			ax--
		}
		bex++
		bey++
		hang--
		lie--
	}
}
