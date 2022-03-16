## 商城后端代码

### 新增单元测试

### 技术栈
- golang
- gin
- gorm
- viper读取文件 

### Feature
实现用户注册的接口

step

1、添加项目依赖go mod init，新建项目文件夹


2、定义常量
- 配置文件conf/config.yaml,读取配置文件config/config.go
- 枚举类、model、resp中的entity(返回数据用)

3、编写mvc
- repository(dao)、service、handler(前端传入参数解析成结构体，对返回结果用entity统一进行封装)

4、初始化和添加路由
- handler.go中对viper读取配置文件、数据库连接、handler+service+repo分别进行初始化
- 添加router，将请求url映射到对应的handler方法