# 部署API Server  

## 生成Docker镜像并推到仓库中  
```bash
docker build -t jackyzhangfd/cicd-kube-apiserver:1.0 .  
docker push jackyzhangfd/cicd-kube-apiserver:1.0  
```

## 测试API Server镜像的正确性  
```bash
kubectl apply -f ./artifacts/0-verify-image-repository.yaml  
kubectl describe pod test-cicd-apiserver-pod
```
首先看pod中镜像的拉取是否成功，不成功的可能需要额外配置；其次看各个container是否启动成功了。由于代码bug等，一般是我们自己api server的那个container启动会出错，为了看启动失败的原因需要：  
```bash
kubectl logs test-cicd-apiserver-pod cicd-apiserver
```
来看那个container的log信息  

Tips：设置集群，从而能从docker hub来拉取镜像  
https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/  
https://kubernetes.io/docs/concepts/containers/images/#using-a-private-registry  
（docker hub在minikube环境没问题,不用配置就可拉取）  

