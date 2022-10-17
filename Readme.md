# 创建Web Server  
经过前序步骤，我们把Aggregated API Server的内容构建起来了，现在我们需要给它套上一层web server的包装，那么它就可以作为一个web server去响应http 请求了。  
每个Aggregated API Server都会基于Kubernetes的子项目apiserver（也叫generic server），Kubernetes的API Server自己也是如此。generic server内含了一个Aggregated API Server 90%的逻辑，它也包含了web server的基础设施所以我们并不需要做太多工作。  

## 制作Server  
虽然Generic Server提供了大部分的基础设施，我么你还是需要把它包裹一下，形成自己的API Server实例，其中最为重要的环节是把我们之前制作的API Object注入到Web Server中，从而Web Server能响应关于这些object的Restful请求。这也是我们注入自己特有配置的时机，只是我们在示例程序中没有加入任何自有config。  
制作的整个过程大概就是：制定config结构体（其中内嵌了generic config）-> 写进一步完善config数据的方法 -> 由完善后的config结构体创建出Server实例。代码：  
pkg/apiserver/apiserver.go  

注意：我们在前序阶段中制作了API Object向Scheme注册的install方法，该方法的调用也是在这个包（apiserver）的init过程中。也就是说，当我们apiserver被引入时，init方法就会被执行，就会保证涉及到的api object都被注册。  

## 启动Server  
程序接收“用户”的命令行参数，按照这些参数的指示去启动刚刚制作的Server。逻辑上工作过程是：收到命令行参数，形成“options”结构体实例 -> 校验和完善options结构体 -> 生成server config结构体实例 -> complete 该config -> 由config生成server实例 -> 启动server。Kubernetes钟爱cobra命令行框架，在我们的aggregated api server代码中也是用的这个框架做用户终端的，便马上来说就是先要做一个cobra command对象，然后再main方法里去调用这个command对象启动server。 具体代码：  
- cmd/server/server.go 制作command对象  
- main.go 调用command启动server  

注意：我们之前制作的admission plugin以及admission用到的informer都是在制作config和启动server时加入进来的。
