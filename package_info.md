
# Architecture overview

![Architecture overview](https://cdn.jsdelivr.net/gh/prometheus/prometheus@c34257d069c630685da35bcef084632ffd5d6209/documentation/images/architecture.svg)

# 源码包结构说明
* cmd main函数程序启动类，程序入口
* config 配置类主要包含配置方法，配置结构体
* consoles，console_libraries 控制台及其依赖
* discovery 用作服务发现的代码
* docs md帮助手册
* documentation 一些配置实例yaml的存放地
* model 定义prometheus所使用的结构体
* notifier 报警模块
* plugins 生成插件的包根据根目录的plugins.yml
* prompb //TODO
* **promql** prometheus sql 包 对应架构图中的PromQL
* rules 告警规则
* **scrape** pull metric数据 对应架构图中的retrieval
* scripts 脚本
* **storage tsdb** prometheus的核心，时序数据库的定义及实现
* template 查询模板
* tracing //TODO
* util 工具包
* web 前端对接接口