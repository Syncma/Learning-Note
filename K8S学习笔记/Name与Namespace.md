# Name和Namespace

<!-- TOC -->

- [Name和Namespace](#name%e5%92%8cnamespace)
  - [Name](#name)
  - [Namespace](#namespace)

<!-- /TOC -->

## Name

* 由于k8s内部，使用“资源”来定义每一种逻辑概念（功能），故每种资源都应该有自己的名称
* 资源有api版本(apiVersion)，类别(kind)，元数据(metadata)，定义清单(spec)，状态(status)等配置信息
* 名称通常定义在资源的元数据信息里面



## Namespace

* 随着项目增多，人员增加，集群规模的扩大，需要一种能够隔离k8s内每种资源的方法，这就是名称空间

* 名称空间可以理解为k8s内部的虚拟集群组

* 不同名称空间内的资源，名称可以相同；相同名称空间内的同种资源，名称不能相同

* 合理的使用k8s名称空间，使集群管理员能够更好的对交付到k8s里面的服务进行分类管理和浏览

* k8s里面默认的存在的名称空间有: default, kube-system, kube-public

* 查询k8s里面特定的资源要带上相应的名称空间