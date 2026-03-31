package mr

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io"
	"log"
	"net/rpc"
	"os"
	"sort"
	"time"
)

// Map functions return a slice of KeyValue.
type KeyValue struct {
	Key   string
	Value string
}

// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

// main/mrworker.go calls this function.
func Worker(mapf func(string, string) []KeyValue,
	reducef func(string, []string) string) {

	for {
		args := RequestTaskArgs{}
		reply := RequestTaskReply{}

		if !call("Coordinator.RequestTask", &args, &reply) {
			return
		}

		switch reply.Type {
		case TaskMap:
			runMapTask(reply, mapf)
			reportDone(TaskMap, reply.TaskID)
		case TaskReduce:
			runReduceTask(reply, reducef)
			reportDone(TaskReduce, reply.TaskID)
		case TaskWait:
			time.Sleep(500 * time.Millisecond)
		case TaskExit:
			return
		}
	}

	// uncomment to send the Example RPC to the coordinator.
	// CallExample()

}

type byKey []KeyValue

func (a byKey) Len() int           { return len(a) }
func (a byKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

func runMapTask(task RequestTaskReply, mapf func(string, string) []KeyValue) {
	content, err := os.ReadFile(task.FileName)
	if err != nil {
		log.Fatalf("cannot read %v", task.FileName)
	}

	kva := mapf(task.FileName, string(content))
	buckets := make([][]KeyValue, task.NReduce)
	for _, kv := range kva {
		r := ihash(kv.Key) % task.NReduce
		buckets[r] = append(buckets[r], kv)
	}

	for r := 0; r < task.NReduce; r++ {
		tmp, err := os.CreateTemp(".", fmt.Sprintf("mr-%d-%d-", task.TaskID, r))
		if err != nil {
			log.Fatalf("cannot create temp map output: %v", err)
		}

		enc := json.NewEncoder(tmp)
		for _, kv := range buckets[r] {
			if err := enc.Encode(&kv); err != nil {
				log.Fatalf("cannot encode intermediate key/value: %v", err)
			}
		}

		if err := tmp.Close(); err != nil {
			log.Fatalf("cannot close temp map output: %v", err)
		}

		finalName := fmt.Sprintf("mr-%d-%d", task.TaskID, r)
		if err := os.Rename(tmp.Name(), finalName); err != nil {
			log.Fatalf("cannot commit map output %v: %v", finalName, err)
		}
	}
}

func runReduceTask(task RequestTaskReply, reducef func(string, []string) string) {
	intermediate := []KeyValue{}

	for m := 0; m < task.NMap; m++ {
		name := fmt.Sprintf("mr-%d-%d", m, task.TaskID)
		file, err := os.Open(name)
		if err != nil {
			continue
		}

		dec := json.NewDecoder(file)
		for {
			var kv KeyValue
			if err := dec.Decode(&kv); err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalf("cannot decode %v: %v", name, err)
			}
			intermediate = append(intermediate, kv)
		}

		file.Close()
	}

	sort.Sort(byKey(intermediate))

	tmp, err := os.CreateTemp(".", fmt.Sprintf("mr-out-%d-", task.TaskID))
	if err != nil {
		log.Fatalf("cannot create temp reduce output: %v", err)
	}

	i := 0
	for i < len(intermediate) {
		j := i + 1
		for j < len(intermediate) && intermediate[j].Key == intermediate[i].Key {
			j++
		}

		values := make([]string, 0, j-i)
		for k := i; k < j; k++ {
			values = append(values, intermediate[k].Value)
		}

		output := reducef(intermediate[i].Key, values)
		fmt.Fprintf(tmp, "%v %v\n", intermediate[i].Key, output)

		i = j
	}

	if err := tmp.Close(); err != nil {
		log.Fatalf("cannot close temp reduce output: %v", err)
	}

	finalName := fmt.Sprintf("mr-out-%d", task.TaskID)
	if err := os.Rename(tmp.Name(), finalName); err != nil {
		log.Fatalf("cannot commit reduce output %v: %v", finalName, err)
	}
}

func reportDone(taskType TaskType, taskID int) {
	args := ReportTaskArgs{Type: taskType, TaskID: taskID}
	reply := ReportTaskReply{}
	call("Coordinator.ReportTask", &args, &reply)
}

// example function to show how to make an RPC call to the coordinator.
//
// the RPC argument and reply types are defined in rpc.go.
func CallExample() {

	// declare an argument structure.
	args := ExampleArgs{}

	// fill in the argument(s).
	args.X = 99

	// declare a reply structure.
	reply := ExampleReply{}

	// send the RPC request, wait for the reply.
	// the "Coordinator.Example" tells the
	// receiving server that we'd like to call
	// the Example() method of struct Coordinator.
	ok := call("Coordinator.Example", &args, &reply)
	if ok {
		// reply.Y should be 100.
		fmt.Printf("reply.Y %v\n", reply.Y)
	} else {
		fmt.Printf("call failed!\n")
	}
}

// send an RPC request to the coordinator, wait for the response.
// usually returns true.
// returns false if something goes wrong.
func call(rpcname string, args interface{}, reply interface{}) bool {
	// c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	sockname := coordinatorSock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		return false
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}
