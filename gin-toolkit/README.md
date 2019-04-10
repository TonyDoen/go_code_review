
### 启动&重启&隔离依赖
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

### 常用命令

#### 新建项目
`sh ./cg.sh 项目名称`

#### 更新服务端框架
`govendor fetch -v ...`

#### 启动依赖
`docker-compose up -d`
