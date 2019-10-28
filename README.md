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
2. ~~日志记录到文件~~ 2019.10.25
3. mysql
4. redis
5. docker部署
6. ~~swagger文档~~ 2019.10.25
7. jwt


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

#### 日志记录到文件
- 在`config`指定`logPath`的值，`logPath`的值为日志目录
- info日志文件名称形如20191025.log，error日志文件名称形如20191025.log.wf
- 记录日志
    ```$golang
    // 记录info日志
    log.Info["hello"] = "world"
    log.Info["great"] = "wall"
    // 输出日志文本为： time="2019-10-25T18:29:39+08:00" level=info hello=world great=wall
    
    // 记录warning日志
    _ = log.Warn("You should probably take a look at this")
    // 输出日志文本为： time="2019-10-25T18:28:44+08:00" level=warning msg=You should probably take a look at this
    
    // 记录error日志
    _ = log.Error("Something failed but I'm not quitting.")
    // 输出日志文本为： time="2019-10-25T18:28:44+08:00" level=error msg=Something failed but I'm not quitting.
    ```
#### swagger文档
```$xslt
// main.go注释解释
@title 文档名称
@version 文档版本
@description 描述
@host 访问地址
@BasePath 例：/api/v1
@license.name Apache 2.0
@license.url http://www.apache.org/licenses/LICENSE-2.0.html
//示例：
// @title gin-project-template
// @version 1.0
// @description gin-project-template demo
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8080
// @BasePath /api/v1

// api注释解释
@Summary 简介
@Description 解释
@Tags api组名
@Produce  Response content type
@Param 参数名称 path/query/body 参数类型 是否必须true/false "参数解释"
@Success 200 {object} struct 用于展示success结果
@Failure 400 {object} struct 用于展示error结果
@Router 路由 [get/post/put/del/...]
```