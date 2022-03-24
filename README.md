# 武汉理工大学研究生健康报送
**仅用于节省时间，身体出现不适立刻停止使用脚本并向学校报告！**

## 使用方法
1. [修改配置文件](./res/userInfo.json)
      主要修改 ```sn(学号)``` 和 ```idCard(登录密码，默认身份证后六位)```
2. 登录微信小程序 -> 个人中心 -> 关联设置 -> 取消关联
   (再次关联，只需要打开小程序登录即可)
3. ```make build``` 产生二进制执行文件
4. ```make run``` 运行打卡程序

### 注意
如果想要单独的二进制文件作为脚本运行，但提示  
```open file err =  open ./res/userInfo.json: no such file or directory  ```  

```open ./res/session_backup.txt: no such file or directory```
  
  表示生成的可执行文件需要指定对应的 userInfo.json。通过--path 指定路径即可
  ```./auto-check --path /res/userInfo.json```



## 联系方式
https://qiqi-note.top/about-me


## 有时间随缘更新~

- [ ] Docker构建
- [ ] 更清晰的log
- [ ] 多用户支持