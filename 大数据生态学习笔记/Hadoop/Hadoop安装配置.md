# Hadoop安装配置

<!-- TOC -->

- [Hadoop安装配置](#hadoop%e5%ae%89%e8%a3%85%e9%85%8d%e7%bd%ae)
  - [安装](#%e5%ae%89%e8%a3%85)
  - [docker 安装](#docker-%e5%ae%89%e8%a3%85)
  - [目录结构](#%e7%9b%ae%e5%bd%95%e7%bb%93%e6%9e%84)
  - [本地模式](#%e6%9c%ac%e5%9c%b0%e6%a8%a1%e5%bc%8f)
  - [伪分布模式](#%e4%bc%aa%e5%88%86%e5%b8%83%e6%a8%a1%e5%bc%8f)

<!-- /TOC -->




## 安装
测试环境： Fedora 29 x64

1.下载hadoop-2.10.0.tar.gz，并解压到本地路径

可以使用下面的源进行安装：
[中科大源](http://mirrors.ustc.edu.cn/apache/hadoop/core/hadoop-2.10.0/)

这里使用系统自带的openjdk （**`发现不行，有些命令没有`**）

这里**`推荐使用官网jdk,这里下载13.0.1 linux版本`**
[官网地址](https://www.oracle.com/technetwork/java/javase/downloads/index.html)
```
[jian@laptop tmp]$ tar xf jdk-13.0.1_linux-x64_bin.tar.gz -C  ~/prj/bigdata/
```

2.设置环境变量：

 vi ~/.bash_profile 或者~/.bashrc

新增下面几行内容： **`这里要按照自己的hadoop路径进行相关配置`**
```
#JAVA_HOME
export JAVA_HOME=/home/jian/prj/bigdata/jdk1.8.0_231
export PATH=$PATH:$JAVA_HOME/bin

HADOOP_HOME=/home/jian/prj/bigdata/hadoop-2.10.0
export HADOOP_HOME

PATH=$HADOOP_HOME/bin:$HADOOP_HOME/sbin:$PATH
export PATH
```
           
使用source命令生效环境变量：
```
 source ~/.bash_profile
```
   

3.执行hadoop命令看看效果
```
[jian@laptop share]$ hadoop
Usage: hadoop [--config confdir] [COMMAND | CLASSNAME]
 CLASSNAME            run the class named CLASSNAME
or
 where COMMAND is one of:
 fs                   run a generic filesystem user client
 version              print the version
 jar <jar>            run a jar file
                      note: please use "yarn jar" to launch
                            YARN applications, not this command.
 checknative [-a|-h]  check native hadoop and compression libraries availability
 distcp <srcurl> <desturl> copy file or directories recursively
 archive -archiveName NAME -p <parent path> <src>* <dest> create a hadoop archive
 classpath            prints the class path needed to get the
                      Hadoop jar and the required libraries
 credential           interact with credential providers
 daemonlog            get/set the log level for each daemon
 trace                view and modify Hadoop tracing settings
Most commands print help when invoked w/o parameters.

```
## docker 安装

* 内容待补充

## 目录结构

```
[jian@laptop hadoop-2.10.0]$ ll
total 152
drwxr-xr-x 2 jian jian   4096 Oct 23 03:23 bin
drwxr-xr-x 3 jian jian   4096 Oct 23 03:23 etc
drwxr-xr-x 2 jian jian   4096 Oct 23 03:23 include
drwxr-xr-x 3 jian jian   4096 Oct 23 03:23 lib
drwxr-xr-x 2 jian jian   4096 Oct 23 03:23 libexec
-rw-r--r-- 1 jian jian 106210 Oct 23 03:23 LICENSE.txt
-rw-r--r-- 1 jian jian  15841 Oct 23 03:23 NOTICE.txt
-rw-r--r-- 1 jian jian   1366 Oct 23 03:23 README.txt
drwxr-xr-x 3 jian jian   4096 Oct 23 03:23 sbin
drwxr-xr-x 4 jian jian   4096 Oct 23 03:23 share
```


## 本地模式

[官方文档](https://hadoop.apache.org/docs/stable/hadoop-project-dist/hadoop-common/SingleCluster.html)

  特点：**`不具备HDFS，只能测试MapReduce程序`**


1.grep案例演示：
```
[jian@laptop output] $ mkdir input
[jian@laptop output] $ cp etc/hadoop/*.xml input
[jian@laptop output] $ bin/hadoop jar share/hadoop/mapreduce/hadoop-mapreduce-examples-3.2.1.jar grep input output 'dfs[a-z.]+'
[jian@laptop output]  $ cat output/*
[jian@laptop output]$ ll
total 4
-rw-r--r-- 1 jian jian 11 Nov  9 14:29 part-r-00000
-rw-r--r-- 1 jian jian  0 Nov  9 14:29 _SUCCESS
[jian@laptop output]$ cat part-r-00000
1 dfsadmin
```

这里要注意：**`如果output目录存在就会报错`**

```
xxxxx
19/11/09 14:38:50 INFO jvm.JvmMetrics: Cannot initialize JVM Metrics with processName=JobTracker, sessionId= - already initialized
org.apache.hadoop.mapred.FileAlreadyExistsException: Output directory file:/home/jian/prj/bigdata/hadoop-2.10.0/output already exists
```

**所以`output目录是不需要创建的`**

2.wordcount案例演示：

```
[jian@laptop hadoop-2.10.0]$ mkdir wcinput
[jian@laptop hadoop-2.10.0]$ cd wcinput/
[jian@laptop wcinput]$ vi wc.input
[jian@laptop wcinput]$ cat wc.input
hadoop yarn
hadoop mapreduce
hello
hello123
[jian@laptop wcinput]$ cd ..
[jian@laptop hadoop-2.10.0]$ bin/hadoop jar share/hadoop/mapreduce/hadoop-mapreduce-examples-2.10.0.jar wordcount wcinput/ wcoutput
WARNING: An illegal reflective access operation has occurred
xxxx

[jian@laptop hadoop-2.10.0]$ cat wcoutput/part-r-00000
hadoop 2
hello 1
hello123 1
mapreduce 1
yarn 1
```
       
 注意：
```
MR有一个默认的排序规则， 具体的规则是什么？ 也可以自定义排序规则的 
具体可以网上搜索下
```

## 伪分布模式 

 特点：**`具备Hadoop的所有功能，在单机上模拟一个分布式的环境`**

```
（1）HDFS：主：NameNode，数据节点：DataNode
（2）Yarn：容器，运行MapReduce程序
	主节点：ResourceManager
	从节点：NodeManager
```

修改配置文件：
```
[jian@laptop hadoop]$ pwd
/home/jian/prj/bigdata/hadoop-2.10.0/etc/hadoop

[jian@laptop hadoop]$ cat core-site.xml
<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet type="text/xsl" href="configuration.xsl"?>
<!--
  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License. See accompanying LICENSE file.
-->
<!-- Put site-specific property overrides in this file. -->
<configuration>
   <!-- 指定HDFS中NameNode的地址-->
    <property>
        <name>fs.defaultFS</name>
        <value>hdfs://localhost:9000</value>
    </property>
   <!-- 指定Hadoop中运行时产生文件的存储目录-->
    <property>
        <name>hadoop.tmp.dir</name>
		<value>/home/jian/prj/bigdata/hadoop-2.10.0/data/tmp</value>
    </property>
</configuration>
```

```
[jian@laptop hadoop]$ cat hdfs-site.xml

<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet type="text/xsl" href="configuration.xsl"?>
<!--
  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License. See accompanying LICENSE file.
-->
<!-- Put site-specific property overrides in this file. -->
<configuration>
    <!--指定HDFS副本的数量--> //默认值是3，如果只有一台机器，副本不是在本地存三份， 数据也只会存一份，要特别注意
    <property>
        <name>dfs.replication</name>
        <value>1</value>
    </property>
</configuration>
```

修改JAVA_HOME:

```
[jian@laptop hadoop]$ vi hadoop-env.sh
[jian@laptop hadoop]$ grep "JAVA_HOME" hadoop-env.sh
# The only required environment variable is JAVA_HOME.  All others are
# set JAVA_HOME in this file, so that it is correctly defined on
export JAVA_HOME=/home/jian/prj/bigdata/jdk1.8.0_231

```


启动集群
1.**`格式化NameNode 第一次启动时候格式化，以后就不要总格式化`**

```
[jian@laptop hadoop]$ hadoop namenode -format
```

原因：
```
格式化NameNode, 会产生新的集群ID, 导致NameNode和DataNode的集群ID不一致，集群找不到已往数据。

所以在格式NameNode时候，一定要先删除data数据和log日志，然后再格式化NameNode
```

NameNode数据：
```
[jian@laptop current]$ pwd
/home/jian/prj/bigdata/hadoop-2.10.0/data/tmp/dfs/name/current
[jian@laptop current]$ cat VERSION
#Sat Nov 09 15:16:26 CST 2019
namespaceID=852948635
blockpoolID=BP-2024675374-192.168.2.194-1573283786692
storageType=NAME_NODE
cTime=1573283786692
clusterID=CID-57f23117-5c60-4dcd-b5d8-3cb13e5ecdc9
layoutVersion=-63
```

DataNode数据：
```
[jian@laptop current]$ pwd
/home/jian/prj/bigdata/hadoop-2.10.0/data/tmp/dfs/data/current
[jian@laptop current]$ cat VERSION
#Sat Nov 09 15:32:06 CST 2019
datanodeUuid=84c8080a-14a4-4002-96f5-e0259373a424
storageType=DATA_NODE
cTime=0
clusterID=CID-57f23117-5c60-4dcd-b5d8-3cb13e5ecdc9
layoutVersion=-57
storageID=DS-38d8f934-48ea-4c10-a16d-d16de6b8e887
```


这一步是做什么用？
```
这种格式化HDFS的方式是需要把原来HDFS中的数据全部清空，
然后再格式化并安装一个全新的HDFS。

注：
这种格式化方式需要将HDFS中的数据全部清空
```

```
[jian@laptop hadoop-2.10.0]$ pwd
/home/jian/prj/bigdata/hadoop-2.10.0
[jian@laptop hadoop-2.10.0]$ bin/hdfs namenode -format
```

 2.启动NameNode
```
[jian@laptop hadoop-2.10.0]$ sbin/hadoop-daemon.sh start namenode
```
3.启动DataNode
```
[jian@laptop hadoop-2.10.0]$ sbin/hadoop-daemon.sh start datanode
```

查看进程情况：

```
[jian@laptop hadoop-2.10.0]$ jps
31987 NameNode
3238 DataNode
3624 Jps
```

也可以直接执行下面的命令，这个sh脚本会自动把namenode,datanode一起开启：
```
[jian@laptop hadoop-2.10.0]$ sbin/start-all.sh
```


使用浏览器访问： [这里](http://localhost:50070)

查看是否像下面一样：

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/hadoop.png)



创建目录：
```
[jian@laptop hadoop-2.10.0]$ bin/hdfs dfs -mkdir -p /user/jian/input
```

查看目录：
```
[jian@laptop hadoop-2.10.0]$ bin/hdfs dfs  -ls
WARNING: An illegal reflective access operation has occurred
WARNING: Illegal reflective access by org.apache.hadoop.security.authentication.util.KerberosUtil (file:/home/jian/prj/bigdata/hadoop-2.10.0/share/hadoop/common/lib/hadoop-auth-2.10.0.jar) to method sun.security.krb5.Config.getInstance()
WARNING: Please consider reporting this to the maintainers of org.apache.hadoop.security.authentication.util.KerberosUtil
WARNING: Use --illegal-access=warn to enable warnings of further illegal reflective access operations
WARNING: All illegal access operations will be denied in a future release
Found 1 items
drwxr-xr-x   - jian supergroup          0 2019-11-09 15:37 input
```
**`会出现好多Warning信息， 这时因为jdk版本过高了，把jdk降低到1.8及以下，建议1.8 (jdk-8u231-linux.tar.gz)`**

```
[jian@laptop hadoop-2.10.0]$ bin/hdfs dfs  -ls
Found 1 items
drwxr-xr-x   - jian supergroup          0 2019-11-09 15:37 input
```

本地文件上传到HDFS里面：
```
[jian@laptop hadoop-2.10.0]$ bin/hdfs dfs -put wcinput/wc.input /user/jian/input

[jian@laptop hadoop-2.10.0]$ bin/hadoop jar share/hadoop/mapreduce/hadoop-mapreduce-examples-2.10.0.jar wordcount /user/jian/input /user/jian/output
```

查看文件
```
[jian@laptop hadoop-2.10.0]$ bin/hdfs dfs -cat /user/jian/output/*
hadoop 2
hello 1
hello123 1
mapreduce 1
yarn 1
```

可以在浏览器中看到是否像下面显示一样：

![enter image description here](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/hadoop-dir.png)
