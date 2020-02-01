# Volume篇
<!-- TOC -->

- [Volume篇](#volume%e7%af%87)
  - [介绍](#%e4%bb%8b%e7%bb%8d)
  - [使用](#%e4%bd%bf%e7%94%a8)

<!-- /TOC -->
## 介绍

Volume是Pod中能够被多个容器访问的目录

k8s的Volume定义在Pod上，它被一个Pod中的多个容器挂载到具体的文件目录下

Volume与Pod的生命周期相同，但与容器的生命周期不相关

当容器终止或重启，Volume中的数据也不会丢失



## 使用

要使用Volume，pod需要指定volume的类型和内容(**spec.volumes**字段)，

和映射到容器的位置(**spec.containers.volumeMounts**字段)

k8s支持多种类型的Volume



1.EmptyDir

EmptyDir类型的Volume创建于pod被调度到某个宿主机上的时候，

而同一个pod内的容器能读写EmptyDir中的同一个文件

一旦这个pod离开了这个宿主机，EmptyDir中的数据就会被永久删除

目前这个类型主要用来临时空间，比如web服务器写日志或者临时目录



例子：

* 待补充



2.还有其他类型的Volume...