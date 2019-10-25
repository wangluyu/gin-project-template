# gin-project-template
a project template for gin

## project structure
```
┌── api
│   ├── v1
├── model
├── config
│   ├── production.yaml
│   ├── development.yaml
│   └── test.yaml
├── middleware
├── pkg
├── router
├── script
├── build
│   ├── Dockerfile
│   └── docker-compose.yaml
├── doc
├── test
├── vendor
├── main.go
├── Makefile
└── README.md
```

## Todo list
1. ~~配置文件~~ 2019.10.25
2. 日志
3. docker部署
4. swagger文档
5. jwt


## 如何使用

#### 配置文件
- 在`config/app.yaml`指定`evn`的值，`evn`的值为配置文件名称
- 配置文件每增加一个配置项，需要在`pkg/util/conf.go`里添加/修改对应的struct
- 获取配置项
    ```$golang
    // 获取app配置
    appConf, err := util.FetchAppConf()
    // 获取自定义配置
    conf, err := util.FetchConf()
    mysqlConf := conf.Mysql
    mysqlHost := conf.Mysql.Host // mysqlConf.Host
    ```