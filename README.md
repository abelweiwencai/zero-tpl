# zero-tpl


## quick start

1. 数据库表生成，到数据库执行 `goddess-mall-server/database/goddess-mall.sql`

2. 生成model数据
    1. 修改GORMT的配置文件 `database/config.yaml`
    2. 生成model代码 `gormt -f database/config.yaml`
4. 修改配置文件 `apps/apigateway/etc/apigateway-api.yaml`
3. 安装依赖库
```
cd apps/apigateway
go mod tidy
```
4. 启动服务
```
go run apigateway.go -f etc/apigateway-api.yaml
```

## 新建APP流程
```
cd apps/
goctl api new app_name
cd app_name
go mod tidy
goctl api go -dir ./ -style goZero -api ./apiserver/apiserver.api
```

## reference
- go-zero官方文档：https://go-zero.dev/cn/monolithic-service.html