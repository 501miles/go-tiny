package timeline

import (
	"errors"
	"github.com/501miles/go-tiny/tool/gen_id/snowflake"
	"github.com/501miles/go-tiny/tool/time_tool"
	"math"
	"sync"
	"time"
)

var (
	TIMELINE_ERROR_ALREADY_STARTED = errors.New("Timeline already started.")
)

type Timeline struct {
	lock     sync.Mutex
	Id       uint16
	Gap      time.Duration
	JobDict  map[int64]map[int64]func()
	JobIndex map[int64]int64
}

var timelineDict = make(map[uint16]*Timeline)
var lock sync.RWMutex

func CreateTimeline(index uint16, gap time.Duration) (*Timeline, error) {
	lock.Lock()
	defer lock.Unlock()
	if _, ok := timelineDict[index]; ok {
		return nil, TIMELINE_ERROR_ALREADY_STARTED
	}
	timeline := &Timeline{
		Id:       index,
		JobDict:  make(map[int64]map[int64]func()),
		JobIndex: make(map[int64]int64),
		Gap:      gap,
	}
	timelineDict[index] = timeline
	return timeline, nil
}

func GetTimeline(index uint16) *Timeline {
	v, _ := timelineDict[index]
	return v
}

func (t *Timeline) Start() {
	ticker := time.NewTicker(t.Gap)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			t.processOnce()
		}
	}
}

func (t *Timeline) processOnce() {
	t1 := time_tool.NowTimeUnix10()
	t.lock.Lock()
	dict, ok := t.JobDict[t1]
	if ok {
		var funcList []func()
		for k, v := range dict {
			delete(t.JobIndex, k)
			funcList = append(funcList, v)
		}
		delete(t.JobDict, t1)
		t.lock.Unlock()
		for _, f := range funcList {
			go func() {
				f()
			}()
			time.Sleep(7 * time.Millisecond)
		}
	} else {
		t.lock.Unlock()
	}
}

func (t *Timeline) Register(tm time.Time, f func()) int64 {
	t.lock.Lock()
	defer t.lock.Unlock()
	id := snowflake.GenInt64()
	t0 := float64(tm.UnixNano()) * 0.000000001
	t1 := int64(math.Ceil(t0))
	tDict, _ := t.JobDict[t1]
	if tDict == nil {
		tDict = map[int64]func(){}
	}
	tDict[id] = f
	t.JobDict[t1] = tDict
	t.JobIndex[id] = t1
	return id
}

func (t *Timeline) UnRegister(id int64) {
	t.lock.Lock()
	defer t.lock.Unlock()
	indexId, _ := t.JobIndex[id]
	if indexId > 0 {
		subDict, _ := t.JobDict[indexId]
		if len(subDict) == 1 {
			delete(t.JobDict, indexId)
		} else {
			delete(t.JobDict[indexId], id)
		}
		delete(t.JobIndex, id)
	}
}
