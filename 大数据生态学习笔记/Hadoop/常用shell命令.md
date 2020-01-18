# 常用shell 命令

<!-- TOC -->

- [常用shell 命令](#常用shell-命令)
    - [显示当前目录结构](#显示当前目录结构)
    - [创建目录](#创建目录)
    - [删除操作](#删除操作)
    - [从本地加载文件到 HDFS](#从本地加载文件到-hdfs)
    - [从 HDFS 导出文件到本地](#从-hdfs-导出文件到本地)
    - [查看文件内容](#查看文件内容)
    - [显示文件的最后内容](#显示文件的最后内容)
    - [拷贝文件](#拷贝文件)
    - [移动文件](#移动文件)
    - [统计当前目录下各文件大小](#统计当前目录下各文件大小)
    - [合并下载多个文件](#合并下载多个文件)
    - [统计文件系统的可用空间信息](#统计文件系统的可用空间信息)
    - [更改文件复制因子](#更改文件复制因子)
    - [权限控制](#权限控制)
    - [文件检测](#文件检测)

<!-- /TOC -->

## 显示当前目录结构

```
# 显示当前目录结构
hadoop fs -ls <path>
# 递归显示当前目录结构
hadoop fs -ls -R <path>
# 显示根目录下内容
hadoop fs -ls /
```
## 创建目录
```
# 创建目录
hadoop fs -mkdir <path>
# 递归创建目录
hadoop fs -mkdir -p <path>
```

## 删除操作
```
# 删除文件
hadoop fs -rm <path>
# 递归删除目录和文件
hadoop fs -rm -R <path>
```

## 从本地加载文件到 HDFS
```
# 二选一执行即可
hadoop fs -put [localsrc] [dst]
hadoop fs - copyFromLocal [localsrc] [dst]
```

## 从 HDFS 导出文件到本地
```
# 二选一执行即可
hadoop fs -get [dst] [localsrc]
hadoop fs -copyToLocal [dst] [localsrc]
```

## 查看文件内容
```
# 二选一执行即可
hadoop fs -text <path>
hadoop fs -cat <path>
```

## 显示文件的最后内容
```
hadoop fs -tail <path>
# 和Linux下一样，会持续监听文件内容变化 并显示文件的最后一千字节
hadoop fs -tail -f <path>
```

## 拷贝文件
```
hadoop fs -cp [src] [dst]
```

## 移动文件
```
hadoop fs -mv [src] [dst]
```

## 统计当前目录下各文件大小
```
	* 默认单位字节
	* -s : 显示所有文件大小总和，
	* -h : 将以更友好的方式显示文件大小（例如 64.0m 而不是 67108864）

hadoop fs -du <path>
```
## 合并下载多个文件
```
	* -nl 在每个文件的末尾添加换行符（LF）
	* -skip-empty-file 跳过空文件
	
hadoop fs -getmerge
# 示例 将HDFS上的hbase-policy.xml和hbase-site.xml文件合并后下载到本地的/usr/test.xml
hadoop fs -getmerge -nl /test/hbase-policy.xml /test/hbase-site.xml /usr/test.xml
```
## 统计文件系统的可用空间信息
```
hadoop fs -df -h /
```

## 更改文件复制因子
```
hadoop fs -setrep [-R] [-w] <numReplicas> <path>
	* 更改文件的复制因子。如果 path 是目录，则更改其下所有文件的复制因子
	* -w : 请求命令是否等待复制完成
# 示例
hadoop fs -setrep -w 3 /user/hadoop/dir1
```
## 权限控制
```
# 权限控制和Linux上使用方式一致
# 变更文件或目录的所属群组。 用户必须是文件的所有者或超级用户。
hadoop fs -chgrp [-R] GROUP URI [URI ...]
# 修改文件或目录的访问权限 用户必须是文件的所有者或超级用户。
hadoop fs -chmod [-R] <MODE[,MODE]... | OCTALMODE> URI [URI ...]
# 修改文件的拥有者 用户必须是超级用户。
hadoop fs -chown [-R] [OWNER][:[GROUP]] URI [URI ]
```
## 文件检测
```
hadoop fs -test - [defsz] URI
可选选项：
	* -d：如果路径是目录，返回 0。
	* -e：如果路径存在，则返回 0。
	* -f：如果路径是文件，则返回 0。
	* -s：如果路径不为空，则返回 0。
	* -r：如果路径存在且授予读权限，则返回 0。
	* -w：如果路径存在且授予写入权限，则返回 0。
	* -z：如果文件长度为零，则返回 0。

# 示例
hadoop fs -test -e filename
```