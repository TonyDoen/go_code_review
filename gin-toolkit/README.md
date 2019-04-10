
### 一. 使用步骤
1. 配置 go 运行环境（包括 执行环境安装 $GOROOT, $GOPATH 变量配置）

2. 在 $GOPATH/src 目录下创建 github.com/TonyDoen/go_code_review/gin-toolkit 目录
```
mkdir -p $GOPATH/src/github.com/TonyDoen/go_code_review/gin-toolkit
```
3. 把 gin-toolkit 代码 移动到上面创建的目录
```
mv ./* $GOPATH/src/github.com/TonyDoen/go_code_review/gin-toolkit/
```
4. 新建项目
```
sh ./gen.sh 项目名称
```


### Tips
如需要修改 $GOPATH/src/github.com/TonyDoen/go_code_review/gin-toolkit/ 目录
1. 创建你自己的目录 $GOPATH/src/ `your/file/path` /gin-toolkit
2. 替换模板项目里的 默认路径
```
grep -rl 'github.com/TonyDoen/go_code_review' ./  | xargs sed -i "" "s/github.com\/TonyDoen\/go_code_review/your\/file\/path/g"
```


### 二. 启动&重启&隔离依赖
1. docker-compose up -d，会在本地启动所有的依赖环境（mysql/redis/nsq）

2. docker-compose restart [servicename] 如果不填写servicename则重启所有的服务，填写servicename重启指定的服务

3. 启动多个mysql
```
数据隔离
启动mysql
docker-compose run -d -p 3306:3306 mysql
再启动另外一个mysql
docker-compose run -d -p 3307:3306 mysql
启动redis
docker-compose run -d -p 6379:6379 redis
再启动另外一个redis
docker-compose run -d -p 6380:6379 redis
```

### 三. 常用命令

##### 更新服务端框架
`govendor fetch -v ...`

##### 启动依赖
`docker-compose up -d`


### 四. 目录说明

##### 项目根目录
如不修改默认目录，即 $GOPATH/src/github.com/TonyDoen/go_code_review/yourProjectName

##### 项目目录下通常存分为如下目录：
`common/`
通用组件目录（待优化）

`conf/`
配置文件目录

`controller/`
控制器层目录

`model/`
数据模型层目录

`router/`
路由层目录

`service/`
业务逻辑层目录

`sql/`
数据库生成文件

`vendor/`
依赖管理目录

