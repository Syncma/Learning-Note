# Gunicorn介绍

为什么要用gunicron?

gunicron是什么？

它使用的协议是前文所讲的WSGI，详细的使用教程请点击这里（http://gunicorn.org/）。
Gunicorn采用prefork模式，Gunicorn 服务器与各种 Web 框架兼容，只需非常简单的执行，轻量级的资源消耗，以及相当迅速。
它的特点是与 web框架结合紧密，部署特别方便。
 缺点也很多，不支持 HTTP 1.1，并发访问性能不高，与 uWSGI，Gevent 等有一定的性能差距。   ----性能怎么检测出来的？



1. Gunicorn设计


Gunicorn 是一个 master 进程，spawn 出数个工作进程的 web 服务器。
master 进程控制工作进程的产生与消亡，工作进程只需要接受请求并且处理。
这样分离的方式使得 reload 代码非常方便，也很容易增加或减少工作进程。 
工作进程这块作者给了很大的扩展余地，它可以支持不同的IO方式，如 Gevent,Sync 同步进程，Asyc 异步进程，Eventlet 等等。
master 跟 worker 进程完全分离，使得 Gunicorn 实质上就是一个控制进程的服务。


2. Gunicorn源码结构


从 Application.run() 开始，首先初始化配置，从文件读取，终端读取等等方式完成 configurate。
然后启动 Arbiter，Arbiter 是实质上的 master 进程的核心，它首先从配置类中读取并设置，然后初始化信号处理函数，建立 socket。
然后就是开始 spawn 工作进程，根据配置的工作进程数进行 spawn。然后就进入了轮询状态，收到信号，处理信号然后继续。
这里唤醒进程的方式是建立一个 PIPE，通过信号处理函数往 pipe 里 write，然后 master 从 select.select() 中唤醒。


工作进程在 spawn 后，开始初始化，然后同样对信号进行处理，并且开始轮询，处理 HTTP 请求，调用 WSGI 的应用端，得到 resopnse 返回。然后继续。


Sync 同步进程的好处在于每个 request 都是分离的，每个 request 失败都不会影响其他 request，但这样导致了性能上的瓶颈。
