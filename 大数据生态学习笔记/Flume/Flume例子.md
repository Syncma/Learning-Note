# Flume例子


<!-- TOC -->

- [Flume例子](#flume%e4%be%8b%e5%ad%90)
	- [flume简单理解](#flume%e7%ae%80%e5%8d%95%e7%90%86%e8%a7%a3)
	- [例子](#%e4%be%8b%e5%ad%90)

<!-- /TOC -->
## flume简单理解
可以简单理解成：

```
这是一个关于池子的故事。有一个池子，它一头进水，另一头出水，进水口可以配置各种管子，出水口也可以配置各种管子，可以有多个进水口、多个出水口。

水术语称为Event，进水口术语称为Source、出水口术语成为Sink、池子术语成为Channel，

Source+Channel+Sink，术语称为Agent。如果有需要，还可以把多个Agent连起来。

```


## 例子

1.配置服务端文件

```
[jian@laptop flume]$ mkdir job
[jian@laptop flume]$ cd job
[jian@laptop job]$ cat thrift-flume-logger.conf
# Name the components on this agent
a1.sources = r1
a1.sinks = s1
a1.channels = c1

# Describe/configure the source
#a1.sources.r1.type = netcat
a1.sources.r1.type = thrift
a1.sources.r1.channels = c1
a1.sources.r1.bind = localhost
a1.sources.r1.port = 4141

# Describe the sink
a1.sinks.s1.channel = c1
a1.sinks.s1.type = logger

# Use a channel which buffers events in memory
a1.channels.c1.type = memory
a1.channels.c1.capacity = 1000
a1.channels.c1.transactionCapacity = 100

# Bind the source and sink to the channel
a1.sources.r1.channels = c1
a1.sinks.s1.channel = c1
```
2.启动flume服务

```
[jian@laptop flume]$ bin/flume-ng agent -c conf/ -f job/thrift-flume-logger.conf --name a1 -Dflume.root.logger=INFO,console
xxx
2019-12-08 21:34:25,131 (lifecycleSupervisor-1-1) [INFO - org.apache.flume.source.ThriftSource.start(ThriftSource.java:206)] Started Thrift source.
```
3.测试，另外开一个终端
```
[jian@laptop flume]$ nc localhost 4141
hello  #输入hello
查看服务端日志：
2019-12-08 21:36:08,359 (Thread-1) [ERROR - org.apache.thrift.server.TThreadedSelectorServer$SelectorThread.run(TThreadedSelectorServer.java:549)] run() on SelectorThread exiting due to uncaught error
java.lang.OutOfMemoryError: Java heap space
at java.nio.HeapByteBuffer.<init>(HeapByteBuffer.java:57)
at java.nio.ByteBuffer.allocate(ByteBuffer.java:335)
at org.apache.thrift.server.AbstractNonblockingServer$FrameBuffer.read(AbstractNonblockingServer.java:371)
at org.apache.thrift.server.AbstractNonblockingServer$AbstractSelectThread.handleRead(AbstractNonblockingServer.java:203)
at org.apache.thrift.server.TThreadedSelectorServer$SelectorThread.select(TThreadedSelectorServer.java:586)
at org.apache.thrift.server.TThreadedSelectorServer$SelectorThread.run(TThreadedSelectorServer.java:541)
```
解决方式：

修改 flume下的conf/flume-env.sh文件：
```
export JAVA_OPTS="-Xms512m -Xmx1024m -Dcom.sun.management.jmxremote" 
其中：
-Xms<size> set initial Java heap size
-Xmx<size> set maximum Java heap size
 
主要修改Xmx和Xms两个参数
```

4.安装 thrift
```
[root@laptop ~]# dnf install thrift
```

5.编译flume.thrift
```
解压apache-flume-1.9.0-src.tar.gz找到flume.thrift

[jian@laptop tmp]$ tar xf apache-flume-1.9.0-src.tar.gz
[jian@laptop thrift]$ pwd
/tmp/apache-flume-1.9.0-src/flume-ng-sdk/src/main/thrift
[jian@laptop thrift]$ thrift --gen py flume.thrift 


生成一个gen-py目录重命名为flumepy，把flumepy复制到python的安装目录：xxx/lib/python2.7/site-packages下（手工安装包）

[jian@laptop thrift]$ mv gen-py/ flumepy
[jian@laptop thrift]$ mv flumepy/ ~/.pyenv/versions/3.6.7/lib/python3.6/site-packages/
```
6.安装python的thrift包
```
[jian@laptop site-packages]$ pip install thrift
```

7. python测试代码 demo.py
```
from flumepy.flume import ThriftSourceProtocol
from flumepy.flume.ttypes import ThriftFlumeEvent
from thrift.transport import TTransport, TSocket
from thrift.protocol import TCompactProtocol

#生命一个flume的客户端对象，封装发送数据的功能
class FlumeClient(object):
	def __init__(self,
		thrift_host,
		thrift_port,
		timeout=None,
		unix_socket=None):

		self.timeout = timeout
		self._socket = TSocket.TSocket(thrift_host, thrift_port, unix_socket)
		self._transport_factory = TTransport.TFramedTransportFactory()
		self._transport = self._transport_factory.getTransport(self._socket)

		self._protocol = TCompactProtocol.TCompactProtocol(
		trans=self._transport)
		self.client = ThriftSourceProtocol.Client(iprot=self._protocol,
		oprot=self._protocol)
		self.connect()

	#建立thrift连接
	def connect(self):
	try:
		if self.timeout:
			self._socket.setTimeout(self.timeout)
			if not self.is_open():
				self._transport = self._transport_factory.getTransport(
				self._socket)
				self._transport.open()
	except Exception as e:
		print(e)
		self.close()

#判断当前连接是否打开
def is_open(self):
	return self._transport.isOpen()

#发送数据
def send(self, event):
	try:
		self.client.append(event)
	except Exception as e:
		print(e)
	finally:
		self.connect()

#批量发送数据
def send_batch(self, events):
	try:
		self.client.appendBatch(events)
	except Exception as e:
		print(e)
	finally:
		self.connect()

#关闭连接
def close(self):
	self._transport.close()


#发送测试代码
if __name__ == '__main__':
	#建立连接
	flume_client = FlumeClient('127.0.0.1', 4141)
	#生成一个event ， 并发送
	#ThriftFlumeEvent(header=none,body=none)
	#header是一个字典
	#body是一个字符串
	event = ThriftFlumeEvent({
	'k1': 'v1',
	'k2': 'v2'
	}, '这是一个event的body部分'.encode())
	flume_client.send(event)

	#生成10个event一次行发送
	events = [
	ThriftFlumeEvent({
	'k1': '1',
	'k2': '2'
	}, ('body部分%s' % _).encode()) for _ in range(10)
	]
	flume_client.send_batch(events)
	flume_client.close()
```


运行后查看日志显示：

```
2019-12-08 22:21:50,486 (SinkRunner-PollingRunner-DefaultSinkProcessor) [INFO - org.apache.flume.sink.LoggerSink.process(LoggerSink.java:95)] Event: { headers:{k1=v1, k2=v2} body: E8 BF 99 E6 98 AF E4 B8 80 E4 B8 AA 65 76 65 6E ............even }
2019-12-08 22:21:50,498 (SinkRunner-PollingRunner-DefaultSinkProcessor) [INFO - org.apache.flume.sink.LoggerSink.process(LoggerSink.java:95)] Event: { headers:{k1=1, k2=2} body: 62 6F 64 79 E9 83 A8 E5 88 86 30 body......0 }
2019-12-08 22:21:50,498 (SinkRunner-PollingRunner-DefaultSinkProcessor) [INFO - org.apache.flume.sink.LoggerSink.process(LoggerSink.java:95)] Event: { headers:{k1=1, k2=2} body: 62 6F 64 79 E9 83 A8 E5 88 86 31 body......1 }
2019-12-08 22:21:50,499 (SinkRunner-PollingRunner-DefaultSinkProcessor) [INFO - org.apache.flume.sink.LoggerSink.process(LoggerSink.java:95)] Event: { headers:{k1=1, k2=2} body: 62 6F 64 79 E9 83 A8 E5 88 86 32 body......2 }
2019-12-08 22:21:50,499 (SinkRunner-PollingRunner-DefaultSinkProcessor) [INFO - org.apache.flume.sink.LoggerSink.process(LoggerSink.java:95)] Event: { headers:{k1=1, k2=2} body: 62 6F 64 79 E9 83 A8 E5 88 86 33 body......3 }
2019-12-08 22:21:50,499 (SinkRunner-PollingRunner-DefaultSinkProcessor) [INFO - org.apache.flume.sink.LoggerSink.process(LoggerSink.java:95)] Event: { headers:{k1=1, k2=2} body: 62 6F 64 79 E9 83 A8 E5 88 86 34 body......4 }
2019-12-08 22:21:50,499 (SinkRunner-PollingRunner-DefaultSinkProcessor) [INFO - org.apache.flume.sink.LoggerSink.process(LoggerSink.java:95)] Event: { headers:{k1=1, k2=2} body: 62 6F 64 79 E9 83 A8 E5 88 86 35 body......5 }
2019-12-08 22:21:50,500 (SinkRunner-PollingRunner-DefaultSinkProcessor) [INFO - org.apache.flume.sink.LoggerSink.process(LoggerSink.java:95)] Event: { headers:{k1=1, k2=2} body: 62 6F 64 79 E9 83 A8 E5 88 86 36 body......6 }
2019-12-08 22:21:50,500 (SinkRunner-PollingRunner-DefaultSinkProcessor) [INFO - org.apache.flume.sink.LoggerSink.process(LoggerSink.java:95)] Event: { headers:{k1=1, k2=2} body: 62 6F 64 79 E9 83 A8 E5 88 86 37 body......7 }
2019-12-08 22:21:50,500 (SinkRunner-PollingRunner-DefaultSinkProcessor) [INFO - org.apache.flume.sink.LoggerSink.process(LoggerSink.java:95)] Event: { headers:{k1=1, k2=2} body: 62 6F 64 79 E9 83 A8 E5 88 86 38 body......8 }
2019-12-08 22:21:50,500 (SinkRunner-PollingRunner-DefaultSinkProcessor) [INFO - org.apache.flume.sink.LoggerSink.process(LoggerSink.java:95)] Event: { headers:{k1=1, k2=2} body: 62 6F 64 79 E9 83 A8 E5 88 86 39 body......9 }
```














