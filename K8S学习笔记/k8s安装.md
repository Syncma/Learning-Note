# k8s安装

<!-- TOC -->

- [k8s安装](#k8s%e5%ae%89%e8%a3%85)
  - [minikube方式](#minikube%e6%96%b9%e5%bc%8f)
    - [介绍](#%e4%bb%8b%e7%bb%8d)
    - [kubectl安装](#kubectl%e5%ae%89%e8%a3%85)
    - [MiniKube安装](#minikube%e5%ae%89%e8%a3%85)
      - [安装](#%e5%ae%89%e8%a3%85)
      - [使用](#%e4%bd%bf%e7%94%a8)
      - [部署应用](#%e9%83%a8%e7%bd%b2%e5%ba%94%e7%94%a8)
  - [kubeadm方式](#kubeadm%e6%96%b9%e5%bc%8f)
  - [区别](#%e5%8c%ba%e5%88%ab)

<!-- /TOC -->



## minikube方式



### 介绍

Minikube 是一种可以让您在本地轻松运行 Kubernetes 的工具。

Minikube可以实现一种轻量级的Kubernetes集群，通过在本地计算机上创建虚拟机并部署只包含单个节点的简单集群



官网文档：

[英文文档](https://kubernetes.io/docs/tutorials/hello-minikube/)

[minikube中文文档](https://kubernetes.io/zh/docs/setup/learning-environment/minikube/)

[minikube英文文档](https://kubernetes.io/docs/setup/learning-environment/minikube/)





### kubectl安装

MiniKube 的安装需要先安装 kubectl （k8s客户端）及相关驱动

这里使用二进制方式进行安装

本地环境: Fedora x29  x64



1.查看最新的版本号

```
[jian@laptop tmp]$ curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt
v1.17.2

# 根据版本号进行下载相应的客户端程序
[jian@laptop tmp]$ curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.17.2/bin/linux/amd64/kubectl

```



也可以直接下载最新版本：

```
[jian@laptop tmp]$ curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
```



2.赋予可执行权限

```
[jian@laptop tmp]$ chmod +x ./kubectl
```



3.放入系统环境路径

```
[jian@laptop tmp]$ sudo mv kubectl /usr/local/bin/kubectl
```



4.测试版本信息

```
[root@laptop bin]# kubectl version --client
```



### MiniKube安装

MiniKube 是使用 Go 语言开发的，所以安装其实很方便，这里也使用二进制方式进行安装

[下载地址](https://github.com/kubernetes/minikube/releases)

这里我们下载1.6.2   minikube-linux-amd64 版本



#### 安装

```
[jian@laptop tmp]$ curl -Lo minikube https://github.com/kubernetes/minikube/releases/download/v1.6.2/minikube-linux-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/
```



#### 使用

1.开启docker 服务

```
[root@laptop ~]# systemctl start docker

[root@laptop system]# systemctl status docker.service
● docker.service - Docker Application Container Engine
   Loaded: loaded (/usr/lib/systemd/system/docker.service; disabled; vendor preset: disabled)
   Active: active (running) since Thu 2020-01-30 11:49:20 CST; 7s ago
     Docs: https://docs.docker.com
     ....
```



2.默认启动使用的是 VirtualBox 驱动，使用 `--vm-driver` 参数可以指定其它驱动

```
参数说明：
--image-registry 使用阿里云镜像进行加速
--vm-driver=none 不使用任何驱动

[root@laptop ~]# minikube start --vm-driver=none --image-repository registry.cn-hangzhou.aliyuncs.com/google_containers
....
🏄  Done! kubectl is now configured to use "minikube"
```





3.检测状态

```
[root@laptop ~]# minikube status
host: Running
kubelet: Running
apiserver: Running
kubeconfig: Configured
```





4.启动k8s dashboard 

```
[root@laptop ~]# minikube dashboard
🤔  Verifying dashboard health ...
🚀  Launching proxy ...
🤔  Verifying proxy health ...
http://127.0.0.1:35173/api/v1/namespaces/kubernetes-dashboard/services/http:kubernetes-dashboard:/proxy/

点击上面的地址就可以打开dashboard
```



5.查看集群情况

```
[root@laptop ~]# kubectl cluster-info
```



6.查看节点情况

```
[root@laptop ~]# kubectl get nodes
```



#### 部署应用

1.创建一个deployment

```
[root@laptop ~]# kubectl run hello-world --image=nginx:1.7.9 --port=80
```



2.查看deployment

```
[root@laptop ~]# kubectl get deployments
NAME          READY   UP-TO-DATE   AVAILABLE   AGE
hello-world   0/1     1            0           67s
```



3.查看pod

```
[root@laptop ~]# kubectl get pods
NAME                          READY   STATUS    RESTARTS   AGE
hello-world-f7dbcbd8f-ghq2r   1/1     Running   0          93s
```



4.查看日志命令

```
[root@laptop ~]# minikube logs
```



5.再次查看deployment

```
[root@laptop ~]# kubectl get deployments
NAME          READY   UP-TO-DATE   AVAILABLE   AGE
hello-world   1/1     1            1           3m38s
```



6.创建服务

默认情况下，Pod 只能通过 Kubernetes 集群中的内部 IP 地址访问。

要使得 容器可以从 Kubernetes 虚拟网络的外部访问，您必须将 Pod 暴露为 Kubernetes [*Service*](https://k8smeetup.github.io/docs/concepts/services-networking/service/)。

```
[root@laptop ~]# kubectl expose deployment hello-world --type=NodePort
# z注意这里--type=NodePort
```



7.查看服务

```
[root@laptop ~]# kubectl get services
NAME          TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
hello-world   NodePort    10.96.43.104   <none>        80:31806/TCP   3m55s
kubernetes    ClusterIP   10.96.0.1      <none>        443/TCP        108m
```



8.访问应用

```
# --url：将返回访问的URL 
[root@laptop ~]# minikube service hello-world --url
http://192.168.1.102:31806

可以通过浏览器直接访问
```



## kubeadm方式

kubeadm是Kubernetes1.6开始官方推出的快速部署Kubernetes集群工具

其思路是将Kubernetes相关服务容器化(Kubernetes静态Pod)以简化部署



* 安装过程待补充(需要多台机器才能弄)



## 区别

> minikube是单机版
> kubeadm 是运行在docker里面的k8s集群
>
> minikube 基本上你可以认为是一个实验室工具，只能单机部署，里面整合了 k8s 最主要的组件，无法真正搭建集群，且由于程序做死无法安装各种扩展插件（比如网络插件、dns 插件、ingress 插件等等），主要作用是给你了解 k8s 用的。
>
> 而 kudeadm 搭建出来是一个真正的 k8s 集群，可用于生产环境（HA 需要自己做），和二进制搭建出来的集群几乎没有区别。