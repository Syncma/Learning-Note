# python例子
<!-- TOC -->

- [python例子](#python%e4%be%8b%e5%ad%90)
  - [安装](#%e5%ae%89%e8%a3%85)
  - [例子](#%e4%be%8b%e5%ad%90)

<!-- /TOC -->
## 安装
```python
pip install kafka-python
```

[官网地址](https://github.com/dpkp/kafka-python)



## 例子
producer.py:

```python
from time import sleep
from json import dumps
from kafka import KafkaProducer

producer = KafkaProducer(bootstrap_servers=['localhost:9092'],
                         value_serializer=lambda x: dumps(x).encode('utf-8'))

for e in range(100):
    data = {'number': e}
    producer.send('testTopic', value=data)
    print('Send Message is {} '.format(data))
    sleep(5)
```


consumer.py:
```python
from json import loads
from kafka import KafkaConsumer


consumer = KafkaConsumer('testTopic',
bootstrap_servers=['localhost:9092'],
auto_offset_reset='earliest',
enable_auto_commit=True,
group_id='my-group',
value_deserializer=lambda x: loads(x))

for message in consumer:
     message = message.value
     print('Recieve Message is {} '.format(message))
```

