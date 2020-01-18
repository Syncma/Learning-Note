# 启动YARN

<!-- TOC -->

- [启动YARN](#%e5%90%af%e5%8a%a8yarn)
  - [配置](#%e9%85%8d%e7%bd%ae)

<!-- /TOC -->

## 配置

1.配置yarn-env.sh
```
[jian@laptop hadoop]$ pwd
/home/jian/prj/bigdata/hadoop-2.10.0/etc/hadoop
[jian@laptop hadoop]$ grep "JAVA_HOME" yarn-env.sh
# export JAVA_HOME=/home/y/libexec/jdk1.6.0/
export JAVA_HOME=/home/jian/prj/bigdata/jdk1.8.0_231
```

2.配置yarn-site.xml
```
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
<name>yarn.nodemanager.aux-services</name>
<value>mapreduce_shuffle</value>
</property>
    <!-- 指定YARN的ResourceManager的地址-->
<property>
<name>yarn.resourcemanager.hostname</name>
            <value>0.0.0.0</value>
</property>


</configuration>

```
3.配置mapred-env.sh
```
[jian@laptop hadoop]$ vi mapred-env.sh
[jian@laptop hadoop]$ grep "JAVA_HOME"  mapred-env.sh
# export JAVA_HOME=/home/y/libexec/jdk1.6.0/
export JAVA_HOME=/home/jian/prj/bigdata/jdk1.8.0_231
```


4.配置mapred-site.xml
```
[jian@laptop hadoop]$ mv mapred-site.xml.template mapred-site.xml
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
</configuration>
```



5.启动集群

启动NameNode, DataNode
```
[jian@laptop hadoop-2.10.0]$ sbin/hadoop-daemon.sh start namenode
[jian@laptop hadoop-2.10.0]$ sbin/hadoop-daemon.sh start datanode

[jian@laptop hadoop]$ jps
31987 NameNode
3238 DataNode
28522 Jps
```

启动ResourceManager
```
[jian@laptop hadoop-2.10.0]$ sbin/yarn-daemon.sh  start resourcemanager
starting resourcemanager, logging to /home/jian/prj/bigdata/hadoop-2.10.0/logs/yarn-jian-resourcemanager-laptop.out
```

启动NodeManager
```
[jian@laptop hadoop-2.10.0]$ sbin/yarn-daemon.sh  start nodemanager
starting nodemanager, logging to /home/jian/prj/bigdata/hadoop-2.10.0/logs/yarn-jian-nodemanager-laptop.out
```


查看服务：
```
[jian@laptop hadoop-2.10.0]$ jps
1537 ResourceManager
31987 NameNode
3238 DataNode
2407 Jps
2137 NodeManager
```


6.浏览器打开：[点击这里](http://localhost:8088)

![enter image description here](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/hd-yarn1.png)


7.测试：

```
[jian@laptop hadoop-2.10.0]$ hdfs dfs -rm -r /user/jian/output
Deleted /user/jian/output

[jian@laptop hadoop-2.10.0]$ hadoop jar share/hadoop/mapreduce/hadoop-mapreduce-examples-2.10.0.jar wordcount /user/jian/input /user/jian/output
```

打开页面看到里面有数据了：

![enter image description here](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/hd-yarn2.png)