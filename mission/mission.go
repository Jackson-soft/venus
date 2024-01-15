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

func newTask() *Task {
	task := new(Task)
	task.Params = make([]reflect.Value, 0)
	return task
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
			New: func() any { return newTask() },
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
func (e *EventBus) Producer(handler any, params ...any) error {
	fn := reflect.ValueOf(handler)
	if fn.Kind() != reflect.Func {
		return ErrTaskNotFunc
	}

	parameterLen := fn.Type().NumIn()
	if len(params) != parameterLen {
		return ErrNumberOfParameters
	}

	task, ok := e.taskPool_.Get().(*Task)
	if !ok {
		task = newTask()
	} else {
		// 清空一下参数
		task.Params = task.Params[:0]
	}

	task.Handler = fn

	for i := 0; i < parameterLen; i++ {
		t1 := reflect.TypeOf(params[i]).Kind()
		if t1 == reflect.Interface || t1 == reflect.Pointer {
			t1 = reflect.TypeOf(params[i]).Elem().Kind()
		}
		t2 := reflect.New(fn.Type().In(i)).Elem().Kind()
		if t2 == reflect.Interface || t2 == reflect.Pointer {
			t2 = reflect.Indirect(reflect.ValueOf(fn.Type().In(i))).Kind()
		}
		if t1 != t2 {
			return ErrTypeOfParameters
		}

		task.Params = append(task.Params, reflect.ValueOf(params[i]))
	}

	e.taskQueue_ <- task

	return nil
}
