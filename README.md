# dy 字节青训营项目

该项目是字节青训营的的最终大项目，通过该项目我也是学会了很多东西

## 项目启动

- 安装go语言环境
- 执行sql建立表语句
- 下载项目所需要的依赖

```shell
go mod tidy
```

- 编译项目

````shell
go build
````

- 运行项目

~~~shell
.\dy.exe
~~~



## 项目完成情况

完成了基础接口和互动接口俩部分

- ### 基础接口

  - #### /douyin/feed/ - 视频流接口

  - #### /douyin/user/register/ - 用户注册接口

  - #### /douyin/user/login/ - 用户登录接口

  - #### /douyin/user/ - 用户信息

  - #### /douyin/publish/action/ - 视频投稿

  - #### /douyin/publish/list/ - 发布列表

- ## 互动接口

  - #### /douyin/favorite/action/ - 赞操作

  - #### /douyin/favorite/list/ - 喜欢列表

  - #### /douyin/favorite/list/ - 喜欢列表

  - #### /douyin/comment/list/ - 视频评论列表