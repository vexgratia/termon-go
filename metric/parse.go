package metric

type Parsed struct {
	Tag  string
	Type MetricType
}

func (m *Metric) Parse() Parsed {
	return ParseTable[m.Name]
}

var ParseTable = map[string]Parsed{
	"/cgo/go-to-c-calls:calls":                     {"Calls to C", CALLS},
	"/cpu/classes/gc/mark/assist:cpu-seconds":      {"GC assist", SECONDS},
	"/cpu/classes/gc/mark/dedicated:cpu-seconds":   {"GC dedicated", SECONDS},
	"/cpu/classes/gc/mark/idle:cpu-seconds":        {"GC idle", SECONDS},
	"/cpu/classes/gc/pause:cpu-seconds":            {"GC pause", SECONDS},
	"/cpu/classes/gc/total:cpu-seconds":            {"GC total", SECONDS},
	"/cpu/classes/idle:cpu-seconds":                {"Idle", SECONDS},
	"/cpu/classes/scavenge/assist:cpu-seconds":     {"Scavenge assist", SECONDS},
	"/cpu/classes/scavenge/background:cpu-seconds": {"Scavenge bg", SECONDS},
	"/cpu/classes/scavenge/total:cpu-seconds":      {"Scavenge total", SECONDS},
	"/cpu/classes/total:cpu-seconds":               {"Total", SECONDS},
	"/cpu/classes/user:cpu-seconds":                {"User", SECONDS},
	"/gc/cycles/automatic:gc-cycles":               {"Automatic", CYCLES},
	"/gc/cycles/forced:gc-cycles":                  {"Forced", CYCLES},
	"/gc/cycles/total:gc-cycles":                   {"Total", CYCLES},
	"/gc/heap/allocs-by-size:bytes":                {"Allocs by size", BYTES},
	"/gc/heap/allocs:bytes":                        {"Heap allocs", BYTES},
	"/gc/heap/allocs:objects":                      {"Heap allocs", OBJECTS},
	"/gc/heap/frees-by-size:bytes":                 {"Frees by size", BYTES},
	"/gc/heap/frees:bytes":                         {"Heap frees", BYTES},
	"/gc/heap/frees:objects":                       {"Heap frees", OBJECTS},
	"/gc/heap/goal:bytes":                          {"Heap goal", BYTES},
	"/gc/heap/objects:objects":                     {"Heap", OBJECTS},
	"/gc/heap/tiny/allocs:objects":                 {"Tiny allocs", OBJECTS},
	"/gc/limiter/last-enabled:gc-cycle":            {"Limiter last", CYCLE},
	"/gc/pauses:seconds":                           {"Pauses", SECONDS},
	"/gc/stack/starting-size:bytes":                {"Stack start", BYTES},
	"/memory/classes/heap/free:bytes":              {"Heap free", BYTES},
	"/memory/classes/heap/objects:bytes":           {"Heap objects", BYTES},
	"/memory/classes/heap/released:bytes":          {"Heap released", BYTES},
	"/memory/classes/heap/stacks:bytes":            {"Heap stacks", BYTES},
	"/memory/classes/heap/unused:bytes":            {"Heap unused", BYTES},
	"/memory/classes/metadata/mcache/free:bytes":   {"Mcache free", BYTES},
	"/memory/classes/metadata/mcache/inuse:bytes":  {"Mcache inuse", BYTES},
	"/memory/classes/metadata/mspan/free:bytes":    {"Mspan free", BYTES},
	"/memory/classes/metadata/mspan/inuse:bytes":   {"Mspan inuse", BYTES},
	"/memory/classes/metadata/other:bytes":         {"Metadata other", BYTES},
	"/memory/classes/os-stacks:bytes":              {"OS stacks", BYTES},
	"/memory/classes/other:bytes":                  {"Other", BYTES},
	"/memory/classes/profiling/buckets:bytes":      {"Profiling buckets", BYTES},
	"/memory/classes/total:bytes":                  {"Total", BYTES},
	"/sched/gomaxprocs:threads":                    {"GOMAXPROCS", THREADS},
	"/sched/goroutines:goroutines":                 {"Goroutines", NUMBER},
	"/sched/latencies:seconds":                     {"Sched latency", SECONDS},
	"/sync/mutex/wait/total:seconds":               {"Mutex wait", SECONDS},
}
