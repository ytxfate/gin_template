package routinepool

import "fmt"

// goroutine池
//
// 根据tasks的容量, 创建指定数量的goroutine来执行任务,
// 执行完成通过ress返回数据.
//
// 若 handler 需携带其他参数, 那么需要以闭包的方式实现传参, 例如:
//
//	RoutinePool(tasksChan, ressChan, func(task T) (E, error) {return handlerFunc(task, otherParams)})
func RoutinePool[T any, E any](tasks chan T, ress chan E,
	handler func(task T) (E, error),
) {
	flagChan := make(chan struct{}, cap(tasks))
	defer close(flagChan)
	// 根据 tasks 创建多个 goroutine
	for i := 0; i < cap(tasks); i++ {
		go func() {
			defer func() { flagChan <- struct{}{} }()
			for task_ := range tasks {
				res, err := handler(task_)
				if err == nil {
					ress <- res
				} else {
					fmt.Println(task_, err)
				}
			}
		}()
	}
	// 创建了多少个goroutine,就需要循环多少次,保证所有goroutine都退出
	for i := 0; i < cap(flagChan); i++ {
		<-flagChan
	}
	close(ress)
}
