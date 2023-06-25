package metric

import (
	"math"

	queue "github.com/vexgratia/collection-go/generic/queue"
)

const newCapacity = 100

type Metric struct {
	Name     string
	Current  float64
	Capacity int
	Queue    *queue.Queue[float64]
}

func New() *Metric {
	q := queue.New[float64]()
	for i := 0; i < newCapacity; i++ {
		q.Enqueue(math.NaN())
	}
	return &Metric{
		Capacity: newCapacity,
		Queue:    q,
	}
}

var AllMetrics = []string{
	"/cgo/go-to-c-calls:calls",
	"/cpu/classes/gc/mark/assist:cpu-seconds",
	"/cpu/classes/gc/mark/dedicated:cpu-seconds",
	"/cpu/classes/gc/mark/idle:cpu-seconds",
	"/cpu/classes/gc/pause:cpu-seconds",
	"/cpu/classes/gc/total:cpu-seconds",
	"/cpu/classes/idle:cpu-seconds",
	"/cpu/classes/scavenge/assist:cpu-seconds",
	"/cpu/classes/scavenge/background:cpu-seconds",
	"/cpu/classes/scavenge/total:cpu-seconds",
	"/cpu/classes/total:cpu-seconds",
	"/cpu/classes/user:cpu-seconds",
	"/gc/cycles/automatic:gc-cycles",
	"/gc/cycles/forced:gc-cycles",
	"/gc/cycles/total:gc-cycles",
	"/gc/heap/allocs-by-size:bytes",
	"/gc/heap/allocs:bytes",
	"/gc/heap/allocs:objects",
	"/gc/heap/frees-by-size:bytes",
	"/gc/heap/frees:bytes",
	"/gc/heap/frees:objects",
	"/gc/heap/goal:bytes",
	"/gc/heap/objects:objects",
	"/gc/heap/tiny/allocs:objects",
	"/gc/limiter/last-enabled:gc-cycle",
	"/gc/pauses:seconds",
	"/gc/stack/starting-size:bytes",
	"/memory/classes/heap/free:bytes",
	"/memory/classes/heap/objects:bytes",
	"/memory/classes/heap/released:bytes",
	"/memory/classes/heap/stacks:bytes",
	"/memory/classes/heap/unused:bytes",
	"/memory/classes/metadata/mcache/free:bytes",
	"/memory/classes/metadata/mcache/inuse:bytes",
	"/memory/classes/metadata/mspan/free:bytes",
	"/memory/classes/metadata/mspan/inuse:bytes",
	"/memory/classes/metadata/other:bytes",
	"/memory/classes/os-stacks:bytes",
	"/memory/classes/other:bytes",
	"/memory/classes/profiling/buckets:bytes",
	"/memory/classes/total:bytes",
	"/sched/gomaxprocs:threads",
	"/sched/goroutines:goroutines",
	"/sched/latencies:seconds",
	"/sync/mutex/wait/total:seconds",
}
var CPUMetrics = []string{
	"/cpu/classes/gc/mark/assist:cpu-seconds",
	"/cpu/classes/gc/mark/dedicated:cpu-seconds",
	"/cpu/classes/gc/mark/idle:cpu-seconds",
	"/cpu/classes/gc/pause:cpu-seconds",
	"/cpu/classes/gc/total:cpu-seconds",
	"/cpu/classes/idle:cpu-seconds",
	"/cpu/classes/scavenge/assist:cpu-seconds",
	"/cpu/classes/scavenge/background:cpu-seconds",
	"/cpu/classes/scavenge/total:cpu-seconds",
	"/cpu/classes/total:cpu-seconds",
	"/cpu/classes/user:cpu-seconds",
}
var GCMetrics = []string{
	"/gc/cycles/automatic:gc-cycles",
	"/gc/cycles/forced:gc-cycles",
	"/gc/cycles/total:gc-cycles",
	"/gc/heap/allocs-by-size:bytes",
	"/gc/heap/allocs:bytes",
	"/gc/heap/allocs:objects",
	"/gc/heap/frees-by-size:bytes",
	"/gc/heap/frees:bytes",
	"/gc/heap/frees:objects",
	"/gc/heap/goal:bytes",
	"/gc/heap/objects:objects",
	"/gc/heap/tiny/allocs:objects",
	"/gc/limiter/last-enabled:gc-cycle",
	// "/gc/pauses:seconds",
	"/gc/stack/starting-size:bytes",
}
var OtherMetrics = []string{
	"/cgo/go-to-c-calls:calls",
	"/sched/gomaxprocs:threads",
	"/sched/goroutines:goroutines",
	"/sched/latencies:seconds",
	"/sync/mutex/wait/total:seconds",
}
var MemoryMetrics = []string{
	"/memory/classes/heap/free:bytes",
	"/memory/classes/heap/objects:bytes",
	"/memory/classes/heap/released:bytes",
	"/memory/classes/heap/stacks:bytes",
	"/memory/classes/heap/unused:bytes",
	"/memory/classes/metadata/mcache/free:bytes",
	"/memory/classes/metadata/mcache/inuse:bytes",
	"/memory/classes/metadata/mspan/free:bytes",
	"/memory/classes/metadata/mspan/inuse:bytes",
	"/memory/classes/metadata/other:bytes",
	"/memory/classes/os-stacks:bytes",
	"/memory/classes/other:bytes",
	"/memory/classes/profiling/buckets:bytes",
	"/memory/classes/total:bytes",
}
