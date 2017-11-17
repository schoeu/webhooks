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
# 'test_router' is url router to execute the command follow it.
# 'echo 1' is command to be executed
./webhook -add "test_router:echo 1"

```

or

```
### Create new config file, content like this.

test1:echo 123
runtask:sh builder.sh

```

then

```
# Start the server
# Default port is 8910
./webhook
```

### Enjoy it.




