# Service篇
<!-- TOC -->

- [Service篇](#service%e7%af%87)
  - [介绍](#%e4%bb%8b%e7%bb%8d)
  - [Service定义](#service%e5%ae%9a%e4%b9%89)
  - [基本用法](#%e5%9f%ba%e6%9c%ac%e7%94%a8%e6%b3%95)
  - [引出服务概念](#%e5%bc%95%e5%87%ba%e6%9c%8d%e5%8a%a1%e6%a6%82%e5%bf%b5)
  - [多端口服务](#%e5%a4%9a%e7%ab%af%e5%8f%a3%e6%9c%8d%e5%8a%a1)
  - [外部服务service](#%e5%a4%96%e9%83%a8%e6%9c%8d%e5%8a%a1service)

<!-- /TOC -->
## 介绍

通过创建service，可以为一组具有**相同功能的容器应用提供一个统一的入口地址**

并且将请求负载分发到后端的各个容器应用上



## Service定义

通过yaml文件定义



## 基本用法

一般来说，对外提供服务的应用程序需要通过某种机制来实现

对于容器应用最简单的方式是通过TCP/IP机制及监听IP和端口号来实现

例子：

rc-demo.yaml

```yaml
apiVersion: v1
kind: ReplicationController
metadata:
 name: myweb
spec:
 replicas: 1 
 template:
  metadata:
   name: myweb
   labels:
    app: myweb
  spec:
    containers:
    - name: myweb
      image: tomcat
      ports:
      - containerPort: 8080
```



创建rc:

```
[root@laptop tmp]# kubectl create -f rc-demo.yaml 
replicationcontroller/myweb created

[root@laptop tmp]# kubectl get pods
NAME                          READY   STATUS    RESTARTS   AGE
hello-world-f7dbcbd8f-ghq2r   1/1     Running   2          22h
myweb-mvp5n                   1/1     Running   0          38s
nginx-8wxrn                   1/1     Running   0          130m
```



查看myweb IP端口信息：

```
[root@laptop ~]# kubectl get pods -l app=myweb -o yaml |grep podIP
    podIP: 172.17.0.9
    podIPs:
```



测试：

```
[root@laptop ~]# curl 172.17.0.9:8080
```



## 引出服务概念

上面的例子是直接访问Pod的IP地址来访问应用服务的，这是不可靠的

因为当Pod所在的Node发生故障时，Pod会被k8s重新调度到另外一台Node，Pod地址也会发生变化

所以我们可以通过配置文件定义service，再通过kubectl create 创建

可以通过service地址来访问后端的Pod



例子：

service.yaml

```yaml
apiVersion: v1
kind: Service
metadata:
 name: myweb-svc
spec:
 ports:
 - port: 8081
   targetPort: 8080
 selector:
  app: myweb

```



创建服务：

```
[root@laptop tmp]# kubectl create -f service.yaml 
service/myweb-svc created
```



查看服务：

```
[root@laptop tmp]# kubectl get svc
NAME          TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
hello-world   NodePort    10.96.43.104   <none>        80:31806/TCP   22h
kubernetes    ClusterIP   10.96.0.1      <none>        443/TCP        24h
myweb-svc     ClusterIP   10.96.187.21   <none>        8081/TCP       79s


[root@laptop tmp]# kubectl describe svc myweb-svc
Name:              myweb-svc
Namespace:         default
Labels:            <none>
Annotations:       <none>
Selector:          app=myweb
Type:              ClusterIP
IP:                10.96.187.21
Port:              <unset>  8081/TCP
TargetPort:        8080/TCP
Endpoints:         172.17.0.9:8080
Session Affinity:  None
Events:            <none>
```



## 多端口服务

有时一个容器应用也可能需要提供多个端口的服务

那么在service定义中也可以设置将多个端口对应多个应用服务



## 外部服务service

* 内容待补充