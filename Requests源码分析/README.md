# Requests源码分析

最近python学习到了瓶颈了，这次准备从Kenneth Reitz大神的requests入手

分析源码，看大神的代码是一种学习的好方法，让我从中学到很多以前不知道的知识

requests源码 [git地址](https://github.com/psf/requests)



## 01 <i class="icon-list"></i> 目录
|章节|标题|进度
|:-:|:-:|:-:|
|   Day1  | [get方法入口](https://github.com/Syncma/Learning-note/blob/master/Requests%E6%BA%90%E7%A0%81%E5%88%86%E6%9E%90/Requests%20%E6%BA%90%E7%A0%81%E9%98%85%E8%AF%BB-Day1.md)|`完成`
|   Day2  | [Session](https://github.com/Syncma/Learning-note/blob/master/Requests%E6%BA%90%E7%A0%81%E5%88%86%E6%9E%90/Requests%20%E6%BA%90%E7%A0%81%E9%98%85%E8%AF%BB-Day2.md)|`完成`
|   Day3 | [hook](https://github.com/Syncma/Learning-note/blob/master/Requests%E6%BA%90%E7%A0%81%E5%88%86%E6%9E%90/Requests%20%E6%BA%90%E7%A0%81%E9%98%85%E8%AF%BB-Day3.md)|`完成`
|   Day4  | [adapters](https://github.com/Syncma/Learning-note/blob/master/Requests%E6%BA%90%E7%A0%81%E5%88%86%E6%9E%90/Requests%20%E6%BA%90%E7%A0%81%E9%98%85%E8%AF%BB-Day4.md)|`完成`
|   Day5  | [Request](https://github.com/Syncma/Learning-note/blob/master/Requests%E6%BA%90%E7%A0%81%E5%88%86%E6%9E%90/Requests%20%E6%BA%90%E7%A0%81%E9%98%85%E8%AF%BB-Day5.md)|`完成`
|   pytest学习笔记  | [pytest](https://github.com/Syncma/Learning-note/blob/master/Requests%E6%BA%90%E7%A0%81%E5%88%86%E6%9E%90/pytest%20%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.md)|`完成`