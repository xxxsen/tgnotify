TGNotify
===

[TOC]

一个用于给自己发送消息提醒的机器人。

## 使用流程

1. 建立自己的机器人, 使用BotFather进行快速创建, 并拿到自己的机器人token, 具体的过程可以自己找资料。
2. 部署自己的机器人后端(参考`后端部署`一节)。
3. 使用`/reg` 在机器人上进行注册。(相关命令可以看下面的`TG命令`一节)
4. 使用curl给自己发送消息(参考`CURL消息`一节)。

可以通过目录下的 docker-compose.yml 进行部署, 需要修改自己的bot token。

### TG命令

以下命令只能在机器人聊天框使用

- /reg 注册用户
- /reset 更换注册信息
- /info 获取当前注册的信息
- /delete 删除自己的注册信息

### CURL消息

机器人建立完成, 并且后端也部署完成后, 可以使用下面的curl命令来给自己发送消息(curl发送后会在机器人里面展示自己发送的消息)。

基本格式:

```shell 
curl -d "这里是要发送的消息" -H 'user: 这里填写在机器人里面注册的用户名' -H 'code: 这里填写注册的code(不是token)' "http://你自己的域名:端口"  -v
```

上面发送的内容还支持markdown和html格式, 可以使用`-H 'mode: markdown'` 或者`-H 'mode: html'` 来发送markdown或者html格式的内容。

**NOTE**: 发送完后可以通过http的响应码来判断发送十分成功

一个例子:

```shell
curl -d "hello world" -H 'user: ${chatid}' -H 'code: ${code}' "https://test.com" -v
```

## 后端部署

推荐使用docker-compose进行部署

### docker-compose部署

传入下面几个ENV即可

- LISTEN: 监听地质
- SAVE_FILE: 用户数据保存文件
- LOG_LEVEL: 日志级别
- TOKEN: 机器人token


### 手工部署

略


###### tags: `tgnotify` `readme` 