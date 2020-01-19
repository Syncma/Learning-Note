# Hive管理

<!-- TOC -->

- [Hive管理](#hive管理)
    - [Hive的启动方式](#hive的启动方式)
        - [命令行方式](#命令行方式)
        - [Web界面方式](#web界面方式)
        - [远程服务启动方式](#远程服务启动方式)

<!-- /TOC -->

## Hive的启动方式

### 命令行方式
1.直接输入<Hive_HOME>/bin/hive的执行程序
2.或者输入hive --service cli

```
常用的命令：
1.清屏
Ctrl +L 或者! clear

2.查看数据仓库中的表
show tables

3.查看数据仓库中的内置函数
show functions

4.查看表结构
desc 表名

5.查看HDFS上的文件
dfs -ls目录

6.执行操作系统的命令
!命令

7.执行HQL语句
select xxx from xxx

8.执行SQL脚本
source xxx

9.静默模式  -不产生MapReduce 调试过程，直接产生结果
hive  -S 

10.交互模式
hive -e 'show tables';
```


### Web界面方式


安装步骤：


Hive从2.0版本开始，为HiveServer2提供了一个简单的WEB UI界面，界面中可以直观的看到当前链接的会话、历史日志、配置参数以及度量信息

服务开启：
```
[jian@laptop conf]$ hive --service hiveserver2 &
```
默认端口是10002

浏览器访问：[访问这里](http://localhost:10002)


### 远程服务启动方式

端口号10000
启动方式: hive --service  hiveserver &

备注:
以JDBC或者ODBC的程序登陆到hive中操作数据时， 必须使用远程服务启动方式