# go 如何用连接池＆协程池实现高并发＆qps项目的研究与实现


## 参考１ [每分钟处理一百万请求 Handling 1 Million Requests per Minute with Go](http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/)


* 基本概念 
    
   其实最主要的还是一个链接池的概念，相信用过sync.pool这个golang内置的库或者有研究的话，理解起来就很方便了。
   
   
   * payload 工作任务对象
   
       ```
       具体的任务对象，可以有不同种类的工作任务函数
       ```

   * worker 工作者对象
       
       ```
       包含：启动、停止、创建等动作
       
       Start: Start方法启动worker的运行循环，监听退出通道以防必要的时候我们需要停止它
       
       WorkerPool chan chan Job		//工作队列
       JobChannel chan Job 			//工作对象池－工作缓冲通道
       
       ```
   * dispatcher 管理者（管理工作任务的分发和调度）
       
       ```sql
        包含: 创建、运行、任务分发
        
        任务分发就是：从工作队列JobQueue监听是否有新的工作,
        如果有的话则派发到WorkerPool工作者队列里面的JobChannel，
        worker则从该通道中取出对应的工作，并执行对应的工作任务
   
       ```
   
* 列举一些可以应用的领域

   * 分布式服务ＤＢ
   * 分布式架构请求处理







