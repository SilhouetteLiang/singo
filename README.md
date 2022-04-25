# Singo
本地Go项目


1、运行

go run asciijson.go
go build


2、木木不会

- [ ] GROM操作mysql数据库

	https://blog.csdn.net/qq_43514659/article/details/121466144?spm=1001.2014.3001.5502

- [ ] Go_web gin框架开发（通用脚手架）
https://blog.csdn.net/qq_43514659/article/details/121587913?spm=1001.2014.3001.5502

- [ ] zap日志的基本使用（go必会知识*）

	https://blog.csdn.net/qq_43514659/article/details/121567519?spm=1001.2014.3001.5502




3、使用Singo 搭建本地项目
	https://gourouting.github.io/

仓库地址：origin	https://github.com/SilhouetteLiang/singo.git (fetch)

liangpengzhan@outlook.com
Wslpz888888

	Personal access tokens:	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

	git remote set-url origin https://<your_token>@github.com/<USERNAME>/<REPO>.git


	git remote set-url origin https://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@github.com/SilhouetteLiang/singo.git


在/Users/pengzhanliang/go/src/gin-liang/singo执行  go run main.go

Redis  密码
sl2016SL!

	mysql账户密码		test  Lpz888888.
	修改mysql权限，
	1、需要想修改  vim /usr/local/etc/my.cnf  跳过权限表  
	2、然后使用 root登录  添加用户或者删除用户，执行flush privileges才能成功
	3、/usr/local/Cellar/mysql/8.0.17/bin执行mysql.server stop	加载修改后的usr/local/etc/my.cnf 文件
	4、vim /usr/local/etc/my.cnf  调整为需要权限表  


4、接口
	host：http://127.0.0.1:3000
	
		post：请求		

			1.检测是否可以ping通		/api/v1/ping
			2.注册接口				/api/v1/user/register
											nickname:nickname1
											user_name:user_name2
											password:123456
											password_confirm:123456
			3.登录接口				api/v1/user/login
								user_name
								password
			
			4、微信注册、手机号注册接口
			5、获取手机短信验证码接口(需要花钱购买第移动短信服务系统)
			6、
		get：请求
			
			1.关于我					v1/user/me
		

		delete：请求
			1.登出接口				v1/user/logout
			


												
									
					
使用算法实现一个机器人，关于
	心里咨询的
		实现一个算法给不同人推荐不同的心里咨询方案

	
	心理资讯：
		把客户自己浏览的内容、某一类内容的浏览时间 ======== 作为推荐内容的依据，推荐相关联的咨询。
		





	心里咨询

推荐算法
平台应用



























	
	















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