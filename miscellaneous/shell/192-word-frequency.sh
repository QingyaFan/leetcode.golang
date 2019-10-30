#!/bin/bash

# 写一个 bash 脚本以统计一个文本文件 words.txt 中每个单词出现的频率
# 假设： words.txt只包括小写字母和 ' '
#       每个单词只由小写字母组成
#       单词间由一个或多个空格字符分隔

cat words.txt | xargs -n 1 | sort | uniq -c | sort -nr | awk '{print $2" "$1}'