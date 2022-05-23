创业本地Go项目

---------------------------------------------------------------------------------------------------------------------
Linux云服务器操作
    ssh root@103.153.138.60 
    cGK{8*uv6Uc<

1.查看进程pid             四种方法：https://blog.csdn.net/weixin_31960565/article/details/112804008
    pidof singo
    pgrep singo
    pstree -p | grep "singo"
    ps aux | grep "singo"

2.后台挂起
    nohup ./singo & 
    ps aux|grep singo

---------------------------------------------------------------------------------------------------------------------

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

docker删除镜像
	docker rm 
	1.首先用docker ps -a 列出所有容器 可以看到这里有两个mysql正在运行 我们想要删除镜像 就必须先删除容器 不然就会报各种错误
	2.然后我们输入docker rm 7887634e4663删除容器
	3.找到镜像id docker images
	4.这时候输入 docker rmi 镜像id 就可以删除我们想要删除的镜像了

docker启动服务
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

GROM操作mysql数据库
	https://blog.csdn.net/qq_43514659/article/details/121466144?spm=1001.2014.3001.5502
Go_web gin框架开发（通用脚手架）
    https://blog.csdn.net/qq_43514659/article/details/121587913?spm=1001.2014.3001.5502
zap日志的基本使用（go必会知识*）
    https://blog.csdn.net/qq_43514659/article/details/121567519?spm=1001.2014.3001.5502

使用Singo 搭建本地项目
    liangpengzhan@outlook.com       Wslpz888888
    Redis  密码   sl2016SL!


在/Users/pengzhanliang/go/src/gin-liang/singo执行  go run main.go

mysql账户密码		test  Lpz888888.
修改mysql权限，
1、需要想修改  vim /usr/local/etc/my.cnf  跳过权限表  
2、然后使用 root登录  添加用户或者删除用户，执行flush privileges才能成功
3、/usr/local/Cellar/mysql/8.0.17/bin执行mysql.server stop	加载修改后的usr/local/etc/my.cnf 文件
4、vim /usr/local/etc/my.cnf  调整为需要权限表  

---------------------------------------------------------------------------------------------------------------------

V2版本
	使用算法实现一个机器人，关于
	心里咨询:
		实现一个算法给不同人推荐不同的心里咨询方案
	心理资讯：
		把客户浏览的内容、某一类内容的浏览时间 ======== 作为推荐内容的依据，推荐相关联的咨询。

推荐算法平台
实物平台应用

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
平台化＋算法推荐
---------------------------------------------------------------------------------------------------------------------

