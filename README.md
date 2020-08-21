# Go-viper
![viper](https://github.com/spf13/viper/raw/master/.github/logo.png?raw=true)

[Viper](https://github.com/spf13/viper)是适用于Go应用程序的完整配置解决方案。它被设计用于在应用程序中工作，并且可以处理所有类型的配置需求和格式。

# Install

```go
go get github.com/spf13/viper
```

# What is Viper
Viper可以适配任何应用程序，可以处理所有类型的配置需求和格式。它支持:

- 设置默认值
- 读取JSON, TOML, YAML, HCL, envfile和Java属性配置文件
- 实时查看和重新读取配置文件(可选)
- 从环境变量中读取
- 从远程配置系统(etcd或Consul)读取，并观察变化
- 从命令行标志读取
- 读取缓冲区
- 设置明确的值

可以将Viper看作所有应用程序配置需要的注册中心。

# Read Config

[mian.go](./main.go)

# License
[MIT](./LICENSE)

Copyright (c) 2020 golang collection