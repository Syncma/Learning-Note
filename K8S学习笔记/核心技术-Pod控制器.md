# Pod控制器

<!-- TOC -->

- [Pod控制器](#pod%e6%8e%a7%e5%88%b6%e5%99%a8)
  - [RC](#rc)
  - [RS](#rs)
  - [Deployment](#deployment)
  - [HPA](#hpa)

<!-- /TOC -->

## RC

全称：**ReplicationController**

当我们定义了一个RC 并提交到k8s集群中，Master节点上的Control Manager组件就得到通知

定期检测系统中存活的Pod，确保Pod实例的数量等于RC的设置值

如果有过多或过少的Pod运行，系统会停掉或创建一些Pod

也可以修改RC副本数量，实现Pod动态缩放功能



例子：

demo.yaml

```yaml
apiVersion: v1
kind: ReplicationController
metadata:
  name: nginx
spec:
 replicas: 3
 selector:
  app: nginx
 template:
  metadata:
   labels:
    app: nginx
  spec:
   containers:
   - name: nginx
     image: nginx
     ports:
     - containerPort: 80
```



创建Pod:

```
[root@laptop tmp]# kubectl create -f demo.yaml 
replicationcontroller/nginx created
```



查看Pod情况：

```
[root@laptop tmp]# kubectl get pods
NAME                          READY   STATUS    RESTARTS   AGE
hello-world-f7dbcbd8f-ghq2r   1/1     Running   2          19h
nginx-47nk6                   1/1     Running   0          34s
nginx-8wxrn                   1/1     Running   0          34s
nginx-tdg6g                   1/1     Running   0          34s
```



可以看到创建了三个nginx, 如果我们想动态扩建到5个，可以使用下面的命令进行操作：

```
[root@laptop tmp]# kubectl scale rc nginx --replicas=5
replicationcontroller/nginx scaled
```



再次查看Pod情况：

```
[root@laptop tmp]# kubectl get pods
NAME                          READY   STATUS    RESTARTS   AGE
hello-world-f7dbcbd8f-ghq2r   1/1     Running   2          20h
nginx-47nk6                   1/1     Running   0          2m4s
nginx-5sbhm                   1/1     Running   0          13s
nginx-8wxrn                   1/1     Running   0          2m4s
nginx-q629v                   1/1     Running   0          13s
nginx-tdg6g                   1/1     Running   0          2m4s
```



由于Replication Controller与k8s代码中的模块Replication Controller同名，所以在新版本的k8s中，

升级为新的概念 Replica Set来取代Replication Controller

值得注意：

> 最好不要越过RC直接创建Pod
>
> 因为Replication Controller 会通过RC管理Pod
>
> 所以即使应用只有一个Pod副本也建议使用RC来定义Pod



## RS

全称： Replica Set

它与RC区别：

> RS（Replica Set）支持基于集合的Label Selector
>
> 而RC 只支持基于等式的Label Selector
>
> **`可以详细看看Label章节`**



虽然RS可以独立使用（一般很少单独使用），一般建议使用deployment来自动管理RS

与RC没有本质的区别，基本用法是相同的



## Deployment

目的是为了更好的解决Pod编排问题

内部使用RS来实现，deployment定义和RS定义类似，除了API声明和Kind类型有所区别：

例子：

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: frontend
spec:
 replicas: 1
 selector:
  matchLabels:
   tier: frontend
  matchExpressions:
   - {key: tier, operation:In, values:[fronted]}
 template:
  metadata:
   labels:
    app: app-demo
    tier: frontend
  spec:
   containers:
   - name: tomcat-demo
     image: tomcat
     ports:
     - containerPort: 8080
```





## HPA

全称： Horizontal Pod Autoscal -Pod横向扩容

通过分析RC控制的所有目标Pod的负载变化情况，来确定是否需要针对性的调整目标Pod的副本数

这就是HPA的实现原理

k8s对Pod扩容与缩容提供了手动和自动的模式

* 手动

  ```
  kubectl scale deployment xxx --replicas 1
  ```

  

* 自动

  需要用户根据性能指标来指定Pod副本范围，系统会在这个范围内根据性能指标进行调整

  可以根据Pod的CPU利用率来扩容，也可以根据内存或者用户自定义的指标