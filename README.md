## 所需环境
- fabric
- beego
- mysql
- redis
- fastdfs

环境安装方式很常规自行google就行
项目需要放在$GOPATH/src目录下
## 执行
项目配置文件在conf/app.conf
```
appname = airhome
httpport = 8080
httpaddr = "127.0.0.1"
runmode = dev
copyrequestbody = true
sessionon = true
redisaddr = "127.0.0.1"
redisport = 6379
redisdbnum = 0
mysqladdr = "127.0.0.1"
mysqlport = 3306
mysqldbname = "airhome"
mysqlusername = "root"
mysqlpassword = "123456"
fdfs_http_addr = "127.0.0.1:9090"
```
conf/client.conf，这是fastdfs的client配置服务。配置好了才能上传图片。
首先在mysql中创建名字为mysqldbname = "airhome" 的数据库。
执行beego命令：bee run，会自动创建表并可以启动项目。在mysql中执行conf/area.sql中的语句添加数据。
不出意外的话说明beego启动成功，此时还没有启动fabric，结束执行，并在目录中执行fabricSDK启动命令：
进入项目的根目录并使用`make`命令：

-   任务`all`  ：  `make`或`make all`
-   任务`clean`  ：清理一切并释放网络（  `make clean`  ）
-   任务`build`  ：只需构建应用程序（  `make build`  ）
-   任务`env-up`  ：只需建立网络（  `make env-up`  ）

**退出时一定要执行make clean关闭正在运行的docker容器**




