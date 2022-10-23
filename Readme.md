# 注册API Object
这一步我们聚焦在scheme信息的填充。API Object必须经过向scheme注册后才可能起作用。具体讲解参考我在B站上的视频:[Kubernetes源码开发之旅四：Aggregated API Server](https://www.bilibili.com/video/BV1Ve4y1U7oE/?vd_source=9304721f2aeb71f0f883054e229f5b22)

## 向Scheme注册  
Scheme就像一个注册表，它会包含server支持的所有API Object以及它们的重要操作函数。需要注意的是Object的每个version都需要进行注册。由于Kubernete Generic API Server的存在，我们向scheme注册的动作变得十分简单。具体参见：  
pkg/apis/cicd/register.go  
pkg/apis/cicd/v1alpha/register.go  
它们分别注册internal version和v1alpha version；  
pkg/apis/cicd/install/install.go  
这个文件提供方法集中调用各个version的注册操作，server在启动过程中会调用它
