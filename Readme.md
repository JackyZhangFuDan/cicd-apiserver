# API Object的存储  
API Server需要把支持的API Object实例存入DB - ETCD - 中，Generic API Server已经帮我们的Aggregated API Server完成了大部分繁琐复杂的数据库相关处理，我们只要实现一些Interface来注入我们的逻辑就ok了。讲解参见我的B站视频<...>  
代码全部在pkg/registry文件夹内  

## 实现数据库存储时各个策略的接口  
在数据库创建，修改API Object的时候，会调用我们的一些列strategy接口方法，主要是create策略和update策略，这些策略里包含的就是我们自己的校验逻辑，我们需要先实现它们。代码在:  
pkg/registry/cicd/jenkinsservice/strategy.go  

## 制作Store实例  
Store负责和ETCD打交道，存取object。在API Server代码中，数据库和Restful服务绑定的比较深入，所以这里我们会再把store包裹一下，形成一个REST结构体实例，可以想象，它会最终响应http restful请求。  
代码：pkg/registry/cicd/jenkinsservice/etcd.go  

