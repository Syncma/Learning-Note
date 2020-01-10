# Dockerfile


<!-- TOC -->

- [Dockerfile](#dockerfile)
  - [是什么](#%e6%98%af%e4%bb%80%e4%b9%88)
  - [Dockerfile构建过程解析](#dockerfile%e6%9e%84%e5%bb%ba%e8%bf%87%e7%a8%8b%e8%a7%a3%e6%9e%90)
    - [基础知识](#%e5%9f%ba%e7%a1%80%e7%9f%a5%e8%af%86)
    - [Docker执行Dockerfile的大致流程](#docker%e6%89%a7%e8%a1%8cdockerfile%e7%9a%84%e5%a4%a7%e8%87%b4%e6%b5%81%e7%a8%8b)

<!-- /TOC -->
## 是什么

**Dockerfile是用来构建Docker镜像的构建文件，是由一系列命令和参数构成的脚本**

构建三步骤：
```
编写dockerfile文件->docker build  -> docker run
```

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/dockerfile1.png)




## Dockerfile构建过程解析

### 基础知识

每条保留字指令都必须为大写字母且后面要跟随至少一个参数
指令按照从上到下，顺序执行
`#表示注释`
每条指令都会创建一个新的镜像层，并对镜像进行提交

### Docker执行Dockerfile的大致流程

1）docker 从基础镜像运行一个容器

2）执行一条指令并对容器作为修改

3）执行类型docker commit的操作提交一个新的镜像层

4）docker再基于刚提交的镜像运行一个新容器

5）执行dockerfile中的下一条指令直到所有指令都执行完成



