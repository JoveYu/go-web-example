# 技术选择

| 功能        | 项目                    | 优点                                         | 备注                       |
| ----------- | ----------------------- | -------------------------------------------- | -------------------------- |
| http server | gin                     | 功能强大，扩展相对多，API完善                |                            |
| http client | 无                      |                                              | 没有特别好用的             |
| orm         | gorm                    | API简介，同时支持ORM和sqlbuilder             |                            |
| docs        | swag                    | 注释方式生成文档，额外操作少，自动解析结构体 |                            |
| log         | zerolog                 | 高性能，0内存分配，结构化                    | 缺点，使用需要指定参数类型 |
| config      | viper                   | 功能强大，支持多种方式获取配置               |                            |
| metrics     | opentelemetry+prmetheus | 公认metrics方案                              | 待实现                     |
| trace       | opentelemetry+jaeger    | 公认trace方案                                | 待实现                     |

# 文档地址

`http://127.0.0.1:8080/swagger/index.html` 文档查看网页

`http://127.0.0.1:8080/swagger/doc.json` swaggerjson，可以自动导入到YAPI

# 命令

- `make swagger` 生成文档
- `make run`  运行