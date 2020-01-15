# Linux系统初始化

<!-- TOC -->

- [Linux系统初始化](#linux%e7%b3%bb%e7%bb%9f%e5%88%9d%e5%a7%8b%e5%8c%96)
  - [1.加载BIOS](#1%e5%8a%a0%e8%bd%bdbios)
  - [2.读取MBR](#2%e8%af%bb%e5%8f%96mbr)
  - [3.Bootloader](#3bootloader)
  - [4.加载内核](#4%e5%8a%a0%e8%bd%bd%e5%86%85%e6%a0%b8)
    - [initrd](#initrd)
    - [initramfs](#initramfs)
      - [含义](#%e5%90%ab%e4%b9%89)
      - [测试](#%e6%b5%8b%e8%af%95)
  - [5. Systemd](#5-systemd)
  - [备注](#%e5%a4%87%e6%b3%a8)

<!-- /TOC -->

## 1.加载BIOS
**Basic Input and Output System，基本输入输出系统**

它的作用简单的说：

> BIOS中包含了CPU的相关信息、设备启动顺序信息、硬盘信息、内存信息.....


BIOS 主要做了两件事：

1. **`自检`**：负责检测系统外围关键设备（如：CPU、内存、显卡、I/O、键盘鼠标等）是否正常。例如，最常见的是内存松动的情况，BIOS自检阶段会报错，系统就无法启动起来
<br>
2. **`根据我们在BIOS中设置的系统启动顺序来搜索用于启动系统的驱动器`**，如硬盘、光盘、U盘、软盘和网络等。

	我们以硬盘启动为例，BIOS此时去读取硬盘驱动器的第一个扇区(MBR，512字节)，然后执行里面的代码。
	
	实际上这里BIOS并不关心启动设备第一个扇区中是什么内容，它只是负责读取该扇区内容、并执行。

## 2.读取MBR
硬盘上第0磁道第一个扇区被称为MBR，也就是`Master Boot Record，即主引导记录`

它由三个部分组成：

>主引导程序(Bootloader)

> 硬盘分区表DPT（Disk Partition table）
> 
> 硬盘有效标志（55AA）

系统找到BIOS所指定的硬盘的MBR后，就会将其复制到0×7c00地址所在的物理内存中。

其实被复制到物理内存的内容就是Boot Loader，而具体到你的电脑，那就是grub


## 3.Bootloader
Boot Loader 就是在操作系统内核运行之前运行的一段小程序。

通过这段小程序，可以初始化硬件设备、建立内存空间的映射图，现在主要是grub2
（Grand Unified Bootloader Version 2）

系统读取内存中的grub配置信息（一般为menu.lst或grub.lst），并依照此配置信息来启动不同的操作系统。

Fedora 系统grub 路径：
> /boot/efi/EFI/fedora/grub.cfg

**grub引导也分为两个阶段stage1阶段和stage2阶段, 这里就不详细说明了**


## 4.加载内核
根据grub设定的内核映像所在路径，系统读取内存映像，并进行解压缩操作

系统将解压后的内核放置在内存之中，并调用start_kernel()函数来启动一系列的初始化函数并初始化各种设备，完成Linux核心环境的建立。


关于Linux的设备驱动程序的加载：
>有一部分驱动程序直接被编译进内核镜像中

>另一部分驱动程序则是以模块的形式放在initrd(ramdisk)中

### initrd
bootloader initialized RAM disk，就是由 boot loader 初始化的内存盘


### initramfs

#### 含义

```
在内核镜像中附加一个cpio包，这个cpio包中包含了一个小型的文件系统

当内核启动时，内核将这个 cpio包解开，并且将其中包含的文件系统释放到rootfs中，内核中的一部分初始化代码会放到这个文件系统中，作为用户层进程来执行。

这样带来的明显的好处是精简了内核的初始化代码，而且使得内核的初始化过程更容易定制。
```
[root@laptop fedora]# pwd

/boot/efi/EFI/fedora

[root@laptop fedora]# cat grub.cfg
```
...
### BEGIN /etc/grub.d/10_linux ###
menuentry 'Fedora (5.3.11-100.fc29.x86_64) 29 (Workstation Edition)' --class fedora --class gnu-linux --class gnu --class os --unrestricted $menuentry_id_option 'gnulinux-4.20.14-200.fc29.x86_64-advanced-b0a79884-e88d-4740-983a-438c88ab902e' {
    load_video
    set gfxpayload=keep
    insmod gzio
    insmod part_gpt
    insmod ext2
    if [ x$feature_platform_search_hint = xy ]; then
      search --no-floppy --fs-uuid --set=root  b0a79884-e88d-4740-983a-438c88ab902e
    else
      search --no-floppy --fs-uuid --set=root b0a79884-e88d-4740-983a-438c88ab902e
    fi  
    linuxefi    /boot/vmlinuz-5.3.11-100.fc29.x86_64 root=UUID=b0a79884-e88d-4740-983a-438c88ab902e ro resume=UUID=ec00225d-e4fb-4f85-b3d9-c53201ec55c0 rhgb quiet LANG=en_US.UTF-8
    initrdefi /boot/initramfs-5.3.11-100.fc29.x86_64.img
}
```

可以看到initramfs路径在/boot

```
[root@laptop boot]# ll |grep init
-rw-------  1 root root 25360951 Dec 26 23:13 initramfs-5.3.11-100.fc29.x86_64.img

[root@laptop boot]# uname -a
Linux laptop 5.3.11-100.fc29.x86_64 #1 SMP Tue Nov 12 20:41:25 UTC 2019 x86_64 x86_64 x86_64 GNU/Linux
```

#### 测试

这里我们做个测试看看initramfs里面到底是啥

测试环境: Fedora 29 x64

1.创建测试目录
```

[root@laptop tmp]# pwd
/tmp
[root@laptop tmp]# mkdir test
[root@laptop test]# cd test
[root@laptop test]# pwd
/tmp/test
```

2.拷贝文件到测试目录
```
[root@laptop test]# cp /boot/initramfs-5.3.11-100.fc29.x86_64.img .
[root@laptop test]# file initramfs-5.3.11-100.fc29.x86_64.img 
initramfs-5.3.11-100.fc29.x86_64.img: ASCII cpio archive (SVR4 with no CRC)
```

3.解压
```
# 参数：
# -i, --extract Extract files from an archive
# -d, --make-directories Create leading directories where needed
[root@laptop test]# cpio -id < initramfs-5.3.11-100.fc29.x86_64.img 
198 blocks
[root@laptop test]# ll
total 24772
-rw-r--r-- 1 root root        2 Jan 15 11:10 early_cpio
-rw------- 1 root root 25360951 Jan 15 11:08 initramfs-5.3.11-100.fc29.x86_64.img
drwxr-xr-x 3 root root       60 Jan 15 11:10 kernel

[root@laptop test]# tree kernel/
kernel/
└── x86
    └── microcode
        └── GenuineIntel.bin

2 directories, 1 file
```

4.问题： **`怎么出现了这个玩意？不是说好的文件系统呢？`**

这里我们下载个工具binwalk，这个工具作用是自动完成指定文件的扫描：

1) 安装：
```
[root@laptop test]# dnf install binwalk
```

2）执行方法：
```
[root@laptop test]# binwalk initramfs-5.3.11-100.fc29.x86_64.img 

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
0             0x0             ASCII cpio archive (SVR4 with no CRC), file name: ".", file name length: "0x00000002", file size: "0x00000000"
112           0x70            ASCII cpio archive (SVR4 with no CRC), file name: "early_cpio", file name length: "0x0000000B", file size: "0x00000002"
240           0xF0            ASCII cpio archive (SVR4 with no CRC), file name: "kernel", file name length: "0x00000007", file size: "0x00000000"
360           0x168           ASCII cpio archive (SVR4 with no CRC), file name: "kernel/x86", file name length: "0x0000000B", file size: "0x00000000"
484           0x1E4           ASCII cpio archive (SVR4 with no CRC), file name: "kernel/x86/microcode", file name length: "0x00000015", file size: "0x00000000"
616           0x268           ASCII cpio archive (SVR4 with no CRC), file name: "kernel/x86/microcode/GenuineIntel.bin", file name length: "0x00000026", file size: "0x00018800"
101116        0x18AFC         ASCII cpio archive (SVR4 with no CRC), file name: "TRAILER!!!", file name length: "0x0000000B", file size: "0x00000000"
101376        0x18C00         gzip compressed data, maximum compression, from Unix, NULL date (1970-01-01 00:00:00)
8747022       0x85780E        xz compressed data
```

从上面的结果可以看到101376 有个gzip包, 也可以实现下面的命令直接过滤gzip

```
[root@laptop test]# binwalk -y gzip initramfs-5.3.11-100.fc29.x86_64.img 

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
101376        0x18C00         gzip compressed data, maximum compression, from Unix, NULL date (1970-01-01 00:00:00)
```

使用这个命令来查看gzip包的内容：
```
root@laptop test]# dd if=initramfs-5.3.11-100.fc29.x86_64.img bs=101376 skip=1 |zcat |cpio -id 
249+1 records in
249+1 records out
25259575 bytes (25 MB, 24 MiB) copied, 0.51821 s, 48.7 MB/s
114798 blocks
[root@laptop test]# ll
total 24772
lrwxrwxrwx  1 root root        7 Jan 15 12:46 bin -> usr/bin
drwxr-xr-x  2 root root      140 Jan 15 12:46 dev
drwxr-xr-x 11 root root      560 Jan 15 12:46 etc
lrwxrwxrwx  1 root root       23 Jan 15 12:46 init -> usr/lib/systemd/systemd
-rw-------  1 root root 25360951 Jan 15 12:41 initramfs-5.3.11-100.fc29.x86_64.img
lrwxrwxrwx  1 root root        7 Jan 15 12:46 lib -> usr/lib
lrwxrwxrwx  1 root root        9 Jan 15 12:46 lib64 -> usr/lib64
drwxr-xr-x  2 root root       40 Jan 15 12:46 proc
drwxr-xr-x  2 root root       40 Jan 15 12:46 root
drwxr-xr-x  2 root root       40 Jan 15 12:46 run
lrwxrwxrwx  1 root root        8 Jan 15 12:46 sbin -> usr/sbin
-rwxr-xr-x  1 root root     3121 Jan 15 12:46 shutdown
drwxr-xr-x  2 root root       40 Jan 15 12:46 sys
drwxr-xr-x  2 root root       40 Jan 15 12:46 sysroot
drwxr-xr-x  2 root root       40 Jan 15 12:46 tmp
drwxr-xr-x  8 root root      160 Jan 15 12:46 usr
drwxr-xr-x  3 root root      100 Jan 15 12:46 var
```


当然还有个更简单的办法：
```
[root@laptop test] (cpio -id; zcat | cpio -id) < initramfs-5.3.11-100.fc29.x86_64.img
```

这样就可以看到文件系统了，在这里我们看到init文件

```
lrwxrwxrwx  1 root root       23 Jan 15 12:46 init -> usr/lib/systemd/systemd
```

这个init文件其实是连接到usr/lib/systemd/systemd， 这个systemd是什么东西？


## 5. Systemd


内核会去执行initrd中的init脚本，这时内核将控制权交给了init文件处理。

至此，Linux内核已经建立起来了，基于Linux的程序应该可以正常运行了。

Linux 内核通过执行 init 将 CPU 的控制权限，交给其它的任务，在 CentOS 中可以通过如下命令查看 init 来自于那个包

```
[jian@laptop boot]$ rpm -qif `which init`
Name        : systemd
Version     : 239
Release     : 14.git33ccd62.fc29
Architecture: x86_64
Install Date: Thu 26 Dec 2019 11:10:44 PM CST
Group       : Unspecified
Size        : 12261887
License     : LGPLv2+ and MIT and GPLv2+
Signature   : RSA/SHA256, Tue 03 Sep 2019 09:42:55 PM CST, Key ID a20aa56b429476b4
Source RPM  : systemd-239-14.git33ccd62.fc29.src.rpm
Build Date  : Tue 03 Sep 2019 07:44:49 PM CST
Build Host  : buildhw-10.phx2.fedoraproject.org
Relocations : (not relocatable)
Packager    : Fedora Project
Vendor      : Fedora Project
URL         : https://www.freedesktop.org/wiki/Software/systemd
Bug URL     : https://bugz.fedoraproject.org/systemd
Summary     : System and Service Manager
```


可以参考：

[阮一峰Systemd 入门教程](http://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-commands.html)

[Systemd 入门教程：命令篇](http://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-part-two.html)


[Systemd介绍](https://jin-yang.github.io/post/linux-systemd.html)


[Systemd入门](https://80imike.github.io/posts/3761.html#journalctl)


## 备注

这里简单说下/boot命令底下的文件都是什么？

```
[jian@laptop boot]$ pwd
/boot

jian@laptop boot]$ ls --format=single-column .
config-4.20.14-200.fc29.x86_64
config-5.3.11-100.fc29.x86_64  // 系统kernel的配置文件，内核编译完成后保存的就是这个配置文件
efi   // Extensible Firmware Interface（EFI，可扩展固件接口）是 Intel 为全新类型的 PC 固件的体系结构、接口和服务提出的建议标准。
elf-memtest86+-5.01
extlinux
grub2  //开机管理程序grub相关数据目录
initramfs-0-rescue-17032e20319d44259de0eb5081c2092f.img
initramfs-4.20.14-200.fc29.x86_64.img
initramfs-5.3.11-100.fc29.x86_64.img //虚拟文件系统文件（用initramfs代替了initrd，他们的目的是一样的，只是本身处理的方式有点不同）
loader
memtest86+-5.01
System.map-4.20.14-200.fc29.x86_64
System.map-5.3.11-100.fc29.x86_64 //是系统kernel中的变量对应表；(也可以理解为是索引文件)
vmlinuz-0-rescue-17032e20319d44259de0eb5081c2092f
vmlinuz-4.20.14-200.fc29.x86_64
vmlinuz-5.3.11-100.fc29.x86_64 //系统使用kernel，用于启动的压缩内核镜像


```