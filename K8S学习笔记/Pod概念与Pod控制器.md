# Pod与Pod控制器

<!-- TOC -->

- [Pod与Pod控制器](#pod%e4%b8%8epod%e6%8e%a7%e5%88%b6%e5%99%a8)
  - [Pod](#pod)
  - [Pod控制器](#pod%e6%8e%a7%e5%88%b6%e5%99%a8)

<!-- /TOC -->

## Pod

* Pod是k8s里能够被运行的最小的逻辑单元
* 1个pod里面可以运行多个容器，它们共享UTS+NET+IPC名称空间
* 可以把Pod理解为豌豆荚，而同一个Pod内的每个容器是一颗颗豌豆
* 一个Pod里运行多个容器，又叫边车(SideCar)模式



## Pod控制器

* Pod 控制器是Pod启动的一种模板，用来保证在K8S里启动的Pod始终按照预期运行

* K8S内提供许多Pod控制器，常用的有：

  * **Deployment**

  * **DaemonSet**

  * ReplicaSet

  * Job

  * CronJob