# YARN工作机制

<!-- TOC -->

- [YARN工作机制](#yarn%e5%b7%a5%e4%bd%9c%e6%9c%ba%e5%88%b6)
  - [流程图](#%e6%b5%81%e7%a8%8b%e5%9b%be)
  - [简单讲解](#%e7%ae%80%e5%8d%95%e8%ae%b2%e8%a7%a3)
  - [具体讲解](#%e5%85%b7%e4%bd%93%e8%ae%b2%e8%a7%a3)

<!-- /TOC -->

## 流程图
![enter image description here](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/yarn-work.png)



## 简单讲解

```
1. Client 提交作业到 YARN 上
2. Resource Manager 选择一个 Node Manager，启动一个 Container 并运行 Application Master 实例
3. Application Master 根据实际需要向 Resource Manager 请求更多的 Container 资源（如果作业很小, 应用管理器会选择在其自己的 JVM 中运行任务）
4. Application Master 通过获取到的 Container 资源执行分布式计算

```


## 具体讲解

- 要看源码 待补充