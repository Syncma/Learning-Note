# docker阿里云镜像加速
<!-- TOC -->

- [docker阿里云镜像加速](#docker%e9%98%bf%e9%87%8c%e4%ba%91%e9%95%9c%e5%83%8f%e5%8a%a0%e9%80%9f)
  - [步骤](#%e6%ad%a5%e9%aa%a4)
  - [docker run 逻辑](#docker-run-%e9%80%bb%e8%be%91)

<!-- /TOC -->

## 步骤
1.登录 [阿里云网站](https://cr.console.aliyun.com/cn-hangzhou/instances/mirrors)

2.注册阿里云账户

3.获取加速器地址连接

4.配置本地docker进行镜像加速器

针对Docker客户端版本大于 1.10.0 的用户
您可以通过修改daemon配置文件/etc/docker/daemon.json来使用加速器

```
[root@laptop ~]$ docker -v
Docker version 19.03.4, build 9013bf583a

[root@laptop ~]$ mkdir -p /etc/docker
[root@laptop ~]$ tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://1664le6h.mirror.aliyuncs.com"]
}
EOF

```


5.重启docker后台程序
```
[root@laptop ~]$ systemctl daemon-reload
[root@laptop ~]$ systemctl restart docker

```

6.测试
```
[root@laptop ~]# docker info
Client:
Debug Mode: false
Server:
Containers: 1
  Running: 0
  Paused: 0
  Stopped: 1
Images: 1
Server Version: 19.03.4
Storage Driver: overlay2
  Backing Filesystem: extfs
  Supports d_type: true
  Native Overlay Diff: true
Logging Driver: json-file
Cgroup Driver: cgroupfs
Plugins:
  Volume: local
  Network: bridge host ipvlan macvlan null overlay
  Log: awslogs fluentd gcplogs gelf journald json-file local logentries splunk syslog
Swarm: inactive
Runtimes: runc
Default Runtime: runc
Init Binary: docker-init
containerd version: b34a5c8af56e510852c35414db4c1f4fa6172339
runc version: 3e425f80a8c931f88e6d94a8c831b9d5aa481657
init version: fec3683
Security Options:
  seccomp
   Profile: default
Kernel Version: 4.20.14-200.fc29.x86_64
Operating System: Fedora 29 (Workstation Edition)
OSType: linux
Architecture: x86_64
CPUs: 8
Total Memory: 7.554GiB
Name: laptop
ID: 74Y4:EFIV:S63Y:USTO:V3TB:RU3J:MLXS:6WRK:RHCR:RNTB:LUX3:JGHM
Docker Root Dir: /var/lib/docker
Debug Mode: false
Registry: https://index.docker.io/v1/
Labels:
Experimental: false
Insecure Registries:
  127.0.0.0/8
Registry Mirrors:
  https://1664le6h.mirror.aliyuncs.com/
Live Restore Enabled: false

```


```
[root@laptop ~]# docker run hello-world
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
1b930d010525: Pull complete
Digest: sha256:c3b4ada4687bbaa170745b3e4dd8ac3f194ca95b2d0518b417fb47e5879d9b5f
Status: Downloaded newer image for hello-world:latest
Hello from Docker!
This message shows that your installation appears to be working correctly.
To generate this message, Docker took the following steps:
1. The Docker client contacted the Docker daemon.
2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
(amd64)
3. The Docker daemon created a new container from that image which runs the
executable that produces the output you are currently reading.
4. The Docker daemon streamed that output to the Docker client, which sent it
to your terminal.
To try something more ambitious, you can run an Ubuntu container with:
$ docker run -it ubuntu bash
Share images, automate workflows, and more with a free Docker ID:
https://hub.docker.com/
For more examples and ideas, visit:
https://docs.docker.com/get-started/
```


## docker run 逻辑

![enter image description here](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/docker-run.png)