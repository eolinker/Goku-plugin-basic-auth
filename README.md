# Goku Plugin：Basic Auth

| 插件名称  | 文件名.so |  插件类型  | 错误处理方式 | 作用范围 |  优先级  |
| ------------ | ------------ | ------------ | ------------ | ------------ | ------------ |
| Basic鉴权  | goku-basic_auth | 访问策略 | 继续后续操作 | 转发前  | 1001 |

鉴权方式的一种，设置Basic Auth的Username与Password，不能通过认证的用户将无权访问接口。

# 目录
- [编译教程](#编译教程 "编译教程")
- [安装教程](#安装教程 "安装教程")
- [使用教程](#使用教程 "使用教程")
- [更新日志](#更新日志 "更新日志")

# 编译教程

#### 环境要求
* 系统：基于 Linux 内核（2.6.23+）的系统，CentOS、RedHat 等均可；

* golang版本号：12.x及其以上

* 环境变量设置：
	* GO111MODULE：on
	
	* GOPROXY：https://goproxy.io


#### 编译步骤

1.clone项目

2.进入项目文件夹，执行**build.sh**
```
cd goku-basic_auth && chmod +x build.sh && ./build.sh
```

###### 注：build.sh为通用的插件编译脚本，自定义插件时可以拷贝直接使用。

3.执行第2步将会生成文件： **{插件名}.so**

将该文件上传到**节点服务器运行目录**下的**plugin**文件夹，然后在控制台安装插件即可使用。

# 安装教程
前往 Goku API Gateway 官方网站查看：[插件安装教程](url "https://help.eolinker.com/#/tutorial/?groupID=c-341&productID=19")

# 使用教程

#### 配置页面

进入控制台 >> 策略管理 >> 某策略 >> 策略插件 >> Basic鉴权插件：

![](http://data.eolinker.com/course/6sjKz8Ge98cc1c5c35d9c9a980f618ad4762096e6569715)

#### 配置参数

| 参数名 | 说明   | 
| ------------ | ------------ |  
|  userName | 用户名 | 
| password  | 密码 |
| hildCredential  | 转发时是否隐藏basicAuth验证信息 |

#### 配置示例

```
[
    {
        "userName": "name",
        "password": "password",
        "hideCredentials": false,
        "remark": ""
    },
    {
       "userName": "name2"
       "password": "password2",
        "hideCredentials": false,
        "remark": "",
    }
]
```

#### API请求参数

| 参数名 | 说明  | 必填  |   值可能性   |  参数位置 |
| ------------ | ------------ |  ------------ | ------------ | ------------ | 
|  Strategy-Id | 策略id | 是 |   |  header  | 
|  Authorization-Type  | 鉴权方式 | 是 | Basic  | header  |
| Authorization  |  鉴权值 |  是  |    | header |

###请求示例

###### 以下API测试页面来自于 **eoLinker API Studio** 接口管理平台

![](http://data.eolinker.com/course/xQmXkmtadbe1da64748837e766decf0d21c64473a53409c)
