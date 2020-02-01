# Label篇
<!-- TOC -->

- [Label篇](#label%e7%af%87)
  - [介绍](#%e4%bb%8b%e7%bb%8d)
  - [作用](#%e4%bd%9c%e7%94%a8)

<!-- /TOC -->
## 介绍



* 标签是k8s特色的管理方式，便于分类管理资源对象

* 一个标签可以对应多个资源，一个资源也可以有多个标签，它们是多对多的关系

* 一个资源拥有多个标签，可以实现不同维度的管理

* 标签的组成: key=value



Label常见用法是**metadata.labels**字段，用来给对象添加label，通过**spec.selector**来引用对象



## 作用

目的是对这些资源进行分组管理，分组管理的核心就是label Selector

注意：

> Label和Label Selector 都不能单独定义，必须附加在一些资源对象的定义文件中
>
> 一般附加在RC和Service的资源定义文件中