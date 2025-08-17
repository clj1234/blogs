# 项目依赖

​	1.MySQL

​	2.Gorm

​	3.Gin

# 项目结构

```
blogs/
├── server
    └── api/		`web接口`
	│── config/		`配置文件字段映射，xxx.yaml`
	│── global/   
	│── middleware/	`gin中间件`
	│── model/		`数据库与请求参数映射`
	│── router/		`接口路由`
	│── service/	`业务实现`
	│── utils/
    │── main.go   
    │── go.mod
    │── config.yaml	`项目配置参数`
```



# 依赖导入

```go
go get github.com/golang-jwt/jwt/v5
go get -u github.com/gin-gonic/gin
go get gopkg.in/yaml.v3
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u gorm.io/driver/mysql
```



# 项目启动

### 修改配置文件

```go
/server/config.yaml
```

### 启动项目
进入server目录下运行
```go
go run main.go
```

#测试接口

接口文档连接：https://docs.apipost.net/docs/4e1dec7444a3000?locale=zh-cn