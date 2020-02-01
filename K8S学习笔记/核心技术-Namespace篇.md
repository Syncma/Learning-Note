# NameSpace篇
<!-- TOC -->

- [NameSpace篇](#namespace%e7%af%87)
  - [介绍](#%e4%bb%8b%e7%bb%8d)
  - [使用](#%e4%bd%bf%e7%94%a8)

<!-- /TOC -->
## 介绍

用于实现多用户的资源隔离，通过将集群内部的资源分配到不同的Namespace中形成逻辑上的分组

便于不同的分组在共享使用整个集群的资源，同时也可以被分别管理



 k8s里面默认的存在的名称空间有: default, kube-system, kube-public

```
[root@laptop tmp]# kubectl get namespaces
NAME                   STATUS   AGE
default                Active   23h
kube-node-lease        Active   23h
kube-public            Active   23h
kube-system            Active   23h
kubernetes-dashboard   Active   23h

[root@laptop tmp]# kubectl get pods
NAME                          READY   STATUS    RESTARTS   AGE
hello-world-f7dbcbd8f-ghq2r   1/1     Running   2          21h
nginx-8wxrn                   1/1     Running   0          100m

[root@laptop tmp]# kubectl get pods --namespace=default
NAME                          READY   STATUS    RESTARTS   AGE
hello-world-f7dbcbd8f-ghq2r   1/1     Running   2          21h
nginx-8wxrn                   1/1     Running   0          99m
```



**如果不特别指定namespace， 用户创建的pod，RC，service都将被系统分配到默认的default Namespace中**



## 使用

通过yaml文件里面增加namespace字段来创建



简单例子：

namespace.yaml

```yaml
apiVersion: v1
kind: Namespace
metadata:
 name: development
```





demo.yaml

```yaml
apiVersion: v1
kind: Pod
metadata:
 name: busybox
 namespace: development
spec:
 containers:
 - image: busybox
   name: busybox
```



创建namespace:

```
[root@laptop tmp]# kubectl create -f namespace.yaml 
namespace/development created

[root@laptop tmp]# kubectl get namespaces
NAME                   STATUS   AGE
default                Active   23h
development            Active   20s
kube-node-lease        Active   23h
kube-public            Active   23h
kube-system            Active   23h
kubernetes-dashboard   Active   23h
```



创建Pod:

```
[root@laptop tmp]# kubectl create -f demo.yaml 
pod/busybox created

[root@laptop tmp]# kubectl get pods --namespace=development
NAME      READY   STATUS      RESTARTS   AGE
busybox   0/1     Completed   2          49s
```
