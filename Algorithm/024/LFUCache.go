package _24

import "math"

type LFUC struct {
	size  int
	list  map[int]int
	visit map[int]int
}

func LFUCache(capacity int) LFUC {
	return LFUC{
		size:  capacity,
		list:  make(map[int]int, 0),
		visit: make(map[int]int, 0),
	}
}

func (lfuc *LFUC) Set(key, value int) {
	if len(lfuc.list) >= lfuc.size {
		var min = math.MaxInt32
		var thenum int
		for key, value := range lfuc.visit {
			if value < min {
				thenum = key
			}
		}
		delete(lfuc.list, thenum)
		delete(lfuc.visit, thenum)

		lfuc.list[key] = value
		lfuc.visit[key] = 0
	} else {
		lfuc.list[key] = value
		lfuc.visit[key] = 0
	}

}

func (lfuc *LFUC) get(key int) int {
	if _, ok := lfuc.list[key]; !ok {
		return -1
	} else {
		lfuc.visit[key] += 1
	}
	return lfuc.list[key]
}
