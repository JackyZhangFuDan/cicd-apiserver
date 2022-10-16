# 部署API Server  

编写dockerfile，以及必要的API Object定义文件，例如namespace, service account, role binding, deployment, service, apiservice...等等  

## 生成Docker镜像并推到仓库中  
```bash
cd ~/go/src/cicd-apiserver
docker build -t jackyzhangfd/cicd-kube-apiserver:1.0 .  
docker push jackyzhangfd/cicd-kube-apiserver:1.0  
```

## 测试API Server镜像的正确性  
```bash
cd ./artifacts/deploy
kubectl apply -f .  
kubectl describe pod <我们的api server所在pod的名字>
```
第一条命令会创建出部署涉及到的所有API Objects，例如namesapce等，大多数不会有问题；第二条命令重点关注API Server所在的pod：首先看pod中镜像的拉取是否成功，不成功的可能需要额外配置；其次看各个container是否启动成功了。由于代码bug等，一般是我们自己api server的那个container启动会出错，为了看启动失败的原因需要：  
```bash
kubectl logs <我们的api server所在pod的名字> cicd-apiserver
```
来看那个container的log信息  

Tips：设置集群，从而能从docker hub来拉取镜像  
https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/  
https://kubernetes.io/docs/concepts/containers/images/#using-a-private-registry  
（docker hub在minikube环境没问题,不用配置就可拉取）  

当我们的pod，service都启动好了以后，看一看自定义API Resource “jenkinsservices”出没出现在server的resource列表中：  
```bash
kubectl get api-resources
```

## 创建API Object 实例  
作为测试，我们创建一下Aggregated API Server中定义的API Object ‘JenkinsService’：  
```bash
cd ./artifacts/testcases
kubectl apply -f .
```
注意：你可能需要等Server启动完全后再执行创建命令，否则它还没有准备好呢。最好等5分钟  

上述命令会创建出10个JenkinsService实例，第11个会失败，因为我们在admission里控制了数量，这里的失败是正确的。成功后，可以用如下命令可以列出创建成功的事例：  
```bash
kubectl get jenkinsservices
```