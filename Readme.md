# 代码生成  
这一步我们聚焦生成代码。详细解释请参见我在B站的视频：[Kubernetes源码开发之旅四：Aggregated API Server](https://www.bilibili.com/video/BV1Ve4y1U7oE/?vd_source=9304721f2aeb71f0f883054e229f5b22)

## 增加标签  
我们需要在doc.go等代码中加入指导代码生成的标签。  

## 在代码生成中注入逻辑  
直接生成的代码有不符合我们要求的可能，特别是为api object设置default值和不同版本之间convert所生成的代码。我们可以自己写这些逻辑，以合适的名字去命名方法，从而让代码生成程序直接引用它们。  
参见pkg/apis/cicd/v1alpha/defaults.go和conversions.go  

## 编写脚本  
代码生成是调用了kubernetes的code-gen子工程中脚本完成的，我们需要自己的脚本去调用之。  
参见hack/文件夹下内容  

## 如何运行脚本  
```bash
cd <go path>/src/cicd-apiserver
./hack/code-generation.sh
```
注意一定是在本工程的主目录下去跑代码生成脚本。由于代码生成需要kubernetes的code-gen子工程，所以我们需要确保它的代码被放入了vendor目录中：  
```bash
cd <go path>/src/cicd-apiserver
go mod tidy
go mod vendor
```
那么vendor中就会有当前工程所引用的所有外部依赖工程。hack/tool.go下引用了code-gen，所以它会被copy进vendor子目录
