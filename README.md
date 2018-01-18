使用control进行操作
```
╰─$ ./control
Usage: ./control {start|stop|restart|build|status}
```
### 1. build
```
./control build
```
### 2. start
```
./control start
```

### 3. test
```
curl http://localhost:9090/hello
```

### 4. stop
发出SIGTERM信号让服务停止停止
```
./control stop
```

### 日志输出
/var/log/hello.log
```
time="2018-01-18T12:37:13+08:00" level=info msg="init logger success"
time="2018-01-18T12:37:13+08:00" level=info msg="start server ... ..."
time="2018-01-18T12:37:13+08:00" level=info msg="register signal handler"
time="2018-01-18T12:37:13+08:00" level=info msg="start server" addr="0.0.0.0:9090"


time="2018-01-18T12:37:30+08:00" level=info msg="get signal SIGTERM, prepare exit!!!"
time="2018-01-18T12:37:30+08:00" level=info msg="get signal SIGTERM, success exit!!!"
time="2018-01-18T12:37:30+08:00" level=info msg="server stop listen"
```
