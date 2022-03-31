# 武汉理工大学研究生健康报送
**仅用于节省时间，身体出现不适立刻停止使用脚本并向学校报告！**

## 使用方法
### 一键报送
1. [修改配置文件](./res/userInfo.json)
   主要修改 ```sn(学号)``` 和 ```idCard(登录密码，默认身份证后六位)```
2. 登录微信小程序 -> 个人中心 -> 关联设置 -> 取消关联
   (再次关联，只需要打开小程序登录即可)
3. 执行bin下面的auto-check

### 定时报送
1. [修改配置文件](./res/userInfo.json)
   主要修改 ```sn(学号)``` 和 ```idCard(登录密码，默认身份证后六位)```
2. 登录微信小程序 -> 个人中心 -> 关联设置 -> 取消关联
   (再次关联，只需要打开小程序登录即可)
3. 把bin文件下的[auto-check](./bin/auto-check),与[res](./res)复制到新的文件夹下
4. 执行```crontab -e``` 加入下面一行
```shell
0 */13 * * * 3中新文件路径/auto-check --path 你用来存放userInfo.json的路径/userInfo.json > 输出log文件的路径/log.txt
```
5. 执行```service cron restart```
### 自定义功能
1. [修改配置文件](./res/userInfo.json)
      主要修改 ```sn(学号)``` 和 ```idCard(登录密码，默认身份证后六位)```
2. 登录微信小程序 -> 个人中心 -> 关联设置 -> 取消关联
   (再次关联，只需要打开小程序登录即可)
3. 修改代码
4. ```make build``` 产生二进制执行文件
5. ```make run``` 运行打卡程序

### 注意
1.如果想要单独的二进制文件作为脚本运行，但提示  
```open file err =  open ./res/userInfo.json: no such file or directory  ```  

```open ./res/session_backup.txt: no such file or directory```
  
  表示生成的可执行文件需要指定对应的 userInfo.json。通过--path 指定路径即可
  ```./auto-check --path /res/userInfo.json```
2. res下面的session_backup.txt和userInfo.json需要放置在同一文件夹下 



## 联系方式
https://qiqi-note.top/about-me


## 有时间随缘更新~

- [ ] Docker构建
- [ ] 更清晰的log
- [ ] 多用户支持
