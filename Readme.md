# 创建控制器  

这部分会为API JenkinsService创建控制器。为了方便，这个控制器没有做成单独可执行程序，而是效仿扩展Server，把它合并在了API Server可执行程序内。  
这个控制器是示意性质的，点到为止。它能达到的效果：为每个JenkinsService API实例创建出一个Deployment，跑NGINX这个镜像。为了使得它能够工作，还需要给本Aggregated Server的Service Account授予操作Deployment的权限，这是在rolecluster.yaml文件中进行的，需要将改变应用到集群。