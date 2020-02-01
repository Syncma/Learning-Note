# Pod详解

<!-- TOC -->

- [Pod详解](#pod%e8%af%a6%e8%a7%a3)
  - [Pod介绍](#pod%e4%bb%8b%e7%bb%8d)
  - [Pod 定义](#pod-%e5%ae%9a%e4%b9%89)
  - [Pod使用](#pod%e4%bd%bf%e7%94%a8)
  - [Pod分类](#pod%e5%88%86%e7%b1%bb)
    - [普通Pod](#%e6%99%ae%e9%80%9apod)
    - [静态Pod](#%e9%9d%99%e6%80%81pod)
  - [Pod生命周期和重启策略](#pod%e7%94%9f%e5%91%bd%e5%91%a8%e6%9c%9f%e5%92%8c%e9%87%8d%e5%90%af%e7%ad%96%e7%95%a5)
    - [Pod状态](#pod%e7%8a%b6%e6%80%81)
    - [Pod重启策略](#pod%e9%87%8d%e5%90%af%e7%ad%96%e7%95%a5)
  - [Pod资源配置](#pod%e8%b5%84%e6%ba%90%e9%85%8d%e7%bd%ae)

<!-- /TOC -->

## Pod介绍

**Pod 是k8s的重要概念，要掌握**

每个Pod都有一个特殊的被称为“根容器”的Pause容器

Pause容器对应的镜像属于k8s平台的一部分，除了Pause容器还包含一个或多个紧密相关的业务容器



Pod图示：

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/pod-1.png)





* Pod VS 应用

  每个Pod都是应用的一个实例，有专用的IP

* Pod VS 容器

  一个Pod可以有多个容器，彼此间共享网络和存储资源

  每个Pod中有一个Pause容器保存所有的容器状态，通过管理Pause容器，达到管理Pod中所有容器的效果

* Pod VS 节点

  同一个Pod中的容器总会被调用到相同的Node节点，不同节点Pod的通信基于虚拟二层网络技术实现

* Pod VS Pod

  普通的Pod 和静态Pod



## Pod 定义

通过yaml文件格式定义pod

[yaml格式校验工具](https://www.bejson.com/validators/yaml/)



## Pod使用

在k8s中对运行的容器要求为：

> **容器的主程序需要一直在前台运行，而不是后台运行**
>
> 所以应用需要改造成前台运行的方式



如果我们创建的docker 镜像的启动命令是后台执行程序，则kubelet创建包含这个容器的pod

之后运行该命令，即认为Pod已经结束，将立刻销毁该Pod.

如果该Pod定义了RC，则创建、销毁会陷入一个无限循环的过程中



Pod可以由1个或多个容器组合而成



* 一个容器组成的pod

  demo1.yaml

  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: mytomcat
    labels: 
      name: mytomcat
  spec:
    containers:
    - name: mytomcat
      image: tomcat
      ports:
      - containerPort: 8000
  ```

  

  创建pod:

  ```
  [root@laptop tmp]# kubectl create -f demo1.yaml 
  pod/mytomcat created
  ```

  

  查看pod:

  ```
  [root@laptop tmp]# kubectl get pods //或者kubectl get po
  NAME                          READY   STATUS              RESTARTS   AGE
  mytomcat                      0/1     ContainerCreating   0          2m30s
  
  [root@laptop tmp]# kubectl describe po mytomcat //查看特定的name
  ```

  

  删除pod:

  ```
  [root@laptop tmp]# kubectl delete -f  demo1.yaml 
  
  [root@laptop tmp]# kubectl delete pod --all/[pod_name]
  ```


* 多个容器组成的pod

  demo2.yaml

  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: myweb
    labels: 
      name: tomcat-redis
  spec:
    containers:
    - name: tomcat
      image: tomcat
      ports:
      - containerPort: 8000
    - name: redis
      image: redis
      ports:
      - containerPort: 6379
  ```


## Pod分类

### 普通Pod 

普通pod 一旦被创建，就会被放入到etcd中存储，随后会被k8s master调度到某个具体的Node上并进行绑定，随后该pod对应的Node上的kubelet进程实例化成一组相关的docker容器并启动起来

在默认情况下，当pod里某个容器停止时，k8s会自动检测到这个问题并且重新启动这个pod里所有容器。

如果该pod所在的Node当机，则会将这个Node上所有的pod重新调度到其他节点上



### 静态Pod

静态pod是由kubelet进行管理的仅存在于特定Node上的pod，它们不能通过API Server进行管理。

无法与RC(replicationController)、Deployment和DaemonSet进行关联，并且kubelet也无法对它们进行健康检测



## Pod生命周期和重启策略



### Pod状态

| 状态值    | 说明                                                                                 |
| --------- | ------------------------------------------------------------------------------------ |
| Pending   | API Server已经创建了该pod，但Pod中的一个或多个容器的镜像还没有创建，包括镜像下载过程 |
| Running   | Pod 内所有容器已创建，且至少一个容器处于运行状态，正在启动状态或正在重启状态         |
| Completed | Pod内所有容器均成功执行退出，且不会再重启                                            |
| Failed    | Pod内所有容器均已退出，但至少一个容器退出失败                                        |
| Unknown   | 由于某种原因无法获取Pod状态，例如网络通信不通                                        |



### Pod重启策略

重启策略包括Always, OnFailure和Never , 默认是Always

| 重启策略  | 说明                                                   |
| --------- | ------------------------------------------------------ |
| Always    | 当容器失效时，由kubelet自动重启该容器                  |
| OnFailure | 当容器终止运行且退出码不为0时，由kubelet自动重启该容器 |
| Never     | 不论容器运行状态如何，kubelet都不会重启该容器          |



## Pod资源配置

每个Pod都可以对其能使用的服务器上的计算资源设置限额，

**`k8s中可以设置限额的计算资源有CPU和Memory两种`**

其中CPU的资源单位是CPU数量，是一个绝对值而非相对值

Memory配置也是一个绝对值，它的单位是字节数

k8s中，一个计算资源进行配置需要设定两个参数：

* requests 表示该资源最小申请数量，系统必须满足要求
* limits表示该资源最大允许使用的量，不能突破，该容器使用超过这个量的资源时，会被k8s kill并重启



看一个例子：

```yaml
spec:
 container:
 - name: db
   image: mysql
   resources:
     requests:
       memory: "64Mi"
       cpu: "250m"
     limits:
       memory: "128Mi"
       cpu: "500m"
```



上面的代码表示mysql容器申请至少0.25个cpu和64MiB内存

在运行过程中容器能使用的资源配额是0.5个cpu和128MiB内存