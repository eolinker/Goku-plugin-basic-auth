### 插件名称

| 类别 |  名称 |  字段  | 属性  |
| ------------ | ------------ | ------------ |------------ |
| 策略插件 | Basic鉴权 | goku-basic_auth  | 用户鉴权（静态token） |

### 功能描述

鉴权方式的一种，设置Basic Auth的Username与Password，不能通过认证的用户将无权访问接口。

### 配置页面

进入控制台 >> 策略管理 >> 某策略 >> 策略插件 >> Basic鉴权插件：

![](http://data.eolinker.com/course/6sjKz8Ge98cc1c5c35d9c9a980f618ad4762096e6569715)

### 配置参数

| 参数名 | 说明   | 
| ------------ | ------------ |  
|  userName | 用户名 | 
| password  | 密码 |
| hildCredential  | 转发时是否隐藏basicAuth验证信息 |

### 配置示例

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

### API请求参数

| 参数名 | 说明  | 必填  |   值可能性   |  参数位置 |
| ------------ | ------------ |  
|  Strategy-Id | 策略id | 是 |   |  header  | 
|  Authorization-Type  | 鉴权方式 | 是 | Basic  | header  |
| Authorization  |  鉴权值 |  是  |    | header |

###请求示例

###### 以下API测试页面来自于 **eoLinker AMS** 接口管理平台

![](http://data.eolinker.com/course/xQmXkmtadbe1da64748837e766decf0d21c64473a53409c)
