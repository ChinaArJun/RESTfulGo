# RESTfulGo

## 手把手教你整合最简洁的GO开发框架：gin + grom + jwt
## Enterprise-level RESTful API services Project in Go 

> 运行打开 http://localhost:7777/swagger/index.html 查看接口文档

## 项目用到的技术
```
Http (gin)          #网络框架
middleware (gin)    #gin中间件实现请求日志和请求拦截
Token（jwt）        #API身份效验
Mysql (gorm)       #数据库
FileConf (Viper)   #配置文件读取
CodeBcrypt (bcrypt) #哈希密码加密
APIDocSwagger (gin-swagger) #API文档
TTestAPI (testing)   #测试框架
APIPProf (go/PProf)  #接口性能测试
Makefile        #管理API项目
```

### 目录结构
```
├── admin.sh                     # 进程的start|stop|status|restart控制文件
├── conf                         # 配置文件统一存放目录
│   ├── config.yaml              # 配置文件
│   ├── server.crt               # TLS配置文件
│   └── server.key
├── config                       # 专门用来处理配置和配置文件的Go package
│   └── config.go                 
├── db.sql                       # 在部署新环境时，可以登录MySQL客户端，执行source db.sql创建数据库和表
├── docs                         # swagger文档，执行 swag init 生成的
│   ├── docs.go
│   └── swagger
│       ├── swagger.json
│       └── swagger.yaml
├── handler                      # 类似MVC架构中的C，用来读取输入，并将处理流程转发给实际的处理函数，最后返回结果
│   ├── handler.go
│   ├── sd                       # 健康检查handler
│   │   └── check.go 
│   └── user                     # 核心：用户业务逻辑handler
│       ├── create.go            # 新增用户
│       ├── delete.go            # 删除用户
│       ├── get.go               # 获取指定的用户信息
│       ├── list.go              # 查询用户列表
│       ├── login.go             # 用户登录
│       ├── update.go            # 更新用户
│       └── user.go              # 存放用户handler公用的函数、结构体等
├── main.go                      # Go程序唯一入口
├── Makefile                     # Makefile文件，一般大型软件系统都是采用make来作为编译工具
├── model                        # 数据库相关的操作统一放在这里，包括数据库初始化和对表的增删改查
│   ├── init.go                  # 初始化和连接数据库
│   ├── model.go                 # 存放一些公用的go struct
│   └── user.go                  # 用户相关的数据库CURD操作
├── pkg                          # 引用的包
│   ├── auth                     # 认证包
│   │   └── auth.go
│   ├── constvar                 # 常量统一存放位置
│   │   └── constvar.go
│   ├── errno                    # 错误码存放位置
│   │   ├── code.go
│   │   └── errno.go
│   ├── token
│   │   └── token.go
│   └── version                  # 版本包
│       ├── base.go
│       ├── doc.go
│       └── version.go
├── README.md                    # API目录README
├── router                       # 路由相关处理
│   ├── middleware               # API服务器用的是Gin Web框架，Gin中间件存放位置
│   │   ├── auth.go              # 
│   │   ├── header.go            # 统一请求头
│   │   ├── logging.go           # 日志记录中间件
│   │   └── requestid.go         # 生成请求ID中间件
│   └── router.go
├── service                      # 实际业务处理函数存放位置
│   └── service.go
├── util                         # 工具类函数存放目录
│   ├── util.go 
│   └── util_test.go
└── vendor                         # vendor目录用来管理依赖包
    ├── github.com
    ├── golang.org
    ├── gopkg.in
    └── vendor.json
```

## License

[MIT license](http://opensource.org/licenses/MIT) © [ArJun](https://github.com/ArJun)