# 第四周

## 工程项目结构

### Standard Go Project Layout

    开发工具创建目录，来推动规范

    
####  /cmd

```
|-cmd
|  |
|  |-demo1
|  |   |-main.go
|  |-demo2
|  |   |-main.go   
```

每个应用程序的目录名应该与想要的可执行文件的名称相匹配

cmd目录中不要放太多代码

#### /internal

存放所有的私有的应用程序和库代码，不希望其他人在其应用程序或库中导入的代码

* 在编译层进行隔离，参考Go 1.4 release notes

#### /pkg

```
 |-pkg
 |  |-cache
 |  |   |-redis
 |  |   |-memcache
 |  |-conf
 |  |   |-dsn
 |  |   |-env
```

存放外部应用程序可以使用的库代码，其他项目会导入这些库。

### Kit Project Layout （基础库）

Package Origented Design

Kit项目必须具备的特点
* 统一
* 标准库方式布局
* 高度抽象
* 支持插件

Kit库不应使用vendor

package命名方法 //todo

### Service Application Project Layout

```
|-README.md
|-api
|-cmd
|-configs
|-go.mode
|-go.sum
|-internal
|-test
```

* /api
  * API协议定义目录，protobuf文件，以及生成的go文件
* /configs
  * 配置文件模板或默认配置
* /test
  * 外部测试应用程序和测试数据


#### 微服务管理

* 微服务需要一个全局唯一的名称
* 公司需要维护一个服务树来根据应用名称来管理微服务
* 建议采用`三段式命名法,按业务来命名`


## API 设计

### gRPC

## 配置管理

## 包管理

## 测试

## 引用
https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html

https://www.ardanlabs.com/blog/2017/02/design-philosophy-on-packaging.html

https://github.com/golang-standards/project-layout

https://github.com/golang-standards/project-layout/blob/master/README_zh.md

https://www.cnblogs.com/zxf330301/p/6534643.html

https://blog.csdn.net/k6T9Q8XKs6iIkZPPIFq/article/details/109192475?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522160561008419724839224387%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=160561008419724839224387&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_v2~rank_v28-6-109192475.first_rank_ecpm_v3_pc_rank_v2&utm_term=阿里技术专家详解DDD系列&spm=1018.2118.3001.4449

https://blog.csdn.net/chikuai9995/article/details/100723540?biz_id=102&utm_term=阿里技术专家详解DDD系列&utm_medium=distribute.pc_search_result.none-task-blog-2~all~sobaiduweb~default-0-100723540&spm=1018.2118.3001.4449

https://blog.csdn.net/Taobaojishu/article/details/101444324?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522160561008419724838528569%2522%252C%2522scm%2522%253A%252220140713.130102334..%2522%257D&request_id=160561008419724838528569&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~top_click~default-1-101444324.first_rank_ecpm_v3_pc_rank_v2&utm_term=阿里技术专家详解DDD系列&spm=1018.2118.3001.4449

https://blog.csdn.net/taobaojishu/article/details/106152641

https://cloud.google.com/apis/design/errors

https://kb.cnblogs.com/page/520743/