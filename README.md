# bullet-text-chatroom
[![Build Status](https://www.travis-ci.org/haowanxing/bullet-text-chatroom.svg?branch=master)](https://www.travis-ci.org/haowanxing/bullet-text-chatroom)
[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://github.com/haowanxing/bullet-text-chatroom/blob/master/LICENSE)

> 利用Go-Websocket实现的弹幕聊天室  
项目地址：https://github.com/haowanxing/bullet-text-chatroom  
此项目为个人学习练手项目，欢迎各路大神轻踩

### 项目介绍

* 使用Golang实现的简易Web多人聊天室
* 支持房间ID标识，自由创建和切换房间（Demo中使用URL的路径切换房间）
* 支持用户改名，名称随意起
* 支持用户的加入和离开提醒
* 消息展现方式自由，默认采用弹幕方式，灵活有趣
* （更多有趣内容...)

### 运行方法

> 参考[演示代码](https://github.com/haowanxing/bullet-text-chatroom/blob/master/example/main.go)

### 更新日志

#### 2019-08-08
> 1.替换官方websocket包，采用gorilla/websocket   
2.新增Dockerfile
3.聊天记录中采用红色文字区分系统消息和普通用户消息，并增加回车按钮直接“发射”   

#### 2019-08-06
> 1.新增修改昵称功能
2.采用canvas画布实现聊天弹幕，并创建client.js封装ws客户端

### 遗留问题

1. ~~IOS下测试失败，可能是对websocket的支持问题？~~ 升级IOS12.4后正常（旧版本为IOS9）
2. ~~golang.org/x/net/websocket获取不到？~~ 采用更优秀的gorilla/websocket包

### 展示图

![demo.png](https://github.com/haowanxing/bullet-text-chatroom/blob/master/demo.png)
