# Dockerfile体系结构
<!-- TOC -->

- [Dockerfile体系结构](#dockerfile%e4%bd%93%e7%b3%bb%e7%bb%93%e6%9e%84)
  - [保留字指令](#%e4%bf%9d%e7%95%99%e5%ad%97%e6%8c%87%e4%bb%a4)
  - [dockerfile运行机制](#dockerfile%e8%bf%90%e8%a1%8c%e6%9c%ba%e5%88%b6)
  - [例子](#%e4%be%8b%e5%ad%90)

<!-- /TOC -->

## 保留字指令

```
FROM：基础镜像，当前新镜像是基于哪个镜像的

MAINTANER：镜像维护者的姓名和邮箱地址

RUN：容器构建需要运行的命令

EXPOSE：当前容器对外暴露的端口

EXPOSE 指令是声明运行时容器提供服务端口，这只是一个声明，在运行时并不会因为这个声明应用就会开启这个端口的服务。

在 Dockerfile 中写入这样的声明有两个好处，一个是帮助镜像使用者理解这个镜像服务的守护端口，以方便配置映射；

另一个用处则是在运行时使用随机端口映射时，也就是 docker run -P 时，会自动随机映射 EXPOSE 的端口。

要将 EXPOSE 和在运行时使用 -p <宿主端口>:<容器端口> 区分开来。

-p，是映射宿主端口和容器端口，换句话说，就是将容器的对应端口服务公开给外界访问，
而 EXPOSE 仅仅是声明容器打算使用什么端口而已，并不会自动在宿主进行端口映射。


WORKDIR：指定在创建容器后，终端默认登录的进来工作目录，一个落脚点

ENV：用来构建镜像过程中设置环境变量

ADD：将宿主机目录下的文件拷贝进镜像且ADD命令会自动处理URL和解压tar压缩包

COPY：类似ADD，拷贝文件和目录到镜像中，将从构建上下文目录中<源路径>的文件/目录复制到新的一层镜像内的<目标路径>位置
COPY src dest
COPY ["src", "dest"]

VOLUME：容器数据卷，用于数据保存和持久化工作

CMD：指定一个容器启动时要运行的命令，
Dockerfile中可以有多个CMD指令，但只有最后一个生效，CMD会被docker run之后的参数替换

CMD ["curl", "-s", "http://ip.cn"]
CMD -i 

ENTRYPOINT：指定一个容器启动时要运行的命令，和CMD一样都是在指定容器启动程序以及参数 
指定会追加 不会被覆盖

ENTRYPOINT  ["curl", "-s", "http://ip.cn"]


ONBUILD： 当构建一个被继承的Dockerfile时运行命令，父镜像在被子继承后父镜像的onbuild被触发
```


## dockerfile运行机制

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/dockerfile2.png)




## 例子

```
base镜像(scratch)： Docker Hub中99%的镜像都是通过在base镜像中安装和配置需要的软件构建起来的

自定义镜像(mycentos)
```

```
[root@laptop mydocker]# cat dockerfile

FROM centos
MAINTAINER majian<majian@qq.com>
ENV mypath /tmp
WORKDIR $mypath
RUN yum -y install vim
RUN yum -y install net-tools
EXPOSE 80
CMD /bin/bash
```

```
[root@laptop mydocker]# docker build -f dockerfile -t mycentos:1.3 .

Sending build context to Docker daemon 2.048kB
Step 1/8 : FROM centos
---> 0f3e07c0138f
Step 2/8 : MAINTAINER majian<majian@qq.com>
---> Running in 6d4684ead2fa
Removing intermediate container 6d4684ead2fa
---> e899711fd7b2
Step 3/8 : ENV mypath /tmp
---> Running in 25e2618ca1c0
Removing intermediate container 25e2618ca1c0
---> 498e81a2ddc6
Step 4/8 : WORKDIR $mypath
---> Running in 51c12df8bf5d
Removing intermediate container 51c12df8bf5d
---> 051be51956d5
Step 5/8 : RUN yum -y install vim
---> Running in 40f00feca24d

```

```
[root@laptop mydocker]# docker images

REPOSITORY TAG IMAGE ID CREATED SIZE
mycentos 1.3 b9d313b868fc 32 seconds ago 316MB
centos latest 0f3e07c0138f 6 weeks ago 220MB


[root@laptop mydocker]# docker images mycentos

REPOSITORY TAG IMAGE ID CREATED SIZE
mycentos 1.3 b9d313b868fc About a minute ago 316MB


[root@laptop mydocker]# docker run -it mycentos:1.3

[root@9eb191b85d6e tmp]# pwd
/tmp
[root@9eb191b85d6e tmp]# which vim
/usr/bin/vim
[root@9eb191b85d6e tmp]# which ifconfig
/usr/sbin/ifconfig

```
列出镜像的变更历史：

```
[root@laptop mydocker]# docker images

REPOSITORY TAG IMAGE ID CREATED SIZE
mycentos 1.3 b9d313b868fc 3 minutes ago 316MB
abc/centos latest 0760a8000e5d 5 hours ago 220MB
abc/mytomcat 1.2 f73b412690d6 6 hours ago 507MB
tomcat latest 882487b8be1d 3 weeks ago 507MB
centos latest 0f3e07c0138f 6 weeks ago 220MB

[root@laptop mydocker]# docker history b9d313b868fc

IMAGE CREATED CREATED BY SIZE COMMENT
b9d313b868fc 3 minutes ago /bin/sh -c #(nop) CMD ["/bin/sh" "-c" "/bin… 0B
2c49c2457ddc 3 minutes ago /bin/sh -c #(nop) EXPOSE 80 0B
4c1421515876 3 minutes ago /bin/sh -c yum -y install net-tools 29.3MB
1e5944f9ebf8 3 minutes ago /bin/sh -c yum -y install vim 66.9MB
051be51956d5 4 minutes ago /bin/sh -c #(nop) WORKDIR /tmp 0B
498e81a2ddc6 4 minutes ago /bin/sh -c #(nop) ENV mypath=/tmp 0B
e899711fd7b2 4 minutes ago /bin/sh -c #(nop) MAINTAINER majian<majian@… 0B
0f3e07c0138f 6 weeks ago /bin/sh -c #(nop) CMD ["/bin/bash"] 0B
<missing> 6 weeks ago /bin/sh -c #(nop) LABEL org.label-schema.sc… 0B
<missing> 6 weeks ago /bin/sh -c #(nop) ADD file:d6fdacc1972df524a… 220MB
```


再看一个例子:
```
[root@laptop mydocker]# cat dockerfile

FROM centos

ONBUILD RUN echo "father image onbuild"

[root@laptop mydocker]# docker build -f dockerfile -t myip_father .

Sending build context to Docker daemon 2.048kB
Step 1/2 : FROM centos
---> 0f3e07c0138f
Step 2/2 : ONBUILD RUN echo "father image onbuild"
---> Running in 9f077f1c3c30
Removing intermediate container 9f077f1c3c30
---> 145433c7a943
Successfully built 145433c7a943
Successfully tagged myip_father:latest

```

```
[root@laptop mydocker]# cat dockerfile2

FROM myip_father

[root@laptop mydocker]# docker build -f dockerfile2 -t myip:son .
Sending build context to Docker daemon 3.072kB
Step 1/1 : FROM myip_father
# Executing 1 build trigger
---> Running in d636ed5d3e0c
father image onbuild
Removing intermediate container d636ed5d3e0c
---> ec2d399db156
Successfully built ec2d399db156
Successfully tagged myip:son
```