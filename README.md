# webhook
Web hook to exec command.

## What

webhook is a simple tool wriiten in go. It can executing remote commands over the web hook.

Very easy to use.


## Getting started

### Download

Choose the version of your computer system and download it, then copy to the server client which you want to control.

### Add configuration

```
# 'test_router' is url router to execute the command which followed it.
# 'echo 1' is command to be executed
./webhook -add "test_router:echo 1"

```

or

### Create new config file, content like this.

```
# Now you can request `http://your_server_host_name:8910/test1` to execute the command `echo 123`.
test1:echo 123
runtask:sh builder.sh

```

then

```
# Start the server
# Default port is 8910
./webhook
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





