block-web-server 可以按照预期 阻塞 的 http web server
----

[TOC]

## 功能

### 1. block

`/block{x}` 这个接口的任何方法(get/put) `阻塞` `x 秒` （`x`必须为数字）

比如 
接口`/block1`， 会阻塞 `1秒`  
接口`/block2`， 会阻塞 `2秒`  
接口`/block3`， 会阻塞 `3秒`  

等等

### 2. welcome
其它接口
1. `block`后面不是数字,如 `/blockabc`, 立刻返回`200`
2. 其它接口，如 `/abc`, 立刻返回`200`

## 使用
提供了两种思路
1. win10, 请到 release 里面下载
2. docker, 镜像地址为 `hwhaocool/block-web-server:latest`


