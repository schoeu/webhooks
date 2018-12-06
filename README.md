# Webhooks
Web hook to exec command.

[![Build Status](https://travis-ci.org/schoeu/webhooks.svg?branch=master)](https://travis-ci.org/schoeu/webhooks)
[![Go Report Card](https://goreportcard.com/badge/github.com/schoeu/webhooks)](https://goreportcard.com/report/github.com/schoeu/webhooks)
[![GoDoc](https://godoc.org/github.com/schoeu/webhooks?status.svg)](https://godoc.org/github.com/schoeu/webhooks)


[中文版Readme](./README-zh_CN.md)

## What is webhooks

Webhooks is a simple tool wriiten in golang. It can executing remote commands over the web hook, also very easy to use.


## Getting started

### Download

Choose the version of your computer system and download it, then copy it to the server which you want to control.

- [linux-32bit](http://weather.schoeu.com/webhook_linux_32bit)
- [linux-64bit](http://weather.schoeu.com/webhook_linux_64bit)
- [MAC](http://weather.schoeu.com/webhook_mac)
- [windows-32bit](http://weather.schoeu.com/webhook_32bit.exe)
- [windows-64bit](http://weather.schoeu.com/webhook_64bit.exe)

## Help

Run `./webhooks --help` command you can get full information for use the webhooks.

```
./webhooks --help
```

## Command mode

Just request url `http://your_server_host_name:your_port/run/your_command`

### Examples

- Get runtime path you can request:

    - `http://localhost:8910/run/pwd`

- Get some file list you can request:

    - `http://localhost:8910/run/ls`

- Also you can get text type file content please request:

    ```
    http://localhost:8910/run/cat test.txt

    or

    http://localhost:8910/run/cat test.pdf

    or

    http://localhost:8910/run/cat test.jpg
    ```

and so on ...


## Tasks mode

### Add configuration

```
# 'test_router' is router to execute the command which followed it.
# 'echo 1' is the command to be executed
./webhooks -add "test_router:echo 1"
```

Or

### Create new configuration

```
# Now you can request `http://your_server_host_name:8910/tasks/test1` to execute the command `echo 123`.
test1:echo 123
runtask:sh builder.sh
```

Then

```
# Start the server
# Default port is 8910
./webhooks

# If you want to change port follow this:
./webhooks --port 8911
```


## Token

Webhooks can run script on your server, it must be run after some inspection. So we can go through with token...

```
# Run server with token.
./webhooks --token your_token
```
Then you can request url `http://youdomain:port/run/some_script?token=your_token` to run your script.


## Return json
You only need to take the `json` parameter in the request url.
Such as:

```
http://youdomain:port/run/some_script?token=your_token&json=1
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
