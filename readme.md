TGNotify
===

[TOC]

一个用于给自己发送消息提醒的机器人。

## 使用流程

1. 建立自己的机器人, 使用BotFather进行快速创建, 并拿到自己的机器人token, 具体的过程可以自己找资料。
2. 部署自己的机器人后端(参考`后端部署`一节)。
3. 使用`/reg ${user} ${code}` 在机器人上进行注册。(相关命令可以看下面的`TG命令`一节)
4. 使用curl给自己发送消息(参考`CURL消息`一节)。

可以通过目录下的 docker-compose.yml 进行部署, 需要修改自己的bot token, 还有db存放目录。

### TG命令

以下命令只能在机器人聊天框使用

- /reg ${user} ${code} 注册用户
- /chg ${user} ${code} 更换注册信息
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
curl -d "hello world" -H 'user: test' -H 'code: test' "https://test.com" -v
```

## 相关配置

采用json进行配置, 大致有下面这些信息

- server_config 基本的服务端配置
    - listen 服务监听地址(http协议的)
- db_config db配置
    - host db域名
    - port db端口
    - user db用户名
    - pwd db密码
    - dbname db名
- bot_config
    - token bot的token, 在botfather建立bot后会产生这个东西

基本配置: 

```json
{
    "server_config":{
        "listen":"${server listen}"
    },
    "db_config":{
        "host":"${your db host}",
        "port":${your db port},
        "user":"${your db user name}",
        "pwd":"${your db password}",
        "dbname":"${your dbname}"
    },
    "bot_config":{
        "token":"${your bot token}"
    }
}
```

下面是一个简单的例子

```json
{
    "server_config":{
        "listen":"0.0.0.0:8333"
    },
    "db_config":{
        "host":"192.168.1.10",
        "port":3306,
        "user":"test",
        "pwd":"test",
        "dbname":"tgmessager"
    },
    "bot_config":{
        "token":"abcdefghijklmn"
    }
}
```

## 后端部署

推荐使用docker-compose进行部署

### docker-compose部署

1. 复制工程里面的dockerrun目录到你自己需要放置的位置(dockerrun目录可以改名成你自己需要的名字).
2. 修改shell_args.sh 里面的`BOT_TOKEN`变量为你的bot的token.
3. 运行init_config.sh, 这个shell会在当前目录下生成相关的配置.(运行后就不要再移动dockerrun了)
4. 运行`docker-compose up -d` 来拉起容器.

这里几个脚本的作用

- init_config.sh 生成docker-compose.yml 运行需要的相关目录和配置
- remove_container.sh 停止并删除docker-compose生成的容器
- remove_folder.sh 删除init_config.sh生成的相关目录(包括DB的数据)


### 手工部署

略


###### tags: `tgnotify` `readme` 