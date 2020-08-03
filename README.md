# go-library

### Installation
Add `go-library` to your go.mod file
```
module project_name

go 1.14

require (
       github.com/CLannadZSY/go-library
)
```

### Example
1. `配置文件` 参考,  请不要修改配置文件中的字段名称, 否则将导致读取失败
   [`redis_config`](https://github.com/CLannadZSY/go-library/blob/master/database/redis/redis_config.yaml)

   [`mysql_config`](https://github.com/CLannadZSY/go-library/blob/master/database/sql/mysql_config.yaml)

2. 将配置文件复制到你的项目中

3. 使用方法, 请参考  `*_test.go` 文件

   [`redis`](https://github.com/CLannadZSY/go-library/blob/master/database/redis/redis_test.go)
   [`mysql`](https://github.com/CLannadZSY/go-library/blob/master/database/sql/mysql_test.go)


### Contributing
欢迎  `PR`  和   `ISSUES`

### Release History

* 2020-08-02: 完成```v1.0.0```版本 

### [License](https://github.com/CLannadZSY/go-library/blob/master/LICENSE)