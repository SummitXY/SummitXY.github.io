package main

// 给定字符串删除所有空格, 并返回空格个数
// 要求时间复杂度O(n), 空间复杂度O(1)

import "fmt"

func main() {

	// 字符串必须用切片表示, 不然无法实现要求
	a := []byte{' ', ' ', ' ', ' ', ' '}

	left, right := 0, 0
	for right < len(a) {
		if a[left] != ' ' && a[right] != ' ' {
			left++
			right++
			continue
		}

		if a[left] == ' ' && a[right] == ' ' {
			right++
			continue
		}

		if a[left] == ' ' && a[right] != ' ' {
			a[left] = a[right]
			a[right] = ' '
			left++
			right++
			continue
		}
	}

	fmt.Println("the space coun is ", len(a)-left)
	fmt.Println("the answer is ", a[:left])
}

func main3() {

	fmt.Println("start")

	a := []byte{'a', ' ', ' ', 'b', 'c'}
	size := len(a)
	fast, slow := 0, 0

	count := 0
	for fast < size {
		if a[slow] == ' ' {
			for fast < size {
				if a[fast] != ' ' {
					break
				}
				count++
				fast++
			}
			if fast == size {
				break
			}
			a[slow] = a[fast]
		}
		slow++
		fast++
	}
	fmt.Println("the space and final string are ", count, a[:slow])
	// the space and final string are  2 [97 98 99]
}

func main2() {

	fmt.Println("start")

	a := []byte{'a', ' ', ' ', 'b', 'c'}
	size := len(a)
	fast := 0

	count := 0
	res := a[:0]
	fmt.Println("the length and capacity of res are ", len(res), cap(res))
	// output : the length and capacity of res are  0 5
	for fast < size {
		if a[fast] == ' ' {
			count++
			fast++
			continue
		}
		// != ' '
		res = append(res, a[fast:fast+1]...)
		fast++
	}

	fmt.Println("the space and final string are ", count, res)
	// output : the space and final string are  2 [97 98 99]
}
