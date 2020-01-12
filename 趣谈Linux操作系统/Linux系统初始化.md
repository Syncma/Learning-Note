# Linux系统初始化

<!-- TOC -->

- [Linux系统初始化](#linux%e7%b3%bb%e7%bb%9f%e5%88%9d%e5%a7%8b%e5%8c%96)
  - [1.加载BIOS](#1%e5%8a%a0%e8%bd%bdbios)
  - [2.读取MBR](#2%e8%af%bb%e5%8f%96mbr)
  - [3.Bootloader](#3bootloader)
  - [4.加载内核](#4%e5%8a%a0%e8%bd%bd%e5%86%85%e6%a0%b8)
  - [5.inittab](#5inittab)
  - [6.sysinit](#6sysinit)
  - [7.modules](#7modules)
  - [8.运行级别](#8%e8%bf%90%e8%a1%8c%e7%ba%a7%e5%88%ab)
  - [9.rc.local](#9rclocal)
  - [10.login](#10login)

<!-- /TOC -->

## 1.加载BIOS
**Basic Input and Output System，基本输入输出系统**

它的作用简单的说：

> BIOS中包含了CPU的相关信息、设备启动顺序信息、硬盘信息、内存信息.....

## 2.读取MBR
硬盘上第0磁道第一个扇区被称为MBR，也就是**Master Boot Record，即主引导记录**

它的大小是512字节，里面却存放了预启动信息、分区表信息。

系统找到BIOS所指定的硬盘的MBR后，就会将其复制到0×7c00地址所在的物理内存中。

其实被复制到物理内存的内容就是Boot Loader，而具体到你的电脑，那就是grub


## 3.Bootloader
Boot Loader 就是在操作系统内核运行之前运行的一段小程序。

通过这段小程序，可以初始化硬件设备、建立内存空间的映射图，现在主要是grub2
（Grand Unified Bootloader Version 2）

系统读取内存中的grub配置信息（一般为menu.lst或grub.lst），并依照此配置信息来启动不同的操作系统。


## 4.加载内核
根据grub设定的内核映像所在路径，系统读取内存映像，并进行解压缩操作

系统将解压后的内核放置在内存之中，并调用start_kernel()函数来启动一系列的初始化函数并初始化各种设备，完成Linux核心环境的建立。

至此，Linux内核已经建立起来了，基于Linux的程序应该可以正常运行了。


## 5.inittab
用户层init依据inittab文件来设定运行等级内核被加载后，第一个运行的程序便是/sbin/init，该文件会读取/etc/inittab文件，并依据此文件来进行初始化工作。

其实/etc/inittab文件最主要的作用就是设定Linux的运行等级，

其设定形式是“：id:5:initdefault:”，这就表明Linux需要运行在等级5上。

Linux的运行等级设定如下：

```
0：关机
1：单用户模式
2：无网络支持的多用户模式
3：有网络支持的多用户模式
4：保留，未使用
5：有网络支持有X-Window支持的多用户模式
6：重新引导系统，即重启
```

## 6.sysinit
init进程执行rc.sysinit在设定了运行等级后

Linux系统执行的第一个用户层文件就是/etc/rc.d/rc.sysinit脚本程序，它做的工作非常多

包括设定PATH、设定网络配置（/etc/sysconfig/network）、启动swap分区、设定/proc等等


## 7.modules
启动内核模块具体是依据/etc/modules.conf文件或/etc/modules.d目录下的文件来装载内核模块。

## 8.运行级别
执行不同运行级别的脚本程序根据运行级别的不同，系统会运行rc0.d到rc6.d中的相应的脚本程序，来完成相应的初始化工作和启动相应的服务。


## 9.rc.local
执行/etc/rc.d/rc.local你如果打开了此文件，rc.local就是在一切初始化工作后，Linux留给用户进行个性化的地方。

你可以把你想设置和启动的东西放到这里。

## 10.login
执行/bin/login程序，进入登录状态
