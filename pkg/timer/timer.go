package timer

import (
	"errors"
	"time"
)

/**
 * 用于计算一段代码的执行时间。
 * 时间单位是us
 * calculate execution time
 */

var startTime map[string]int64
var endTime map[string]int64
var executionTime map[string]int64

func start(key string) {
	startTime[key] = time.Now().UnixNano() / 1000
	endTime[key] = startTime[key]
	executionTime[key] = 0
}

func end(key string) {
	endTime[key] = time.Now().UnixNano() / 1000
	if _, ok := startTime[key]; ok {
		executionTime[key] = endTime[key] - startTime[key]
	} else {
		executionTime[key] = 0
	}
}

func fetchOne(key string) (int64, error) {
	if _, ok := executionTime[key]; ok {
		return executionTime[key], nil
	} else {
		return 0, errors.New("key is not exist")
	}
}

func fetchAll() map[string]int64 {
	return executionTime
}

func init() {
	startTime = make(map[string]int64)
	endTime = make(map[string]int64)
	executionTime = make(map[string]int64)
}
