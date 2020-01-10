# Docker容器数据卷

<!-- TOC -->

- [Docker容器数据卷](#docker容器数据卷)
    - [是什么](#是什么)
    - [能干嘛](#能干嘛)
    - [数据卷](#数据卷)
        - [容器内添加](#容器内添加)
        - [DockerFile添加](#dockerfile添加)
    - [数据卷容器](#数据卷容器)
        - [是什么](#是什么-1)
        - [容器之间传递共享](#容器之间传递共享)

<!-- /TOC -->

## 是什么

先看看docker的理念：

* 将运用与运行的环境打包形成容器运行，运行可以伴随着容器，但是我们对数据的要求希望是持久化的
*  容器之间有可能共享数据


Docker容器产生的数据，如果不通过docker commit生成新的镜像，使得数据作为镜像的一部分保存下来，那么当容器删除后，数据自然也就没有了

**为了能保存数据在docker中我们使用卷**

**`一句话类似redis里面的rdb 和aof文件`**


## 能干嘛

**`卷就是目录或文件`**，存在与一个或多个容器中，由docker挂载到容器，但不属于联合文件系统，因此能够绕过UnionFS，提供一些用于存储或共享数据的特征

卷的设计目标就是**数据的持久化**，完全独立于容器的生存周期，因此Docker不会在容器删除时删除其挂载的数据卷

特点：
* 数据卷可在容器之间共享或重用数据
* 卷的更改可以直接生效
* 数据卷中的更改不会包含在镜像的更新中
* 数据卷的生命周期一直持续到没有容器使用它为止


## 数据卷

### 容器内添加

1.直接命令添加

```
docker run -it -v /宿主机绝对路径目录:/容器内目录  镜像名
```

例子：
```
[root@laptop docker]# docker images
REPOSITORY TAG IMAGE ID CREATED SIZE
centos latest 0f3e07c0138f 6 weeks ago 220MB

[root@laptop docker]#  mkdir /myDataVolume
[root@laptop docker]# docker run -it -v /myDataVolume:/dataVolumeContainer centos
[root@f71c79781ee3 /]# ls
bin dev home lib64 media opt root sbin sys usr
dataVolumeContainer etc lib lost+found mnt proc run srv tmp var


[root@laptop /]# ls
bin dudir.sh lib media oprofile_data root shares @System.solv var
boot etc lib64 mnt opt run srv tmp zookeeper_server.pid
dev home lost+found myDataVolume proc sbin sys usr
```


2.查看数据卷是否挂载成功

```
[root@laptop ~]# docker ps
CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES
f71c79781ee3 centos "/bin/bash" 2 minutes ago Up 2 minutes friendly_lehmann

[root@laptop ~]# docker inspect f71c79781ee3
....
"HostConfig": {
"Binds": [
"/myDataVolume:/dataVolumeContainer"
],
....



"Mounts": [
{
"Type": "bind",
"Source": "/myDataVolume",
"Destination": "/dataVolumeContainer",
"Mode": "",
"RW": true,
"Propagation": "rprivate"
}
],
...
```



3.容器和宿主机之间数据共享

4.容器停止退出后，主机修改后数据是否同步  -- **经过测试数据是同步的**

5.命令

```
docker run -it -v /宿主机绝对路径目录:/容器内目录:ro  镜像名   #ro 只读
容器内不能写，宿主机可以写
```




### DockerFile添加

```
[root@laptop mydocker]# cat dockerfile

FROM centos
VOLUME ["/dataVolumeContainer1", "/dataVolumeContainer2"]
CMD echo "finished"
CMD /bin/bash
```

```
[root@laptop mydocker]# docker build -f dockerfile -t abc/centos .

Sending build context to Docker daemon 2.048kB
Step 1/4 : FROM centos
---> 0f3e07c0138f
Step 2/4 : VOLUME ["/dataVolumeContainer1", "/dataVolumeContainer2"]
---> Running in 97ffe5299e14
Removing intermediate container 97ffe5299e14
---> f47d556c8f0a
Step 3/4 : CMD echo "finished"
---> Running in 4bada4dcb5b5
Removing intermediate container 4bada4dcb5b5
---> 23d47518f485
Step 4/4 : CMD /bin/bash
---> Running in 9d7b4e16c397
Removing intermediate container 9d7b4e16c397
---> 0760a8000e5d
Successfully built 0760a8000e5d
Successfully tagged abc/centos:latest
```
```
[root@laptop mydocker]# docker images

REPOSITORY TAG IMAGE ID CREATED SIZE
abc/centos latest 0760a8000e5d 8 seconds ago 220MB
abc/mytomcat 1.2 f73b412690d6 About an hour ago 507MB
tomcat latest 882487b8be1d 3 weeks ago 507MB
centos latest 0f3e07c0138f 6 weeks ago 220MB


[root@laptop mydocker]# docker run -it abc/centos

[root@59f466af5f6b /]# ls -l

total 56
lrwxrwxrwx 1 root root 7 May 11 2019 bin -> usr/bin
drwxr-xr-x 2 root root 4096 Nov 15 04:24 dataVolumeContainer1
drwxr-xr-x 2 root root 4096 Nov 15 04:24 dataVolumeContainer2
...

```


宿主机目录路径：

```
[root@laptop mydocker]# docker ps

CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES
59f466af5f6b abc/centos "/bin/sh -c /bin/bash" 2 minutes ago Up 2 minutes sleepy_shaw
[root@laptop mydocker]# docker inspect 59f466af5f6b

"Volumes": {
"/dataVolumeContainer1": {},
"/dataVolumeContainer2": {}
....
]

```


## 数据卷容器

### 是什么

命名的容器挂载数据卷，其它容器通过挂载这个实现数据共享，挂载数据卷的容器，称之为数据卷容器

```
/dataVolumeContainer1
/dataVolumeContainer2
```

### 容器之间传递共享

```
--volumes-from
docker run -it xxx --volumes-from xxx  xxxx
```

容器之间配置信息的传递，数据卷的的生命周期一直持续到没有容器使用它为止



