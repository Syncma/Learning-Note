# 内部topic _consumer_offsets详解


<!-- TOC -->

- [内部topic _consumer_offsets详解](#%e5%86%85%e9%83%a8topic-consumeroffsets%e8%af%a6%e8%a7%a3)
	- [测试](#%e6%b5%8b%e8%af%95)

<!-- /TOC -->
## 测试

新版Kafka已推荐将consumer的位移信息保存在Kafka内部的topic中，
即**`__consumer_offsets topic`**

并且默认提供了kafka_consumer_groups.sh脚本供用户查看consumer信息。


1.创建topic:
```
[jian@laptop kafka]# bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic testTopic
Created topic testTopic.
```

2.使用脚本生产消息
```
[jian@laptop bin]$ ./kafka-console-producer.sh --broker-list localhost:9092 --topic testTopic
>hello
>world
```
3.验证生产消息成功
```
[jian@laptop bin]$ ./kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list localhost:9092 --topic testTopic --time -1
testTopic:0:2

参数： --time  -1 表示从最新的时间的offset得到数据条数
输出结果表示topic:partition:untilOffset
```

4.创建consumer group
```
[jian@laptop bin]$ ./kafka-console-consumer.sh --bootstrap-server localhost:9092 --consumer.config ../config/consumer.properties --topic testTopic --from-beginning
hello
world
```

5.获取group id
```
[jian@laptop bin]$ ./kafka-consumer-groups.sh --bootstrap-server localhost:9092 --list
test-consumer-group

其实在第4步中在配置文件consumer.properties中配置了group.id
如果没有配置，会随机产生一个test-consumer-group的group.id
```

6.查询consumer offsets topic内容
```
jian@laptop kafka]$ bin/kafka-console-consumer.sh --topic __consumer_offsets --bootstrap-server localhost:9092 --formatter "kafka.coordinator.group.GroupMetadataManager\$OffsetsMessageFormatter" --consumer.config config/consumer.properties --from-beginning

[test-consumer-group,testTopic,0]::OffsetAndMetadata(offset=2, leaderEpoch=Optional[0], metadata=, commitTimestamp=1575528207199, expireTimestamp=None)
[test-consumer-group,testTopic,0]::OffsetAndMetadata(offset=2, leaderEpoch=Optional[0], metadata=, commitTimestamp=1575528212189, expireTimestamp=None)
[test-consumer-group,testTopic,0]::OffsetAndMetadata(offset=2, leaderEpoch=Optional[0], metadata=, commitTimestamp=1575528217195, expireTimestamp=None)
[test-consumer-group,testTopic,0]::OffsetAndMetadata(offset=2, leaderEpoch=Optional[0], metadata=, commitTimestamp=1575528222192, expireTimestamp=None)
[test-consumer-group,testTopic,0]::OffsetAndMetadata(offset=2, leaderEpoch=Optional[0], metadata=, commitTimestamp=1575528227192, expireTimestamp=None)
[test-consumer-group,testTopic,0]::OffsetAndMetadata(offset=2, leaderEpoch=Optional[0], metadata=, commitTimestamp=1575528232193, expireTimestamp=None)
[test-consumer-group,testTopic,0]::OffsetAndMetadata(offset=2, leaderEpoch=Optional[0], metadata=, commitTimestamp=1575528237193, expireTimestamp=None)
[test-consumer-group,testTopic,0]::OffsetAndMetadata(offset=2, leaderEpoch=Optional[0], metadata=, commitTimestamp=1575528242194, expireTimestamp=None)

**默认情况下__consumer_offsets有50个分区**
```

7. 计算指定consumer group在__consumer_offsets topic中分区信息

这时候就用到了第5步获取的group.id(本例中是console-consumer-46965)。
Kafka会使用下面公式计算该group位移保存在__consumer_offsets的哪个分区上：

```java
[jian@laptop tmp]$ cat Test.java

import java.lang.Math;
public class Test {
	public static void main(String args[]){
		int hashCode = Math.abs("test-consumer-group".hashCode());
		int partition = hashCode % 50;
		System.out.println(partition);
    }
}
```


```
[jian@laptop tmp]$ javac Test.java
[jian@laptop tmp]$ java Test
31
```

对应的分区=31，即__consumer_offsets的分区31保存了这个consumer group的位移信息

下面让我们验证一下。


8.获取指定consumer group的位移信息

```
[jian@laptop bin]$ ./kafka-console-consumer.sh --topic __consumer_offsets --partition 31 --bootstrap-server localhost:9092 --formatter "kafka.coordinator.group.GroupMetadataManager\$OffsetsMessageFormatter"

[test-consumer-group,testTopic,0]::OffsetAndMetadata(offset=2, leaderEpoch=Optional[0], metadata=, commitTimestamp=1575531479387, expireTimestamp=None)
[test-consumer-group,testTopic,0]::OffsetAndMetadata(offset=2, leaderEpoch=Optional[0], metadata=, commitTimestamp=1575531484388, expireTimestamp=None)
[test-consumer-group,testTopic,0]::OffsetAndMetadata(offset=2, leaderEpoch=Optional[0], metadata=, commitTimestamp=1575531489388, expireTimestamp=None)
[test-consumer-group,testTopic,0]::OffsetAndMetadata(offset=2, leaderEpoch=Optional[0], metadata=, commitTimestamp=1575531494389, expireTimestamp=None)
[test-consumer-group,testTopic,0]::OffsetAndMetadata(offset=2, leaderEpoch=Optional[0], metadata=, commitTimestamp=1575531499390, expireTimestamp=None)
[test-consumer-group,testTopic,0]::OffsetAndMetadata(offset=2, leaderEpoch=Optional[0], metadata=, commitTimestamp=1575531504390, expireTimestamp=None)
[test-consumer-group,testTopic,0]::OffsetAndMetadata(offset=2, leaderEpoch=Optional[0], metadata=, commitTimestamp=1575531509390, expireTimestamp=None)
[test-consumer-group,testTopic,0]::OffsetAndMetadata(offset=2, leaderEpoch=Optional[0], metadata=, commitTimestamp=1575531514392, expireTimestamp=None)
[test-consumer-group,testTopic,0]::OffsetAndMetadata(offset=2, leaderEpoch=Optional[0], metadata=, commitTimestamp=1575531519393, expireTimestamp=None)
[test-consumer-group,testTopic,0]::OffsetAndMetadata(offset=2, leaderEpoch=Optional[0], metadata=, commitTimestamp=1575531524393, expireTimestamp=None)


####又输入了三条消息发现offset 变成 5了####
[test-consumer-group,testTopic,0]::OffsetAndMetadata(offset=5, leaderEpoch=Optional[0], metadata=,    commitTimestamp=1575531529394, expireTimestamp=None)

```

该consumer group果然保存在分区11上，且位移信息都是对的(这里的位移信息是已消费的位移）
也就是说没有消费者消费信息，offset不会变化



