# YARN日志聚集功能

<!-- TOC -->

- [YARN日志聚集功能](#yarn%e6%97%a5%e5%bf%97%e8%81%9a%e9%9b%86%e5%8a%9f%e8%83%bd)
  - [问题](#%e9%97%ae%e9%a2%98)
  - [解决办法](#%e8%a7%a3%e5%86%b3%e5%8a%9e%e6%b3%95)

<!-- /TOC -->


## 问题
日志聚集概念： 应用运行完成后，将程序运行日志信息上传到HDFS系统上
日志聚集功能好处：可以方便的查看到程序运行详情，方便开发调试

开启日志聚集功能，**需要重新启动NodeManager, ResourceManager和HistoryManager**

![enter image description here](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/yarn-log.png)


## 解决办法

1.配置

```
[jian@laptop hadoop]$ pwd
/home/jian/prj/bigdata/hadoop-2.10.0/etc/hadoop
[jian@laptop hadoop]$ cat yarn-site.xml
<?xml version="1.0"?>
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
<configuration>
<!-- Site specific YARN configuration properties -->
    <!-- Reducer 获取数据的方式-->
<property>
    </property>
    <!-- 指定YARN的ResourceManager的地址-->
<property>
<name>yarn.resourcemanager.hostname</name>
            <value>0.0.0.0</value>
    </property>
    <!--日志聚集功能-->
    <property>
        <name>yarn.log-aggregation-enable</name>
        <value>true</value>
    </property>
    <!--日志时间保留7天-->
    <property>
        <name>yarn.log-aggregation.retain-seconds</name>
        <value>604800</value>
    </property>


</configuration>
```

2.关闭服务

```
[jian@laptop hadoop-2.10.0]$ sbin/mr-jobhistory-daemon.sh stop historyserver
stopping historyserver

[jian@laptop hadoop-2.10.0]$ sbin/yarn-daemon.sh stop nodemanager
stopping nodemanager
nodemanager did not stop gracefully after 5 seconds: killing with kill -9

[jian@laptop hadoop-2.10.0]$ sbin/yarn-daemon.sh stop resourcemanager
stopping resourcemanager
[jian@laptop hadoop]$ jps
31987 NameNode
17908 Jps
3238 DataNode
```

3.启动服务

```
[jian@laptop hadoop-2.10.0]$ sbin/yarn-daemon.sh start resourcemanager
starting resourcemanager, logging to /home/jian/prj/bigdata/hadoop-2.10.0/logs/yarn-jian-resourcemanager-laptop.out
[jian@laptop hadoop-2.10.0]$ sbin/yarn-daemon.sh start nodemanager
starting nodemanager, logging to /home/jian/prj/bigdata/hadoop-2.10.0/logs/yarn-jian-nodemanager-laptop.out
[jian@laptop hadoop-2.10.0]$ sbin/mr-jobhistory-daemon.sh start historyserver
starting historyserver, logging to /home/jian/prj/bigdata/hadoop-2.10.0/logs/mapred-jian-historyserver-laptop.out


[jian@laptop hadoop-2.10.0]$ jps
19858 NodeManager
31987 NameNode
3238 DataNode
20378 JobHistoryServer
21164 Jps
19439 ResourceManager
```

4.测试
```
[jian@laptop hadoop-2.10.0]$ hdfs dfs -rm -r  /user/jian/output
Deleted /user/jian/output

[jian@laptop hadoop-2.10.0]$ hadoop jar share/hadoop/mapreduce/hadoop-mapreduce-examples-2.10.0.jar wordcount  /user/jian/input /user/jian/output
```

打开页面：

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/yarn-log2.png)