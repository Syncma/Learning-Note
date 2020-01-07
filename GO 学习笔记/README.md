# GO 学习笔记

一边学go gin框架，一边复习go语法知识


## 01 <i class="icon-list"></i> 目录
|章节|标题|进度
|:-:|:-:|:-:|
|   第1章  | [go概述](https://github.com/Syncma/Learning-note/blob/master/GO%20%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/Go%20%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%20-%E7%AC%AC1%E7%AB%A0.md)|`完成`
|   第2章  | [程序结构](https://github.com/Syncma/Learning-note/blob/master/GO%20%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/Go%20%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%20-%E7%AC%AC2%E7%AB%A0.md)|`完成`
|   第3章  | [基本数据类型](https://github.com/Syncma/Learning-note/blob/master/GO%20%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/Go%20%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%20-%E7%AC%AC3%E7%AB%A0.md)|`完成`
|   第4章  |[复合数据类型](https://github.com/Syncma/Learning-note/blob/master/GO%20%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/Go%20%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%20-%E7%AC%AC4%E7%AB%A0.md)|`完成`
|   第5章  |[函数](https://github.com/Syncma/Learning-note/blob/master/GO%20%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/Go%20%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%20-%E7%AC%AC5%E7%AB%A0.md) |`完成`
|   第6章  |[方法](https://github.com/Syncma/Learning-note/blob/master/GO%20%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/Go%20%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%20-%E7%AC%AC6%E7%AB%A0.md)|`完成`
|   第7章  |[接口](https://github.com/Syncma/Learning-note/blob/master/GO%20%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/Go%20%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%20-%E7%AC%AC7%E7%AB%A0.md)|`完成`
|   第8章  |[goroutine和通道](https://github.com/Syncma/Learning-note/blob/master/GO%20%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/Go%20%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%20-%E7%AC%AC8%E7%AB%A0.md)|`待补充`
|   第9章  |[使用共享变量实现并发](https://github.com/Syncma/Learning-note/blob/master/GO%20%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/Go%20%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%20-%E7%AC%AC9%E7%AB%A0.md)|`完成`
|   第10章  |[包和go工具](https://github.com/Syncma/Learning-note/blob/master/GO%20%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/Go%20%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%20-%E7%AC%AC10%E7%AB%A0.md)|`完成`
|   第11章  |[测试](https://github.com/Syncma/Learning-note/blob/master/GO%20%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/Go%20%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%20-%E7%AC%AC11%E7%AB%A0.md)|`完成`
|   第12章  |[反射](https://github.com/Syncma/Learning-note/blob/master/GO%20%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/Go%20%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%20-%E7%AC%AC12%E7%AB%A0.md)|`完成`
|   第13章  |[指针](https://github.com/Syncma/Learning-note/blob/master/GO%20%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/Go%20%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%20-%E7%AC%AC13%E7%AB%A0.md)|`完成`
|   第14章  |[循环](https://github.com/Syncma/Learning-note/blob/master/GO%20%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/Go%20%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%20-%E7%AC%AC14%E7%AB%A0.md)|`完成`
|   第15章  |[defer](https://github.com/Syncma/Learning-note/blob/master/GO%20%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/Go%20%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%20-%E7%AC%AC15%E7%AB%A0.md)|`完成`
|   第16章  |[结构体](https://github.com/Syncma/Learning-note/blob/master/GO%20%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/Go%20%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%20-%E7%AC%AC16%E7%AB%A0.md)|`完成`
|   第17章  |[go中的main函数和init函数](https://github.com/Syncma/Learning-note/blob/master/GO%20%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/Go%20%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%20-%E7%AC%AC17%E7%AB%A0.md)|`完成`
|   第18章  |[堆空间&栈空间](https://github.com/Syncma/Learning-note/blob/master/GO%20%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0/Go%20%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0%20-%E7%AC%AC18%E7%AB%A0.md)|`完成`


## 02 <i class="icon-desktop"></i> 参考

- [ ] 《Go程序设计语言》（艾伦 A. A. 多诺万 著 李道兵 译）
