# 创建Aggregated API Server  

众所周知Kubernetes提供了诸多扩展API Server的方式，比较常用的一种是CRD - Customer Resource Definition，它已经可以满足绝大多数的扩展需求。而通过自定义Aggregarted API Server的方式可以获得最大程度的灵活性，只是门槛有点儿高。这个仓库包含了创建Kubernete Aggregated API Server的代码，它演示了如何从头开始建立自己的API Server，放入集群来扩展API Server。  

注意：你需要仔细思考CRD是否已经足以满足你的需求，不要一上来就走Aggregated API Server这条路。  

我通过不同分支来展示server创建的整个过程，每个分支代表一个主题，顺序不能乱因为前后有依赖关系。通过我在B站上的相关视频你可以更好理解这些代码： <....>

## Branch master  
最新的完整代码  

## Branch [phase-1](https://github.com/JackyZhangFuDan/cicd-apiserver/tree/phase-1/)  
初始化工程，建立所需要的API Object  

## Branch [pahse-2](https://github.com/JackyZhangFuDan/cicd-apiserver/tree/phase-2/)  
代码生成，根据tag为建立的API Object生成配套代码  

## Branch [phase-3](https://github.com/JackyZhangFuDan/cicd-apiserver/tree/phase-3/)  
向scheme注册API Object  

## Branch [phase-4](https://github.com/JackyZhangFuDan/cicd-apiserver/tree/phase-4/)  
存储API Object到ETCD  

## Branch [phase-5](https://github.com/JackyZhangFuDan/cicd-apiserver/tree/phase-5/)  
增加Admission到Aggregated API Server  

## Branch [phase-6](https://github.com/JackyZhangFuDan/cicd-apiserver/tree/phase-6/)  
添加代码生成web server  

## Branch [phase-7](https://github.com/JackyZhangFuDan/cicd-apiserver/tree/phase-7/)  
部署Aggregated API Server到集群，并启动  
