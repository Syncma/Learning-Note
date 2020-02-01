# 核心组件

<!-- TOC -->

- [核心组件](#核心组件)
    - [核心组件](#核心组件-1)
    - [CLI客户端](#cli客户端)
    - [核心附件](#核心附件)
    - [k8s三条网络](#k8s三条网络)

<!-- /TOC -->

## 核心组件

* 配置存储中心->etcd服务

* 主控(master)节点

  * kube-apiserver服务
  * kube-controller-manager服务
  * kube-scheduler服务

* 运算(node)节点

  	* kube-kubelet服务
  	* kube-proxy服务

  

  

## CLI客户端

  * kubectl

  

## 核心附件

* CNI网络插件->finnel/calico
* 服务发现用插件->coredns
* 服务暴露用插件->traefik
* GUI管理插件->Dashboard





## k8s三条网络

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/k8s-network.png)