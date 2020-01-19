# Hive的数据类型



<!-- TOC -->

- [Hive的数据类型](#hive%e7%9a%84%e6%95%b0%e6%8d%ae%e7%b1%bb%e5%9e%8b)
  - [基本数据类型](#%e5%9f%ba%e6%9c%ac%e6%95%b0%e6%8d%ae%e7%b1%bb%e5%9e%8b)
  - [复杂的数据类型](#%e5%a4%8d%e6%9d%82%e7%9a%84%e6%95%b0%e6%8d%ae%e7%b1%bb%e5%9e%8b)
  - [时间类型](#%e6%97%b6%e9%97%b4%e7%b1%bb%e5%9e%8b)

<!-- /TOC -->

## 基本数据类型
- tinyint/smallint/int/bigint  整数类型
- float/double 浮点数类型
- boolean  布尔型
- string 字符串型

## 复杂的数据类型
- Array : 数组类型，由一系列相同数据类型的元素组成
- Map: 集合类型，包含key-value键值对，通过key来访问元素
- Struct:  结构类型，可以包含不同数据类型的元素。这些元素可以通过"点语法“的方式来得到所需要的元素。

## 时间类型
- Date : 从Hive 0.12.0开始支持。
- Timestamp: 从Hive 0.8.0开始支持