package mission

import (
	"reflect"
	"sync"
)

// 简单的异步队列

// 消息总线
type EventBus struct {
	taskQueue_ chan *Task // 任务队列
	taskPool_  sync.Pool  // 任务池
}

// 消息体
type Task struct {
	Handler reflect.Value   // 函数体
	Params  []reflect.Value // 函数参数
}

var (
	once     sync.Once
	instance *EventBus
)

func Instance() *EventBus {
	once.Do(func() {
		instance = create()
	})

	return instance
}

func create() *EventBus {
	eb := &EventBus{
		taskQueue_: make(chan *Task, 8),
		taskPool_: sync.Pool{
			New: func() any { return new(Task) },
		},
	}

	go eb.run()

	return eb
}

func (e *EventBus) run() {
	defer func() {
		_ = recover()
	}()

	for task := range e.taskQueue_ {
		e.consumer(task)
	}
}

func (e *EventBus) consumer(task *Task) {
	task.Handler.Call(task.Params)
	e.taskPool_.Put(task)
}

// 参数：函数体与入参
func (e *EventBus) Producer(handler any, params ...any) {
	if reflect.TypeOf(handler).Kind() != reflect.Func {
		return
	}

	task, ok := e.taskPool_.Get().(*Task)
	if !ok {
		task = new(Task)
	}

	task.Handler = reflect.ValueOf(handler)

	task.Params = make([]reflect.Value, 0, len(params))
	for i := range params {
		task.Params = append(task.Params, reflect.ValueOf(params[i]))
	}

	e.taskQueue_ <- task
}
