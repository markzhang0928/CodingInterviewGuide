package main

import "fmt"

func returnButDefer2() (r int) {
	r = 5
	defer func() {
		r = r + 5
	}()
	return r
}

func returnButDefer() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
func main() {
	fmt.Println(returnButDefer())
	fmt.Println(returnButDefer2())
}

//import "fmt"
//
//type number int
//
//func (n number) print() {
//	fmt.Println(n)
//}
//
//func (n *number) pprint() {
//	fmt.Println(*n)
//}
//
//func main() {
//	var n number
//
//	defer n.print()
//	defer n.pprint()
//	defer func() { n.print() }()
//	defer func() { n.pprint() }()
//
//	n = 3
//}

//func main() {
//	var whatever [3]struct{}
//	fmt.Printf("type is %T \n", whatever)
//	for i := range whatever {
//		defer func() {
//			fmt.Printf("task[%d]: %d\n", i, i)
//		}()
//	}
//}

//func main() {
//	sl := []int{11, 12, 13, 14, 15}
//	var wg sync.WaitGroup
//	for i, v := range sl {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			fmt.Printf("task[%d]: %d\n", i, v)
//		}()
//	}
//	wg.Wait()
//}

//func main() {
//	p := workerpool.New(5, workerpool.WithPreAllocWorkers(false), workerpool.WithBlock(false))
//
//	time.Sleep(2 * time.Second)
//	for i := 0; i < 10; i++ {
//		err := p.Schedule(func() {
//			time.Sleep(time.Second * 3)
//		})
//		if err != nil {
//			fmt.Printf("task[%d]: error: %s\n", i, err.Error())
//		}
//	}
//
//	p.Free()
//}
