# YRAN配置历史服务器

<!-- TOC -->

- [YRAN配置历史服务器](#yran%e9%85%8d%e7%bd%ae%e5%8e%86%e5%8f%b2%e6%9c%8d%e5%8a%a1%e5%99%a8)
  - [问题](#%e9%97%ae%e9%a2%98)
  - [解决办法](#%e8%a7%a3%e5%86%b3%e5%8a%9e%e6%b3%95)

<!-- /TOC -->


## 问题
为了查看程序的历史运行情况，需要配置历史服务器

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/yarn-history.png)


## 解决办法
1.修改配置文件

```
[jian@laptop hadoop]$ pwd
/home/jian/prj/bigdata/hadoop-2.10.0/etc/hadoop
[jian@laptop hadoop]$ vi mapred-site.xml
[jian@laptop hadoop]$ cat mapred-site.xml
<?xml version="1.0"?>
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
    <!--指定MR运行在YARN上-->
    <property>
        <name>mapreduce.framework.name</name>
        <value>yarn</value>
    </property>
    <!--历史服务端地址-->
    <property>
        <name>mapreduce.jobhistory.address</name>
        <value>0.0.0.0:10020</value>
    </property>
    <!--历史服务器web端地址-->
    <property>
        <name>mapreduce.jobhistory.webapp.address</name>
        <value>0.0.0.0:19888</value>
    </property>
</configuration>
```

2.启动服务

```
[jian@laptop hadoop-2.10.0]$ sbin/mr-jobhistory-daemon.sh start historyserver
starting historyserver, logging to /home/jian/prj/bigdata/hadoop-2.10.0/logs/mapred-jian-historyserver-laptop.out

[jian@laptop hadoop-2.10.0]$ jps
14739 JobHistoryServer
31987 NameNode
14964 Jps
3238 DataNode
2137 NodeManager
10587 ResourceManager

```


3.测试，点击history就可以打开页面：


![enter image description here](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/yarn-h2.png)
