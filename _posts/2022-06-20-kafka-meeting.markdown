---
layout:     post
title:      "All For Kafka"
subtitle:   " \"Kafka知识点整理\""
date:       2022-06-20 22:00:00
author:     "QuXY"
header-img: "img/post-bg-unix-linux.jpg"
catalog: true
tags:
    - Kafka
    - Middleware
---

## Part of Core of MQ

### 1. How Kafka guarantee 'consume at least one'

Kafka guarantees the message consumed at least noce and only once.

To producers:
1. 