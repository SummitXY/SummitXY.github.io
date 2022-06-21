---
layout:     post
title:      "LeetCode: Double Pointer For Array"
subtitle:   " \"双指针(快慢指针)场景整理\""
date:       2022-06-19 22:00:00
author:     "QuXY"
header-img: "img/post-bg-universe.jpg"
catalog: true
tags:
    - LeetCode
---

## Problem 1

Give you one string and you must delete and count all space and return the new one.
> 删除并算出字符串里所有空格, 并返回新字符串

There are two stricted condition:
1. Not allowed to make new memory space except simple variables
2. The time complexity should be O(n)

### Solution one

The first solution uses two pointer:
```go
func main() {

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
```

two pointer fast 

### Solution two

The second solution uses Golang's feature: slice structure

```go
func main() {

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

```
> the `res := a[:0]` would not make new space, conversely `res` would reuse the array in slice `a`

---

## Problem 2

[Leetcode 142](https://leetcode.cn/problems/linked-list-cycle-ii/):

Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.

```go
func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    fast, slow := head, head
    for fast.Next != nil && fast.Next.Next != nil{
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {
            // start
            p := head
            for p != slow {
                p = p.Next
                slow = slow.Next
            }
            return p
        }
    }

    return nil
}
```

Create two pointer: faster(walks two steps once) and slower(walks one step once)

When faster meets slower :

Assume that faster has walked `f` steps and slower has walked `s` steps and we know `f = 2*s`.

Assume the distance from head to cycble point (exclude cycle point) is `a`, and the length of cycle (include cycle point) is `b`, so we can get `f = s + n*b` (faster walk more `n*b` steps than slower)

From the above formula , we can get `s = n*b`, and we know walking from head with `a+n*b` steps can get to the cycle point, so let slower walks `a` steps continuely can get to the cycyle point

end