---
layout:     post
title:      "LeetCode: Sliding Window Algorithm"
subtitle:   " \"滑动窗口算法(可变窗口/固定窗口)\""
date:       2022-06-23 22:00:00
author:     "QuXY"
header-img: "https://raw.githubusercontent.com/SummitXY/img/master/blogniloy-tesla-fVE7VhswCJI-unsplash.jpg"
catalog: true
tags:
    - LeetCode
---

## Problem 1

[LeetCode209](https://leetcode.cn/problems/minimum-size-subarray-sum/)

Given an array of positive integers nums and a positive integer target, return the minimal length of a contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr] of which the sum is greater than or equal to target. If there is no such subarray, return 0 instead.

```go
func minSubArrayLen(target int, nums []int) int {
    if len(nums) < 2 {
        return 0
    }

    minLen := 1000000
    left, right := 0, 0
    nowSum := 0
    for right < len(nums) {
        nowSum += nums[right]
        right ++
        for nowSum >= target {
            if right - left < minLen {
                minLen = right - left
            }
            nowSum -= nums[left]
            left ++
        }
    }

    if minLen < 1000000 {
        return minLen
    }

    return 0
}
```


## REF

