创业本地Go项目
资源						信息管理系统平台					中国    			美国
    司机、汽车			阳光 							
	房产					贝壳
	商家					阿里
	图书					亚马逊
	资讯					字节
	社交用户				微博
	二手车
			
机械化			降低成本
量产				降低成本			造火箭的成本  特斯拉的成本  因为科技变的很低很低
智能化			降低成本
使用科技降低成本
MAC下使用docker搭建环境
基于Docker镜像部署go项目（实例详解)
https://www.php.cn/docker/488707.html
本地用Docker搭建go开发环境
https://blog.csdn.net/weixin_42687353/article/details/124656022


docker命令总结
命令	用途
docker pull	远程拉取image
docker build	创建image
docker images	列出本地所有的image
docker run	运行容器container
docker ps	列出当前运行的容器container
docker rm	删除已经结束的容器container
docker cp	在本地和容器之间拷贝文件
docker commit	保存改动生成新的image
docker rmi	删除image


自己总结
docker怎么删除镜像
	docker rm 
	1.首先用docker ps -a 列出所有容器 可以看到这里有两个mysql正在运行 我们想要删除镜像 就必须先删除容器 不然就会报各种错误
	2.然后我们输入docker rm 7887634e4663删除容器
	3.找到镜像id docker images
	4.这时候输入 docker rmi 镜像id 就可以删除我们想要删除的镜像了


我自己的

1.启动容器使用docker软件
2.停止容器  docker  stop 6610d427f4a4
3.docker进入容器
	docker ps -a
	docker exec -it 6610d427f4a4 /bin/bash
4.给自己的容器建立本地映射添加名字为	golang-1.18.1-pengzhan
	docker run -itd -p 8080:8080 -v /Users/pengzhanliang/go/src:/go --name golang-1.18.1-pengzhan golang-local


连接远程的token  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
	Personal access tokens:	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
	git remote set-url origin https://<your_token>@github.com/<USERNAME>/<REPO>.git
	git remote set-url origin https://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@github.com/SilhouetteLiang/singo.git
http://www.sosoapi.com/pass/apidoc/share/forward.htm?shareKey=2c06993de448ecebdb7554870146e1ae  密码：123456

题库录入，话术录入，性格测试结果展示，用户列表，支付记录
	1、首页
	2、题库录入
	3、话术录入
	4、结果1
	5、结果2
	6、结果3
	7、获取用户信息

1、运行

go run asciijson.go
go build


2、木木不会
    GROM操作mysql数据库
	https://blog.csdn.net/qq_43514659/article/details/121466144?spm=1001.2014.3001.5502
    Go_web gin框架开发（通用脚手架）
    https://blog.csdn.net/qq_43514659/article/details/121587913?spm=1001.2014.3001.5502
    zap日志的基本使用（go必会知识*）
    https://blog.csdn.net/qq_43514659/article/details/121567519?spm=1001.2014.3001.5502
3、使用Singo 搭建本地项目
    liangpengzhan@outlook.com       Wslpz888888

在/Users/pengzhanliang/go/src/gin-liang/singo执行  go run main.go
Redis  密码   sl2016SL!


mysql账户密码		test  Lpz888888.
修改mysql权限，
1、需要想修改  vim /usr/local/etc/my.cnf  跳过权限表  
2、然后使用 root登录  添加用户或者删除用户，执行flush privileges才能成功
3、/usr/local/Cellar/mysql/8.0.17/bin执行mysql.server stop	加载修改后的usr/local/etc/my.cnf 文件
4、vim /usr/local/etc/my.cnf  调整为需要权限表  

V2版本
	
使用算法实现一个机器人，关于
	心里咨询的
		实现一个算法给不同人推荐不同的心里咨询方案
	心理资讯：
		把客户浏览的内容、某一类内容的浏览时间 ======== 作为推荐内容的依据，推荐相关联的咨询。

推荐算法平台
实物平台应用

平台化＋算法推荐

Singo: Simple Single Golang Web Service

go-crud正式改名为Singo!

使用Singo开发Web服务: 用最简单的架构，实现够用的框架，服务海量用户

https://github.com/Gourouting/singo

## Singo文档

https://gourouting.github.io/

## 视频实况教程

[让我们写个G站吧！Golang全栈编程实况](https://space.bilibili.com/10/channel/detail?cid=78794)

## 使用Singo开发的项目实例

仿B站的G站：https://github.com/Gourouting/giligili

Singo框架为移动端提供Token登录的案例: https://github.com/bydmm/singo-token-exmaple
## 目的

本项目采用了一系列Golang中比较流行的组件，可以以本项目为基础快速搭建Restful Web API

## 特色

本项目已经整合了许多开发API所必要的组件：

1. [Gin](https://github.com/gin-gonic/gin): 轻量级Web框架，自称路由速度是golang最快的 
2. [GORM](https://gorm.io/index.html): ORM工具。本项目需要配合Mysql使用 
3. [Gin-Session](https://github.com/gin-contrib/sessions): Gin框架提供的Session操作工具
4. [Go-Redis](https://github.com/go-redis/redis): Golang Redis客户端
5. [godotenv](https://github.com/joho/godotenv): 开发环境下的环境变量工具，方便使用环境变量
6. [Gin-Cors](https://github.com/gin-contrib/cors): Gin框架提供的跨域中间件
7. 自行实现了国际化i18n的一些基本功能
8. 本项目是使用基于cookie实现的session来保存登录状态的，如果需要可以自行修改为token验证

本项目已经预先实现了一些常用的代码方便参考和复用:

1. 创建了用户模型
2. 实现了```/api/v1/user/register```用户注册接口
3. 实现了```/api/v1/user/login```用户登录接口
4. 实现了```/api/v1/user/me```用户资料接口(需要登录后获取session)
5. 实现了```/api/v1/user/logout```用户登出接口(需要登录后获取session)

本项目已经预先创建了一系列文件夹划分出下列模块:

1. api文件夹就是MVC框架的controller，负责协调各部件完成任务
2. model文件夹负责存储数据库模型和数据库操作相关的代码
3. service负责处理比较复杂的业务，把业务代码模型化可以有效提高业务代码的质量（比如用户注册，充值，下单等）
4. serializer储存通用的json模型，把model得到的数据库模型转换成api需要的json对象
5. cache负责redis缓存相关的代码
6. auth权限控制文件夹
7. util一些通用的小工具
8. conf放一些静态存放的配置文件，其中locales内放置翻译相关的配置文件

## Godotenv


项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建.env文件设置环境变量便于使用(建议开发环境使用)

```shell
MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接地址
REDIS_ADDR="127.0.0.1:6379" # Redis端口和地址
REDIS_PW="" # Redis连接密码
REDIS_DB="" # Redis库从0到10
SESSION_SECRET="setOnProducation" # Seesion密钥，必须设置而且不要泄露
GIN_MODE="debug"
```

## Go Mod

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```shell
go mod init go-crud
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go run main.go // 自动安装
```

## 运行

```shell
go run main.go
```

项目运行后启动在3000端口（可以修改，参考gin文档)

