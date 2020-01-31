# Label和Label选择器

<!-- TOC -->

- [Label和Label选择器](#label%e5%92%8clabel%e9%80%89%e6%8b%a9%e5%99%a8)
  - [Label](#label)
  - [Label选择器](#label%e9%80%89%e6%8b%a9%e5%99%a8)

<!-- /TOC -->

## Label

* 标签是k8s特色的管理方式，便于分类管理资源对象
* 一个标签可以对应多个资源，一个资源也可以有多个标签，它们是多对多的关系
* 一个资源拥有多个标签，可以实现不同维度的管理
* 标签的组成: key=value
* 与标签类似的，还有一种注解(annotations)



## Label选择器

* 给资源打上标签后，可以使用标签选择器过滤指定的标签
* 标签选择器目前有两个，基于等值关系（等于、不等于）和基于集合关系（属于、不属于、存在）
* 许多资源支持内嵌标签选择器字段
  * matchLables
  * matchExpressions