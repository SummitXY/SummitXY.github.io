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

### 1. At least once & At most once

There're two topic:
1. How does Kafka guarantee all messages consumed
2. How does Kafka avoid or handle duplicate messages


In fact, there are lots of messaging system, and Four types of strategy exist:
1. No Guarantee:
2. At most once:
3. At least once:
4. Exactly once:

![](https://raw.githubusercontent.com/SummitXY/img/master/kafka-at-least-once.png)


So why isn't it always exactly once? That's because 


## Ref

[Medium: At least once](https://medium.com/@andy.bryant/processing-guarantees-in-kafka-12dd2e30be0e#id_token=eyJhbGciOiJSUzI1NiIsImtpZCI6IjJiMDllNzQ0ZDU4Yzk5NTVkNGYyNDBiNmE5MmY3YjM3ZmVhZDJmZjgiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJuYmYiOjE2NTU5MDk4NzcsImF1ZCI6IjIxNjI5NjAzNTgzNC1rMWs2cWUwNjBzMnRwMmEyamFtNGxqZGNtczAwc3R0Zy5hcHBzLmdvb2dsZXVzZXJjb250ZW50LmNvbSIsInN1YiI6IjEwNjY4MDExMzA1NjIyNjE4NTI1MiIsImVtYWlsIjoic3VtbWl0eHkxMjNAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImF6cCI6IjIxNjI5NjAzNTgzNC1rMWs2cWUwNjBzMnRwMmEyamFtNGxqZGNtczAwc3R0Zy5hcHBzLmdvb2dsZXVzZXJjb250ZW50LmNvbSIsIm5hbWUiOiJDaGF1bmNleSBRIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FBVFhBSnhhTmFMd2FNUU1jTmN0MmxYaHhGVmRlRi1GQUhjMzJhOUZBdVZjPXM5Ni1jIiwiZ2l2ZW5fbmFtZSI6IkNoYXVuY2V5IiwiZmFtaWx5X25hbWUiOiJRIiwiaWF0IjoxNjU1OTEwMTc3LCJleHAiOjE2NTU5MTM3NzcsImp0aSI6Ijk1ZWZkNzlhMmNlYTM2Yzg3MjM2YWNjM2E3Y2MxZDVjN2U3ZGVkMTkifQ.MtDRUAZoHRjmddik1qIoVcFUy710ENn8d_RHEfaY3UU5uF6Fa77zgGZTBJuztpNE2Dd1trbDPKUngayU_QKaZuOR8Bfid6xdu6JwnCCIGBiC2AVx-tQQ_a0HIil0Zwn--rnM24DH5FHTCK3GVMeX4vgpe6FftV35Q33_fi_Dlc7vap5oihyTPP5tWUhbXfD9wvl65jdYoH3YPhqDkuuM5IqscX7goEy8qunPBCM_zdhiquW4Uvx5YhVBzQMTITQ797Nx5NgbpeCwPX6OUGntAR9A_Xy6Gh75ozlG7FJ3SV4pVEmKdP13IYDpM9YVKhjMBQ51hfKgJjKKgT3hX6oxCw)