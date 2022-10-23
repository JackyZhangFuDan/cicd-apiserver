# 创建项目，定义API Object  
在开始动手编码前，我们需要设计我们的Aggregated API Server（我有时也叫它 Custom API Server）所支持的API Objects，包括这些API Object所具有的属性，它们决定了通过kubectl操作时yaml文件长相。  

对于这个示例Server，它会支持一个API Object：JenkinsService，你可以把一个JS实例看作一组Jenkins实例的集合，用于处理CICD相关工作（注意，在本示例Server里面我不会真正在集群中去创建Jenkins，只是把JenkinsService API Object在Aggregated API Server中定义出来，实际创建需要另建Controller去完成）；它的spec里只有InstanceAmount和InstanceCpu两个属性，分别代表需要多少个Jenkins实例，以及每个Jenkins实例需要多少cpu核心。  

## 我们首先需要初始化一个工程
```bash
cd <go path>/src
mkdir cicd-apiserver
cd cicd-apiserver
go mod init cicd-apiserver
```

你也可以直接fork这个分支，然后以它为基础改成你的project，同样ok  

## 开始编码  
主要部分在文件夹pkg/apis/cicd下。
请参考我在B站上的视频:[Kubernetes源码开发之旅四：Aggregated API Server](https://www.bilibili.com/video/BV1Ve4y1U7oE/?vd_source=9304721f2aeb71f0f883054e229f5b22)
