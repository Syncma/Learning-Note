# MySQL高级知识-Day4

[toc]

## 慢查询

### 慢查询设置

1.默认没有启动慢查询日志，需要手动设置

如果不是调优需要的话，不建议开启该参数，

因为开启慢查询日志会带来一定的性能影响，慢查询日志支持日志记录写入文件

可以执行下面命令进行查看是否开启慢查询：
```
SHOW VARIABLES LIKE '%slow_query_log%';
```


2.开启方式

可以使用下面命令进行开启
```
set slow_query_log=1
```

或者写入mysql配置文件中
<br>

### 慢查询条件

1.什么样的sql才会被写入慢查询日志呢？

查看当前多少秒算慢？
```
SHOW VARIABLES LIKE 'long_query_time%';
```

设置慢的阙值时间 
```
 set long_query_time = 3
```

2.为什么设置后看不出变化？

需要重新连接或新开一个会话才能看到修改值。


### 慢查询检测

当MySQL繁忙的时候运行**show processlist**，会发现有很多行输出，每行输出对应一个MySQL连接。

在运营网站的过程中，可能会遇到网站突然变慢的问题，一般情况下和 MySQL 慢有关系，可以通过开启慢查询，找到影响效率的 SQL ，然后采取相应的措施。

MySQL有一个功能就是可以log下来运行的比较慢的sql语句，默认是没有这个log的，为了开启这个功能，要修改 my.cnf或者在MySQL启动的时候加入一些参数。

如果在my.cnf里面修改，需增加如下几行

>long_query_time = 1

>log-slow-queries =

>log-queries-not-using-indexes

>long_query_time 是指执行超过多久的sql会被log下来，这里是1秒。

>log-slow-queries 设置把日志写在那里，可以为空，系统会给一个缺省的文件
>log-queries-not-using-indexes 就是纪录没使用索引的sql


### 慢查询路径

可以执行下面的命令查看慢查询日志路径：

```
mysql> show variables like '%slow%';
+---------------------+---------------------------------+
| Variable_name       | Value                           |
+---------------------+---------------------------------+
| log_slow_queries    | OFF                             |
| slow_launch_time    | 2                               |
| slow_query_log      | OFF                             |
| slow_query_log_file | /var/run/mysqld/mysqld-slow.log |
+---------------------+---------------------------------+
4 rows in set (0.00 sec)
```



### 慢查询分析

阅读慢速查询日志最好是通过 **`mysqldumpslow`** 命令进行

```
mysqldumpslow –help以下，主要用的是
-s ORDER what to sort by (t, at, l, al, r, ar etc), ‘at’ is default
-t NUM just show the top n queries
-g PATTERN grep: only consider stmts that include this string

-s，是order的顺序，说明写的不够详细，俺用下来，包括看了代码，主要有
c,t,l,r和ac,at,al,ar，分别是按照query次数，时间，lock的时间和返回的记录数来排序，前面加了a的时倒叙
-t，是top n的意思，即为返回前面多少条的数据
-g，后边可以写一个正则匹配模式，大小写不敏感的
```

例子：
```
mysqldumpslow -s c -t 20 host-slow.log
mysqldumpslow -s r -t 20 host-slow.log
```

上述命令可以看出访问次数最多的20个sql语句和返回记录集最多的20个sql。

```
mysqldumpslow -t 10 -s t -g “left join” host-slow.log
```

这个是按照时间返回前10条里面含有左连接的sql语句。

```
Time: 060908 22:17:43
 Query_time: 12 Lock_time: 0 Rows_sent: 86345 Rows_examined: 580963
 
Q:这个是慢查的日志,都是些什么意思?
A:查询用了12妙，返回86345行，一共查了580963行
```

 
