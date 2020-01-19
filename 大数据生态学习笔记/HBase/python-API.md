# python-API

<!-- TOC -->

- [python-API](#python-api)
  - [开启thrift服务](#%e5%bc%80%e5%90%afthrift%e6%9c%8d%e5%8a%a1)
  - [python测试](#python%e6%b5%8b%e8%af%95)

<!-- /TOC -->


## 开启thrift服务
HBase通过thrift机制 -（RPC 框架）可以实现多语言编程，信息通过端口传递


开启thrift服务：
```
[jian@laptop bin]$ hbase-daemon.sh start thrift --infoport 9095 -p 9090

--infoport 是web页面访问端口
-p: 是API访问的接口

```

## python测试

1.安装模块
```
pip install happybase
```

2.例子：
```
import happybase
connection = happybase.Connection('localhost')
table = connection.table('student')
table.put(b'stu003', {b'info:name': b'jacky', b'info:age': b'50'})
row = table.row(b'stu003')
print(row[b'info:name'])  # prints 'value1'
for key, data in table.rows([b'stu001', b'stu003']):
    print(key, data)  # prints row key and data for each row
for key, data in table.scan(row_prefix=b'stu'):
    print(key, data)  # prints 'value1' and 'value2'
# row = table.delete(b'row-key')
```

[官网文档](https://happybase.readthedocs.io/en/latest/user.html)


