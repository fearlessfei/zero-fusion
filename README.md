# zero-fusion

zero-fusion是基于[go-zero](https://github.com/zeromicro/go-zero)框架之上封装多功能聚合的框架。

## 概述

通常框架只提供最基本的功能，业务的开发通常会用到很多中间件，这部分需要开发者根据业务需求去定制化实现并不能立马上手进行业务的开发，zero-fusion就是针对这个问题进行开发的脚手架，让开发只关注业务的开发无需关注业务以外的实现，真正做到开箱即用减少开发的心智负担。

### 特征

zero-fusion尽可能的进行高度的抽象，对于同一个中间件，模块可无缝的切换，例如定时任务模块高度抽象可以切换任意一个定时任务的实现而无需更改业务代码。

- 缓存key的封装，更直观的操作。
- 枚举自动生成，自动生成存储，解析db字段等方法。
- 定时任务服务，应对基本的任务调度场景。
- 内置集成swagger。
- nacos配置中心,其他配置中心逐步实现。
- 内置gorm/gen生成。
- gorm封装，支持主从配置。
- 脚本运行器。
- 。。。

### 规范

每个模块目录下都一个README.md文件用于描述模块的使用方法及注意事项。

### 示例

生成api

``````
goctl api go --type-group --api app/demo/apis/major/apis.api --dir app/demo --home ./goctl/template
``````

运行路由插件，每个服务可能对外对内提供不同的api地址，每部分需要单独部署，go-zero默认不支持。

``````
goctl api plugin -p serviceroute --api app/demo/apis/major/apis.api --dir app/demo
``````

生成swagger

``````
goctl api swagger --api app/demo/apis/major/apis.api --dir app/demo/docs/swagger --filename major
``````

### 最后

go-fusion还处于初级阶段，后续会不断完善增加其他功能模块，例如http,mq等模块的封装，文档后续会随着功能的不断增加会逐渐的完善。