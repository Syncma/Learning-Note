# RDD操作

<!-- TOC -->

- [RDD操作](#rdd%e6%93%8d%e4%bd%9c)
  - [转换操作](#%e8%bd%ac%e6%8d%a2%e6%93%8d%e4%bd%9c)
  - [行动操作](#%e8%a1%8c%e5%8a%a8%e6%93%8d%e4%bd%9c)
  - [惰性机制](#%e6%83%b0%e6%80%a7%e6%9c%ba%e5%88%b6)

<!-- /TOC -->
## 转换操作 

对于RDD而言，每一次转换操作都会产生不同的RDD,供一个操作使用

转换得到的RDD是惰性求值的，也就是说整个转换过程只是记录转换的轨迹，
并不会发生真正的计算，只有遇到行动操作时，才会发生真正的计算。
从血缘关系的源头开始进行 从头到尾的计算操作


常用转换操作：
```
Filter(func)  -筛选出满足函数func的元素, 并返回一个新的数据集
>>> lines = sc.textFile("file:////home/jian/prj/bigdata/spark/README.md")
>>> lineswithspark = lines.filter(lambda line:"Spark" in line)
>>> lineswithspark.foreach(print)
```
Flatmap -与map相似，但每个输入元素都可以映射到0或者多个输出结果

拍扁(Flat)

```
[jian@laptop tmp]$ cat word.txt
Hadoop is good
Spark is fast
Spark is better
>>> lines = sc.textFile("file:///tmp/word.txt")
>>> words = lines.map(lambda line: line.split(" "))
>>> words.foreach(print)
['Spark', 'is', 'better']
['Hadoop', 'is', 'good']
['Spark', 'is', 'fast']


>>> lines = sc.textFile("file:///tmp/word.txt")
>>> words = lines.flatMap(lambda line:line.split(" "))
>>> words.foreach(print)
Hadoop
is
good
Spark
is
fast
Spark
is
better
```

GroupByKey - 应用于(K,V)键值队的数据集，返回一个新的(K, Iterable)形式的数据集

```
>>> words = sc.parallelize([("Hadoop", 1), ("is", 1), ("good", 1),\
... ("Spark", 1), ("is", 1), ("fast", 1), ("Spark", 1)])
>>> word1 = words.groupByKey()
>>> word1.foreach(print)
('Spark', <pyspark.resultiterable.ResultIterable object at 0x7f1057ebf240>)
('good', <pyspark.resultiterable.ResultIterable object at 0x7f1057ebf240>)
('is', <pyspark.resultiterable.ResultIterable object at 0x7f1057ebf240>)
('fast', <pyspark.resultiterable.ResultIterable object at 0x7f1057ebf240>)
('Hadoop', <pyspark.resultiterable.ResultIterable object at 0x7f1057ebf240>)
```



ReduceByKey -应用于(K,V)键值对的数据集时，返回一个新的(K,V)形式的数据集，
其中每个值是将每个key传递到函数func中进行聚合后的结果

```
>>> words = sc.parallelize([("Hadoop", 1), ("is", 1), ("good", 1),\
... ("Spark", 1), ("is", 1), ("fast", 1), ("Spark", 1)])
>>> word1 = words.reduceByKey(lambda a,b:a+b)
>>> word1.foreach(print)
('Spark', 2)
('is', 2)
('fast', 1)
('good', 1)
('Hadoop', 1)
```

map -将每个元素传递到函数func中，并将结果返回成一个新的数据集

![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/20200120150652.png)

```
>>> data = [1,2,3,4,5]
>>> rdd1 = sc.parallelize(data)
>>> rdd2 = rdd1.map(lambda x: x+10)
>>> rdd2.foreach(print)
11
13
15
12
14
```


## 行动操作

count() -返回数据集中的元素个数

collect() -以数组的形式返回数据集中的所有元素

first() -返回数据集中的第一个元素

take(n) 以数组的形式返回数据集中的前n个元素

reduce(func) -通过函数func(输入两个参数并返回一个值）聚合数据集中的元素

foreach(func) -将数据集中的每个元素传递到函数func中运行


```
>>> rdd = sc.parallelize([1,2,3,4,5])
>>> rdd.count()
5
>>> rdd.first()
1
>>> rdd.take(3)
[1, 2, 3]
>>> rdd.reduce(lambda a,b:a+b)
15
>>> rdd.foreach(lambda elem:print(elem))
2
4
3
5
1
```


## 惰性机制
![](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/20200120150727.png)

```
>>> lines = sc.textFile("file:///tmp/word.txt")
>>> lineLengths = lines.map(lambda s:len(s))
>>> totalLength = lineLengths.reduce(lambda a,b:a+b)
>>> print(totalLength)
42
```

