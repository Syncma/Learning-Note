# RDD编程
<!-- TOC -->

- [RDD编程](#rdd%e7%bc%96%e7%a8%8b)
  - [RDD创建](#rdd%e5%88%9b%e5%bb%ba)
  - [键值对RDD](#%e9%94%ae%e5%80%bc%e5%af%b9rdd)

<!-- /TOC -->

## RDD创建

Spark Core

RDD - 提供抽象的数据结构（弹性分布式数据集）

1、从文件系统中加载数据
Spark的SparkContext通过textFile()读取数据生成内存中的RDD


本地文件:
```
from pyspark import SparkConf, SparkContext
conf = SparkConf().setMaster("local").setAppName("My App")
sc = SparkContext(conf=conf)
logFile = "file:///home/jian/prj/bigdata/spark/README.md"
logData = sc.textFile(logFile, 2).cache()
numAs = logData.filter(lambda line: 'a' in line).count()
numBs = logData.filter(lambda line: 'b' in line).count()
print("Lines with a:%s, lines with b:%s" % (numAs, numBs))
```


[jian@laptop spark]$ bin/pyspark
```
Python 3.6.7 (default, Mar 21 2019, 20:23:57)
[GCC 8.3.1 20190223 (Red Hat 8.3.1-2)] on linux
Type "help", "copyright", "credits" or "license" for more information.
19/12/04 12:46:39 WARN NativeCodeLoader: Unable to load native-hadoop library for your platform... using builtin-java classes where applicable
Setting default log level to "WARN".
To adjust logging level use sc.setLogLevel(newLevel). For SparkR, use setLogLevel(newLevel).
Welcome to
      ____              __
     / __/__  ___ _____/ /__
    _\ \/ _ \/ _ `/ __/  '_/
   /__ / .__/\_,_/_/ /_/\_\   version 2.4.4
      /_/
Using Python version 3.6.7 (default, Mar 21 2019 20:23:57)
SparkSession available as 'spark'.
>>> lines = sc.textFile("file:////home/jian/prj/bigdata/spark/README.md")
>>> lines.foreach(print)
```

HDFS文件：
```
>>>lines = sc.textFile("hdfs://localhost:9000/user/hadoop/word.txt")
>>>lines = sc.textFile("/user/hadoop/word.txt")
>>>lines = sc.textFile("word.txt")
```
这三条语句是完全等价的，可以使用任何一种方式


2、通过并行集合（数组）

```
>>> array = [1,2,3,4,5]
>>> rdd = sc.parallelize(array)
>>> rdd.foreach(print)
3
2
5
1
4
```


## 键值对RDD

1. 从文件中加载

```
[jian@laptop spark]$ cat /tmp/word.txt
Hadoop is good
Spark is fast
Spark is better
>>> lines = sc.textFile("file:///tmp/word.txt")
>>> pairRDD = lines.flatMap(lambda line:line.split(" ")).map(lambda word:(word,1))
>>> pairRDD.foreach(print)
('Spark', 1)
('is', 1)
('better', 1)
('Hadoop', 1)
('is', 1)
('good', 1)
('Spark', 1)
('is', 1)
('fast', 1)
```

2.通过并行集合创建RDD
```
>>> list = ["hadoop", "Spark", "Hive", "Spark"]
>>> rdd = sc.parallelize(list)
>>> pairRDD = rdd.map(lambda word:(word,1))
>>> pairRDD.foreach(print)
('Hive', 1)
('Spark', 1)
('hadoop', 1)
('Spark', 1)
```
