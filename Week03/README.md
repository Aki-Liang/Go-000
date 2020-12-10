# 第三周

## Goroutine

### 线程和进程

进程
* 运行应用程序时操作系统会为该应用程序创建一个进程。作为一个应用程序，它像一个为所有资源而运行的容器，包括内存地址空间、文件句柄、设备和线程

线程
* 线程是操作系统调度的一种知性路径，用于处理器执行我们再函数中编写的代码。一个进程从一个线程开始（主线程）， 当该线程终止时进程终止

### Goroutiones and Parallelism(协程和并行)

* go关键字可以快速的将一个函数创建为goroutione。我们可以认为main函数就是作为groutine执行的。
* 操作系统调度线程在可用处理器上运行
* go运行时调度goroutines在绑定到单个操作系统线程的逻辑处理器中运行（P）
* 即使使用单一的逻辑处理器和操作系统线程也可以调度数十万goroutine以惊人的效率和性能并发运行
  
Concurrency is not Parallelism(并发不是并行)
* 并行是指两个或多个线程同时在不同的处理器执行代码
* 将运行时配置为使用多个逻辑处理器，则调度程序将在这些逻辑处理器之间分配goroutine
* 想要获得真正的并行性需要在具有多个物理处理器的计算机上运行程序，否则即使GO运行时使用多个逻辑处理器goroutines仍然会针对单个物理处理器并发运行

### Rules

#### Keep yourself busy or do the work yourself

* 如果goroutine在从另一个goroutine获得结果之前无法取得进展，那么通常情况下自己去做这项工作比委托它更简单
* 这样能够消除将结果从goroutine返回到其启动器所需要的大量状态跟踪和chan操作
  
```
//不要这样做
func main(){
    ……
    go func(){
        if err:=http.ListenAndServer(":8080", nil); err != nil {
            log.Fatal(err)
        }
    }()
    ……
    select{}  //空select语句将永远阻塞
}

//要这样做
func foo(){
    ……
    if err:=http.ListenAndServer(":8080", nil); err != nil {
        log.Fatal(err)
    }
    ……
}
```
#### Leave concurrency to the caller(将异步执行函数的决定权交给该函数的调用方)

```
func ListDirectory(dir string)([]string, error)

func ListDirectory(dir string) chan string
```
chan版本的问题：
* 通过使用一个关闭通道操作来作为不再需要处理的信号，无法告诉调用者因为中途遇到了错误导致通过chan返回的项目集不完整。调用者无法区分空目录与读取目录出错的区别，这两种情况都会导致从ListDirectory返回的通道立即关闭
* 调用者必须持续从chan中读取，直到chan关闭，这是唯一能够知道goroutine已经停止的方法。这是对ListDirectory使用方法的一个严重限制，及时调用者可能已经收到了想要的数据也必须花时间持续读取chan。对于大中型目录该版本可能在内存使用方面更高效，但并不比原始的基于slice的方法快

```
func ListDirectory(dir string, fn func(string))
```
* filepath.WalkDir的模型，如果函数启动groutine则必须向调用者提供显示停止该goroutine的方法

#### Never start a goroutine without knowing when it will stop(不要开启一个不知道何时结束的goroutine)

* 使用`sync.WaitGroup`来追踪每一个创建的goroutine
* 使用`golang.org/x/sync/errgroup`来追踪创建的goroutine和运行结果
* 使用`context`处理超时
  
## Memory model

https://golang.org/ref/mem

### Happen-Before

## Package sync

## chan

## Package context


## 引用文档

https://www.ardanlabs.com/blog/2018/11/goroutine-leaks-the-forgotten-sender.html
https://www.ardanlabs.com/blog/2019/04/concurrency-trap-2-incomplete-work.html
https://www.ardanlabs.com/blog/2014/01/concurrency-goroutines-and-gomaxprocs.html
https://dave.cheney.net/practical-go/presentations/qcon-china.html#_concurrency
https://golang.org/ref/mem
https://blog.csdn.net/caoshangpa/article/details/78853919
https://blog.csdn.net/qcrao/article/details/92759907
https://cch123.github.io/ooo/
https://blog.golang.org/codelab-share
https://dave.cheney.net/2018/01/06/if-aligned-memory-writes-are-atomic-why-do-we-need-the-sync-atomic-package
http://blog.golang.org/race-detector
https://dave.cheney.net/2014/06/27/ice-cream-makers-and-data-races
https://www.ardanlabs.com/blog/2014/06/ice-cream-makers-and-data-races-part-ii.html
https://medium.com/a-journey-with-go/go-how-to-reduce-lock-contention-with-the-atomic-package-ba3b2664b549
https://medium.com/a-journey-with-go/go-discovery-of-the-trace-package-e5a821743c3c
https://medium.com/a-journey-with-go/go-mutex-and-starvation-3f4f4e75ad50
https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html
https://medium.com/a-journey-with-go/go-buffered-and-unbuffered-channels-29a107c00268
https://medium.com/a-journey-with-go/go-ordering-in-select-statements-fd0ff80fd8d6
https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html
https://www.ardanlabs.com/blog/2014/02/the-nature-of-channels-in-go.html
https://www.ardanlabs.com/blog/2013/10/my-channel-select-bug.html
https://blog.golang.org/io2013-talk-concurrency
https://blog.golang.org/waza-talk
https://blog.golang.org/io2012-videos
https://blog.golang.org/concurrency-timeouts
https://blog.golang.org/pipelines
https://www.ardanlabs.com/blog/2014/02/running-queries-concurrently-against.html
https://blogtitle.github.io/go-advanced-concurrency-patterns-part-3-channels/
https://www.ardanlabs.com/blog/2013/05/thread-pooling-in-go-programming.html
https://www.ardanlabs.com/blog/2013/09/pool-go-routines-to-process-task.html
https://blogtitle.github.io/categories/concurrency/
https://medium.com/a-journey-with-go/go-context-and-cancellation-by-propagation-7a808bbc889c
https://blog.golang.org/context
https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html
https://golang.org/ref/spec#Channel_types
https://drive.google.com/file/d/1nPdvhB0PutEJzdCq5ms6UI58dp50fcAN/view
https://medium.com/a-journey-with-go/go-context-and-cancellation-by-propagation-7a808bbc889c
https://blog.golang.org/context
https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html
https://golang.org/doc/effective_go.html#concurrency
https://zhuanlan.zhihu.com/p/34417106?hmsr=toutiao.io
https://talks.golang.org/2014/gotham-context.slide#1
https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39


