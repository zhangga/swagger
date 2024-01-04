# swagger
A service that provides Swagger UI parsing, supporting openapi.yaml files.

通用的swagger解析服务，用于解析openapi.yaml文件，生成swagger ui页面。

特性
1. 通用服务，使用docker部署一个即可访问其他服务的openapi文件。
2. 支持openapi中多个请求源切换。
3. 支持url中指定host，方便分发。
4. 支持跨域访问。

本地测试
* `make run`
* 访问 `http://127.0.0.1:8081/swagger-ui/` 使用
* 指定openapi.yaml文件的host，访问 `http://127.0.0.1:8081/swagger-ui/${YAML_HOST}/` ，注意`/`结尾
