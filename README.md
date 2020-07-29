# 技术选择

| 功能        | 相关库                   | 优点                                         | 备注                                       |
| ----------- | ------------------------ | -------------------------------------------- | ------------------------------------------ |
| http server | gin                      | 功能强大，扩展相对多，API完善                |                                            |
| http client | resty                    | 基本功能完善，API比较简单，支持解析json      |                                            |
| orm         | gorm                     | API简洁，同时支持ORM和sqlbuilder             |                                            |
| docs        | swag                     | 注释方式生成文档，额外操作少，自动解析结构体 |                                            |
| log         | zerolog                  | 高性能，0内存分配，结构化                    | 缺点，为了不解析类型，使用需要指定参数类型 |
| config      | viper                    | 功能强大，支持多种方式获取配置               |                                            |
| metrics     | opentelemetry+prometheus |                                              |                                            |
| trace       | opentelemetry+jaeger     |                                              | 待实现                                     |

# 配置

`debug` 开启调试模式
1. 会同时开启gin的debug
2. 会开启log的debuglevel
3. 会修改log的输出格式

# 地址

`/swagger/index.html` 文档查看网页

`/swagger/doc.json` swaggerjson，可以自动导入到YAPI

`/metrics` 指标地址，提供给prometheus采集

# 命令

- `make swagger` 生成文档
- `make run`  运行
- `make build`  编译
- `make watch`  运行并随代码变化自动重启
- `make docker.build` 编译docker镜像
- `make docker.run` 运行docker镜像
- `make docker.push` 上传docker镜像