# 安装docker

这里使用fedora系统来安装

```
[root@laptop ~]# dnf config-manager --add-repo https://download.docker.com/linux/fedora/docker-ce.repo

[root@laptop ~]# dnf makecache
[root@laptop ~]# dnf install docker-ce

After successful installation of Docker engine, Let’s enable and start the docker service.

[root@laptop ~]# systemctl enable docker.service
[root@laptop ~]# systemctl start docker.service

[root@laptop ~]# docker search fedora
```

[参考地址](https://tecadmin.net/install-docker-on-fedora/)