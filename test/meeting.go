package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	v1 "meeting/api/meeting/v1"
	"sync"
	"sync/atomic"
	"time"
)

var meetingClient v1.MeetingClient
var conn *grpc.ClientConn

//func main() {
//	//Init()
//	//TestCreateMeeting()
//	//conn.Close()
//	TestGoroutine()
//	//ch := make(chan int, 10)
//	//ch <- 1
//	//ch <- 2
//	//ch <- 3
//	//
//	//// 关闭函数非常重要,若不执行close(),那么range将无法结束,造成死循环
//	//// close(ch)
//	//
//	//for v := range ch {
//	//	fmt.Println(v)
//	//}
//	//
//}

var x int64
var l sync.Mutex
var wg sync.WaitGroup

// 普通版加函数
func add() {
	// x = x + 1
	x++ // 等价于上面的操作
	wg.Done()
}

// 互斥锁版加函数
func mutexAdd() {
	l.Lock()
	x++
	l.Unlock()
	wg.Done()
}

// 原子操作版加函数
func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

//func main() {
//
//	start := time.Now()
//	for i := 0; i < 10000; i++ {
//		wg.Add(1)
//		//go add()       // 普通版add函数 不是并发安全的
//		 //go mutexAdd()  // 加锁版add函数 是并发安全的，但是加锁性能开销大
//		go atomicAdd() // 原子操作版add函数 是并发安全，性能优于加锁版
//	}
//	wg.Wait()
//	end := time.Now()
//	fmt.Println(x)
//	fmt.Println(end.Sub(start))
//}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				select {
				case <-ctx.Done():
					return
				case ch <- i:
				}
			}
		}()
		return ch
	}(ctx)

	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			cancel()
			break
		}
	}
}

type job struct {
	id     int
	random int
	result int
}

///携程池的使用
//func TestGoroutine()  {
//	jobchanel :=make(chan *job,10)
//	reschanel:=make(chan *job,10)
//	createGrountines(1000,jobchanel,reschanel)
//	go func(resultChan chan *job) {
//		// 遍历结果管道打印
//		for  {
//			result := <- resultChan
//			fmt.Printf("job id:%v randnum:%v result:%d\n", result.id,
//				result.random, result.result)
//		}
//
//	}(reschanel)
//	var id int
//	// 循环创建job，输入到管道
//	for {
//		fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
//		id++
//		// 生成随机数
//		r_num := rand.Int()
//		job := &job{
//			id:      id,
//			random: r_num,
//		}
//		jobchanel <- job
//	}
//}
func createGrountines(num int, jobchanel chan *job, reschanel chan *job, single chan int) {
	for i := 0; i < num; i++ {
		go func(jobchanel chan *job, reschanel chan *job) {
			for {

				select {
				case <-jobchanel:
					fmt.Printf("jobchanel: %d\n", len(jobchanel))
					job := <-jobchanel
					r_num := job.random
					var sum int
					for r_num != 0 {
						tmp := r_num % 10
						sum += tmp
						r_num /= 10
					}
					job.result = sum
					reschanel <- job
				case single <- 1:
					break
				}

			}

		}(jobchanel, reschanel)
	}
}

// Init 初始化 grpc 链接
func Init() {
	var err error

	conn, err = grpc.Dial("127.0.0.1:9300", grpc.WithInsecure())
	if err != nil {
		panic("grpc link err" + err.Error())
	}
	meetingClient = v1.NewMeetingClient(conn)

}

func TestCreateMeeting() {
	meeting := &v1.MeetingRequest_Meeting{
		Name:      "NewRrandMeet",
		Address:   "疫情高风险区",
		AppDeatil: "高峰论坛",
	}
	rsp, err := meetingClient.Create(context.Background(), &v1.MeetingRequest{
		Meeting: meeting,
	})
	if err != nil {
		panic("grpc 创建失败" + err.Error())
	}
	fmt.Println(rsp)
}

func TestRegisterMeeting() {
	//建立链接
	meeting := &v1.RegisterRequest_Meeting{
		Name: "NewRrandMeet",
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	ctx = context.WithValue(ctx, "UserId", "001")
	rsp, err := meetingClient.Register(ctx, &v1.RegisterRequest{
		Meeting: meeting,
	})
	if err != nil {
		panic("grpc 创建失败" + err.Error())
	}
	fmt.Println(rsp)
}
