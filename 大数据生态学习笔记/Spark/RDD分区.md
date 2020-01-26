# RDD分区
<!-- TOC -->

- [RDD分区](#rdd分区)
    - [分区作用](#分区作用)
    - [分区原则](#分区原则)
    - [自定义分区方法](#自定义分区方法)

<!-- /TOC -->

## 分区作用
* 增加并行度
![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/20200126102104.png)


* 减少通信开销
![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/20200126102212.png)
![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/20200126102241.png)


## 分区原则

**`分区个数=集群中CPU核数`**

可以进行设置参数 来设置默认的分区数目
**spark.default.parallelism**

>local模式  -默认为本地机器的CPU数目
>
>Apache mesos模式 -默认分区数目为8
>
>对于standalone和yarn模式  - 集群中所有CPU核心数目总和  和2 之间 取较大值

也可以手动指定分区
```
sc.textFile(path, partionNum)
```

```
>>> list = [1,2,3,4,5]
>>> rdd = sc.parallelize(list,2)  //设置2个分区
>>>
```

也可以使用repartition方法重新设置分区个数

```
>>> data = sc.parallelize([1,2,3,4,5],2)
>>> len(data.glom().collect())  #显示data这个RDD的分区数量
2
>>> rdd = data.repartition(1) #对data这个RDD重新分区
>>> len(rdd.glom().collect()) #显示rdd这个RDD的分区数量
1
```

## 自定义分区方法

Count.py 

```python
from pyspark import SparkConf, SparkContext


def MyPartitioner(key):
    print("MyPartitioner is running")
    print("The key is %d" % key)
    return key % 10


def main():
    print("The main function is running")
    conf = SparkConf().setMaster("local").setAppName("My App")
    sc = SparkContext(conf=conf)
    data = sc.parallelize(range(10), 5)
    data.map(lambda x:(x,1)) \
        .partitionBy(10,MyPartitioner) \
        .map(lambda x:x[0]) \
        .saveAsTextFile("file:///tmp/partitioner")


if __name__ == "__main__":
    main()
```

[jian@laptop python]$ python Count.py
```
The main function is running
19/12/04 15:04:17 WARN NativeCodeLoader: Unable to load native-hadoop library for your platform... using builtin-java classes where applicable
Using Spark's default log4j profile: org/apache/spark/log4j-defaults.properties
Setting default log level to "WARN".
To adjust logging level use sc.setLogLevel(newLevel). For SparkR, use setLogLevel(newLevel).
19/12/04 15:04:18 WARN Utils: Service 'SparkUI' could not bind on port 4040. Attempting port 4041.
[Stage 0:>                                                          (0 + 1) / 5]MyPartitioner is running
The key is 0
MyPartitioner is running
The key is 1
MyPartitioner is running
The key is 2
MyPartitioner is running
The key is 3
MyPartitioner is running
The key is 4
MyPartitioner is running
The key is 5
[Stage 0:=======================>                                   (2 + 1) / 5]MyPartitioner is running
The key is 6
MyPartitioner is running
The key is 7
MyPartitioner is running
The key is 8
MyPartitioner is running
The key is 9



[jian@laptop partitioner]$ ll
total 40
-rw-r--r-- 1 jian jian 2 Dec  4 15:04 part-00000
-rw-r--r-- 1 jian jian 2 Dec  4 15:04 part-00001
-rw-r--r-- 1 jian jian 2 Dec  4 15:04 part-00002
-rw-r--r-- 1 jian jian 2 Dec  4 15:04 part-00003
-rw-r--r-- 1 jian jian 2 Dec  4 15:04 part-00004
-rw-r--r-- 1 jian jian 2 Dec  4 15:04 part-00005
-rw-r--r-- 1 jian jian 2 Dec  4 15:04 part-00006
-rw-r--r-- 1 jian jian 2 Dec  4 15:04 part-00007
-rw-r--r-- 1 jian jian 2 Dec  4 15:04 part-00008
-rw-r--r-- 1 jian jian 2 Dec  4 15:04 part-00009
-rw-r--r-- 1 jian jian 0 Dec  4 15:04 _SUCCESS
[jian@laptop partitioner]$ pwd
/tmp/partitioner
```

