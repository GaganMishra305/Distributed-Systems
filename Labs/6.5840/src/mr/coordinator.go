package mr

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sync"
	"time"
)

type Coordinator struct {
	mu sync.Mutex

	files   []string
	nReduce int

	mapTasks    []taskState
	reduceTasks []taskState
	done        bool
}

type taskStatus int

const (
	statusIdle taskStatus = iota
	statusInProgress
	statusCompleted
)

type taskState struct {
	status    taskStatus
	startedAt time.Time
}

const taskTimeout = 10 * time.Second

// Your code here -- RPC handlers for the worker to call.

func (c *Coordinator) RequestTask(args *RequestTaskArgs, reply *RequestTaskReply) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.done {
		reply.Type = TaskExit
		return nil
	}

	if !c.allMapsDone() {
		if id, ok := c.pickTask(c.mapTasks); ok {
			c.mapTasks[id].status = statusInProgress
			c.mapTasks[id].startedAt = time.Now()

			reply.Type = TaskMap
			reply.TaskID = id
			reply.FileName = c.files[id]
			reply.NReduce = c.nReduce
			reply.NMap = len(c.files)
			return nil
		}
		reply.Type = TaskWait
		return nil
	}

	if !c.allReducesDone() {
		if id, ok := c.pickTask(c.reduceTasks); ok {
			c.reduceTasks[id].status = statusInProgress
			c.reduceTasks[id].startedAt = time.Now()

			reply.Type = TaskReduce
			reply.TaskID = id
			reply.NReduce = c.nReduce
			reply.NMap = len(c.files)
			return nil
		}
		reply.Type = TaskWait
		return nil
	}

	c.done = true
	reply.Type = TaskExit
	return nil
}

func (c *Coordinator) ReportTask(args *ReportTaskArgs, reply *ReportTaskReply) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	switch args.Type {
	case TaskMap:
		if args.TaskID >= 0 && args.TaskID < len(c.mapTasks) {
			c.mapTasks[args.TaskID].status = statusCompleted
		}
	case TaskReduce:
		if args.TaskID >= 0 && args.TaskID < len(c.reduceTasks) {
			c.reduceTasks[args.TaskID].status = statusCompleted
		}
	}

	if c.allMapsDone() && c.allReducesDone() {
		c.done = true
	}

	return nil
}

func (c *Coordinator) allMapsDone() bool {
	for _, t := range c.mapTasks {
		if t.status != statusCompleted {
			return false
		}
	}
	return true
}

func (c *Coordinator) allReducesDone() bool {
	for _, t := range c.reduceTasks {
		if t.status != statusCompleted {
			return false
		}
	}
	return true
}

func (c *Coordinator) pickTask(tasks []taskState) (int, bool) {
	now := time.Now()
	for i := range tasks {
		if tasks[i].status == statusIdle {
			return i, true
		}
		if tasks[i].status == statusInProgress && now.Sub(tasks[i].startedAt) > taskTimeout {
			return i, true
		}
	}
	return -1, false
}

// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}

// start a thread that listens for RPCs from worker.go
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := coordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
func (c *Coordinator) Done() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.done
}

// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{
		files:       files,
		nReduce:     nReduce,
		mapTasks:    make([]taskState, len(files)),
		reduceTasks: make([]taskState, nReduce),
	}

	c.server()
	return &c
}
