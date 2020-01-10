# K8S
<!-- TOC -->

- [K8S](#k8s)
  - [K8S 介绍](#k8s-%e4%bb%8b%e7%bb%8d)
  - [K8S和Mesos](#k8s%e5%92%8cmesos)
    - [主要区别](#%e4%b8%bb%e8%a6%81%e5%8c%ba%e5%88%ab)
    - [业务需求](#%e4%b8%9a%e5%8a%a1%e9%9c%80%e6%b1%82)
  - [docker三剑客区别](#docker%e4%b8%89%e5%89%91%e5%ae%a2%e5%8c%ba%e5%88%ab)

<!-- /TOC -->


## K8S 介绍

会单独介绍k8s

## K8S和Mesos

Apache Mesos 和 Kubernetes 都是优秀的开源框架，都支持大规模集群管理
（当然开源 Kubernetes 目前还局限于数千，万级节点还需要定制化，而 Apache Mesos 可以轻量级地调度万级节点）

### 主要区别
* 千节点集群，少定制：使用开源 Kubernetes （细粒度设计，契合微服务思想）
* 万节点集群，多定制：使用 Mesos + Marathon （双层调度好犀利）
* 万节点集群，IT 能力强：深度定制 Kubernetes （如网易云）
* 万节点集群，IT 能力强：深入掌握使用 DC/OS （DC/OS 在最基础的 Marathon 和 Mesos 之上添加了很多的组件）
* 大数据集群：Spark on Mesos （建议只基于容器部署计算部分，数据部分另行部署）


### 业务需求

Mesos和Kubernetes使用哪个好，关键还是要**看业务需求**，当然是适合业务需求的最好。

* 两者参数都都挺多的，针对于容器管理，两者也都挺灵活，一般都可以满足，不过学习成本也会相对提高。
* 两者对比，有点类似搭积木，Mesos是零散的积木，你需要自己组装实现自己的业务模型，Kubernetes就是组装好的积木，你直接拿来用就好了。
* Mesos自身定义为一个分布式内核，也就是最核心的资源管理它帮你实现了，如果自己实现业务调度，那么就需要公司有一定的开发能力，当然了，一般的需求使用Mesos的框架Marathon和Aurora等也能满足了。
* 两者发展的侧重点不同，Mesos更侧重底层资源的管理，Kubernetes侧重业务层的调度，容器服务编排，服务发现等。还有现在Kubernetes也可以运行在Mesos上，这样你也可以选择两者结合。
* 现在在国内，Kubernetes感觉更火些，个人觉得这可能跟容器的爆发有关，并且有Google公司的光环。



## docker三剑客区别

1.docker-machine: 是解决docker运行环境问题

docker技术是基于Linux内核的cgroup技术实现的，那么问题来了，如果在非Linux平台上使用docker技术需要依赖安装Linux系统的虚拟机。
docker-machine就是docker公司官方提出的，用于在各种平台上快速创建具有docker服务的虚拟机的技术。


2.docker-compose：是解决本地docker容器编排问题


一般是通过yaml配置文件来使用它，这个yaml文件里能记录多个容器启动的配置信息（镜像、启动命令、端口映射等），最后只需要执行docker-compose对应的命令就会像执行脚本一样地批量创建和销毁容器。


3.docker-swarm：是解决多主机多个容器调度部署问题
swarm是基于docker平台实现的集群技术，他可以通过几条简单的指令快速的创建一个docker集群，接着在集群的共享网络上部署应用，最终实现分布式的服务。

swarm技术相当不成熟，很多配置功能都无法实现，只能说是个半成品，

**目前更多的是使用Kubernetes来管理集群和调度容器**。

