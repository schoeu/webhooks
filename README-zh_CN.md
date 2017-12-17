# Webhooks

> 使用web访问来执行机器命令，触发任务等。

[![Build Status](https://travis-ci.org/schoeu/webhooks.svg?branch=master)](https://travis-ci.org/schoeu/webhooks)
[![Go Report Card](https://goreportcard.com/badge/github.com/schoeu/webhooks)](https://goreportcard.com/report/github.com/schoeu/webhooks)
[![GoDoc](https://godoc.org/github.com/schoeu/webhooks?status.svg)](https://godoc.org/github.com/schoeu/webhooks)


## 简介

Webhooks是一个用`go`语言实现的可以在web端执行机器命令的工具，工具非常的简单易用。通常用来做任务触发，机器自动化维护等需求。


## 开始

### 下载

选择一个对应的系统版本，下载到你想执行命令的机器。

- [linux32位系统](http://ozo2fe2cm.bkt.clouddn.com/webhook_linux_32bit)
- [linux64位系统](http://ozo2fe2cm.bkt.clouddn.com/webhook_linux_64bit)
- [MAC系统](http://ozo2fe2cm.bkt.clouddn.com/webhook_mac)
- [windows32系统](http://ozo2fe2cm.bkt.clouddn.com/webhook_32bit.exe)
- [windows64系统](http://ozo2fe2cm.bkt.clouddn.com/webhook_64bit.exe)


下载后

在*nix机器中对应目录中执行`./webhook_xxx`即可，执行之前可能需要用一下命令授权：
```
chmod -x ./webhook_xxx
```
在windows机器，双击执行，或者在CMD中找到`webhook.exe`执行文件，然后执行
```
webhook.exe
```


## 帮助

执行以下就可以获得该工具的详细帮助信息。

```
./webhooks --help
```

例子： 假如在`123.123.123.123`机器中下载，启动了webhooks服务，那么就可以访问`http://123.123.123.123:8910/`来访问服务。

访问的格式为： ``http://ip:port/[run|tasks]/command[?token=some_token][&json=1]`


## 命令模式

直接访问 `http://机器host:端口/run/你要执行的命令`

### 示例
以linux机器为例

- 获取机器当前执行路径:

    - `http://localhost:8910/run/pwd`

- 获取当前目录中的文件列表:

    - `http://localhost:8910/run/ls`

- 当然你也可以获取目标机器上的一些文本文件:

    ```
    http://localhost:8910/run/cat test.txt

    或者

    http://localhost:8910/run/cat test.pdf

    或者

    http://localhost:8910/run/cat test.jpg
    ```

    这些文件，图片，pdf都会直接在你的浏览器中展现出来。

其他的各种用法请自行发现 ...


## 任务模式

### 添加配置

```
./webhooks -add "test:echo 1"

# 添加该命令后，当我们访问 `http://localhost:8910/tasks/test` 时就会执行后面跟着的`echo 1`命令。


```

或者

### 创建新配置文件

```
vim ./webhooks.conf

# 输入以下内容
test1:echo 123
runtask:sh builder.sh

# 然后保存，重启服务后，我们就可以使用
# `http://localhost:8910/tasks/test1`, `http://localhost:8910/tasks/runtask` 来执行对应的命令了


```

之后

```
# 启动服务
# 默认端口为8910
./webhooks

# 如果想自定义一个端口，比如自定义为8911，可以这样
./webhooks --port 8911
```


## 指令

webhook可以直接执行远程目标机器上的命令，对于目标机器来讲不安全，所以可以添加token来增强安全性。

```
# token模式启动服务, 后面的`your_token`就是自定义token
./webhooks --token your_token
```
现在只能访问 `http://localhost:8910/run/指定命令?token=your_token` 才能运行命令。


## 返回JSON
webhooks默认直接返回内容，如果想要返回json格式的内容，可以这样：

```
http://localhost:8910/run/指定命令?token=your_token&json=1
```


## MIT License

Copyright (c) 2017 Schoeu

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
