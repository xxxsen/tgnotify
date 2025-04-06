TGNotify
===

[TOC]

一个用于给自己发送消息提醒的机器人。

## 配置

```jsonc
    "listen": ":9902", //监听地址
    "chatid": 123, //要给指定用户发送消息的用户chatid
    "token": "2345", //机器人token
    "users": {    //鉴权信息
        "abc": "123456" //key:value
    },
    "log_config": {
        "level": "debug",
        "console": true
    }
```

## 运行

```yaml
services:
  tgnotify:
    image: xxxsen/tgnotify:latest
    container_name: "tgnotify"
    restart: always
    volumes:
      - ./config:/config
    expose:
      - 9902
    command: --config=/config/config.json  
```