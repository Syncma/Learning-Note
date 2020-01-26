# RDD持久化

<!-- TOC -->

- [RDD持久化](#rdd%e6%8c%81%e4%b9%85%e5%8c%96)
  - [持久化](#%e6%8c%81%e4%b9%85%e5%8c%96)

<!-- /TOC -->

## 持久化
对于迭代计算，经常需要多次重复使用同一组数据

可以通过**持久化（缓存）机制**来避免这种重复计算的开销

可以使用**persist()方法**对一个RDD标记为持久化

之所以说”标记为持久化“， 是因为出现persist()语句的地方，
并不会马上计算生成RDD并把它持久化，而是要等到遇到第一个**行动操作**
触发真正计算以后，才会吧计算结果进行持久化

持久化后的RDD将会被保留在计算节点的内存中被后面的行动操作重复使用

例子：
```
>>> list = ["hadoop", "Spark", "Hive"]
>>> rdd = sc.parallelize(list)
>>> print(rdd.count())   //行动操作，触发一次真正从头到尾的计算
3
>>> print(','.join(rdd.collect()))  //行动操作，触发一次真正从头到尾的计算
hadoop,Spark,Hive
```


```

.persist(MEMORY_AND_DISK)  #内存不足，存放磁盘
.persist(MEMORY_ONLY) 存内存

.cache()   等价于.persist(MEMORY_ONLY)
.unpersist()方法： 手动吧持久化的RDD从缓存中移除

```
```
>>> list = ["hadoop", "Spark", "Hive"]
>>> rdd = sc.parallelize(list)
>>> rdd.cache()
#会调用.persist(MEMORY_ONLY), 但是语句执行到这里并不会缓存rrd
因为这时候rrd并没有被计算生成
ParallelCollectionRDD[38] at parallelize at PythonRDD.scala:195
>>> print(rdd.count())
#第一次行动操作，触发一次真正从头到尾的计算，这时上面的rrd.cache()才会被执行
把这个rdd放到缓存中
3
>>> print(','.join(rdd.collect()))
#第二次行动操作，不需要触发从头到尾的计算，只需要重复使用上面缓存中的rdd
hadoop,Spark,Hive
```
