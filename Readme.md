# 部署API Server  

## 生成Docker镜像并推到仓库中  
```bash
docker build -t jackyzhangfd/cicd-kube-apiserver:1.0 .  
docker push jackyzhangfd/cicd-kube-apiserver:1.0  
```

设置集群，从而能从docker hub来拉取镜像  
https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/  
https://kubernetes.io/docs/concepts/containers/images/#using-a-private-registry  

