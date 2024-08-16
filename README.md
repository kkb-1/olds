# 养老系统-后台管理（后端）

## 项目介绍

基于go-zero、gorm、go-redis/v9、zap、etcd的后台管理系统后端

> 特点

* go-zero 是一个集成了各种工程实践的 web 和 rpc 框架。通过弹性设计保障了大并发服务端的稳定性，经受了充分的实战检验。（官方原话）

* 使用gorm适配mysql、postgre

* 用redis作为缓存层，设计有空缓存、热点key续期等缓存策略

* 使用同一proto文件生成client即可跨语言调用rpc

* 使用了etcd作为服务中心
* 采用了bff-service分离策略，以api作为bff层，rpc作为service层
* 总体侧重考虑高并发情况下的服务高负载问题
* 采用了go-zero官方的设计规范，api文件即可直观感受到接口的介绍与细节
* zap定制化（这里感谢**djy**大佬的代码分享让我参考学习），zap集成进gorm
* 实现了rpc、http拦截器自定义并集成进了go-zero中（这里感谢b站up主**爱喝冰美式的程序员**的go实战视频 借鉴参考）

> 预期效果

* 使用kafka、canal、es，来使es同步mysql的数据，达到数据一致性的同时，利用es作为搜索引擎，承担用户管理模块的搜索功能。
* 后续考虑接入普罗米修斯等服务治理可视化的工具（视实际情况而定）
* 书写docker-compose.yaml提供环境参考、并接入k8s
* 将zap集成进go-zero

## 作者后话

接口功能比较简单，主要用作学习。

算是麻雀虽小五脏俱全吧

其实写了好久了，主要是一直在搭建项目demo（因为之前学了gin、hertz，都整了一下demo，后面发现go-zero更符合我的预期效果，就转用了go-zero，又是从0开始的一天捏QAQ）