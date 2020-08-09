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
1. `配置文件` 参考,  请不要修改配置文件中的字段名称, 否则将导致读取失败<br>

   [`redis_config`](https://github.com/CLannadZSY/go-library/blob/master/database/redis/redis_config.yaml)<br>

   [`mysql_config`](https://github.com/CLannadZSY/go-library/blob/master/database/sql/mysql_config.yaml)

2. 将配置文件复制到你的项目中

3. 使用方法, 请参考  `*_test.go` 文件

   [`redis`](https://github.com/CLannadZSY/go-library/blob/master/database/redis/redis_test.go)<br>
   
   [`mysql`](https://github.com/CLannadZSY/go-library/blob/master/database/sql/mysql_test.go)


### Contributing
欢迎  `PR`  和   `ISSUES`

### [Release History](https://github.com/CLannadZSY/go-library/releases)

### [License](https://github.com/CLannadZSY/go-library/blob/master/LICENSE)
