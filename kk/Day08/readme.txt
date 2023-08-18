接口赋值
    原则：赋值（=后面）的对象必须要全部实现被赋值对象（=左边）的所有方法



并发编程开发将一个过程按照并行算法拆分为多个可以独立执行的代码


不带缓冲区的
    写入的时候没有人(goroutine)读会阻塞
    在读的时候没有人(goroutine)写会阻塞
带缓冲区的
    写入的时候缓冲区已满，没有人(goroutine)读会阻塞
    在读的时候没有数据，没有人(goroutine)写会阻塞


监听多个管道
    <- A
    <- B
    select case

在main函数内有功能要求在5s之内执行完成，如果没有执行完成程序自动结束。




复习：
例程： goroutine
    语法： go 函数调用
    分类： 主例程（main）、工作例程（使用go关键字启动的）
通信： 共享数据（同步、sync（WaitGroup（add、done、wait）、Mutex（Lock、Unlock）、RWMutex、Map、Once））、管道（异步、select case）
    管道：
        定义：var name chan type
        初始化： make(chan type,cap)  有缓冲区管道
                make(chan type)  无缓冲区管道
        操作：读： rt := <- channel
             写： channel <-
             关闭：
             for range
        只读/只写：函数之间参数传递，在函数内部只进行读/只进行写，将函数参数声明为只读/只写管道
            只读（声明）：  <-chan type
            只写（声明）：  chan <- type
        select-case：
            select{
                case: <- channel1:
                case: <- channel2:
                case channel3 <-:
                case channel4 <-:
                default:
            }
    Once：只执行一次
    Map：
    Pool：

        

time包：
    延迟： After(duration) <- chan Time
    定时： 
        Tick(duration) <- chan Time
        for-range



runtime包：
    Goshed：例程让出
    NumCPU：
    NumGoroutine：
    GOMAXPROCS

注意事项：waitgroup是一个结构体，函数之间传递的话需要用指针

go MPG调度模型