---
layout:     post
title:      "All For Redis"
subtitle:   " \"Redis知识点整理\""
date:       2022-06-20 22:00:00
author:     "QuXY"
header-img: "img/post-bg-halting.jpg"
catalog: true
tags:
    - Redis
    - Middleware
---

## Part of Data Structure

### 1. The process of rehashing

`hash` is one of five data structure in redis, and the other are `string`, `list`, `set`, `zset`

the C struct of `hash` is:
```c
// dict.h
typedef struct dictEntry {
    void *key;
    union {
        void *val;
        uint64_t u64;
        int64_t s64;
        double d;
    } v;
    struct dictEntry *next;
} dictEntry;


typedef struct dictht {
    dictEntry **table;      // 指向哈希表数组第一位的指针，指向dictEntry*的数组:[&dictEntry{},&dictEntry{},...]
    unsigned long size;     // 索引位置数
    unsigned long used;     // 键值对的数量
    // ...
} dictht;


typedef struct dict {
    dictht ht[2];   // 一个dictht存数据，另一个只用在rehash的时候
    long rehashidx; // 没有在rehash时值为-1，rehash正式开始时值为0
    // ...
} dict;
```

The process of rehashing:
1. `ht[1]` requires new memory with new size, and `rehashidx` would be set 0
2. Everytime one key was updated (or created/delete/query), the hashed key(using old hash algorithm) should be calculated firstly. If the hashed key was bigger than `rehashidx`, the operation would affect `ht[1]` instead of `ht[0]`. Meanwhile if the hashed key was smaller than `rehashidx`, the updating operation would be in the `ht[0]`
3. Everytime one key was updated (or created), copy the `ht[0]` key-value in position `rehashidx`, and add the `rehashidx`
4. When the `rehashidx` equals to old table size, make `ht[0]` point to `ht[1]`, and delete the old memory. 
5. The `rehashidx` is set -1


