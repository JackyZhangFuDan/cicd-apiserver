# 增加Admission  
Admission机制会在一个HTTP请求的处理过程后期被触发，主要用于对用户请求进行修改，对合理性进行校验。这对应于两个环节：  
- mutation阶段  
在这一阶段系统有可能去修改用户提交的请求，例如ISTIO的边车机制就在这个阶段向新创建的pod中注入边车container；  
- validation阶段  
这一阶段不能修改请求内容，只是对它的正确性进行校验，如果校验不过请求也会被拒绝。  

我的B站视频:[Kubernetes源码开发之旅四：Aggregated API Server](https://www.bilibili.com/video/BV1Ve4y1U7oE/?vd_source=9304721f2aeb71f0f883054e229f5b22)  

## 制作Admission Plugin  
一个Plugin是实现了admission.Interface接口的结构体，它可以同时实现mutaion阶段和validation阶段的逻辑，例如validation能力只需实现admission.ValidationInterface接口。代码：  
pkg/admission/plugin/jsplugins.go  

## 为Admission Plugin注入Informer  
在plugin实现自己逻辑的时候，往往是需要获取server中api object的信息的，例如我在示例逻辑中加了一个validation：系统不能创建超过10个jenkins service实例，那么我在Validate方法里就需要获取当前系统共有多少个js实例。在Aggregated API Server 中，（也包括API Server，都一样）不是通过直接search数据库来获取这些信息的，而是通过loopback informer。这里我们演示了如何向admission plugin注入informer。代码：  
pkg/admission/plugin/informerinjector.go  
