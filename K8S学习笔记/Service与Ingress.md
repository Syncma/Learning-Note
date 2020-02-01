# Service与Ingress

<!-- TOC -->

- [Service与Ingress](#service%e4%b8%8eingress)
  - [Service](#service)
  - [Ingress](#ingress)

<!-- /TOC -->

## Service

* 在k8s的世界里，虽然每个pod都会被分配一个单独的IP地址，但这个IP地址会随着pod销毁而消失
* Service（服务）就是用来解决这个问题的核心概念
* 一个Service可以看作一组提供相同服务的Pod的对外访问接口
* Service作用于哪些pod是通过标签选择器来定义的



## Ingress

* Ingress是k8s集群里工作在OSI网络模型下，第7层的应用对外暴露的接口
* Service只能进行L4流量调度，表现方式是ip+port
* Ingress可以调度不同业务域，不同URL访问路径的业务流量