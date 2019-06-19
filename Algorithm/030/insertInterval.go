package _30

type Interval struct {
	Start, End int
}

//给出一个无重叠的按照区间起始端点排序的区间列表。
//
//在列表中插入一个新的区间，你要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
func insert(intervals []*Interval, newInterval *Interval) []*Interval {

	result := make([]*Interval, 0)
	for i := 0; i < len(intervals); i++ {
		if newInterval == nil {
			result = append(result, intervals[i:]...)
			break
		}
		if intervals[i].End < newInterval.Start || intervals[i].Start > newInterval.End {
			if intervals[i].Start > newInterval.End {
				result = append(result, newInterval)
				newInterval = nil
			}
			result = append(result, intervals[i])

		} else if intervals[i].Start <= newInterval.Start && newInterval.Start <= intervals[i].End {
			if newInterval.End <= intervals[i].End {
				result = append(result, intervals[i])
				newInterval = nil
			} else {
				newInterval.Start = intervals[i].Start
			}
		} else if intervals[i].Start <= newInterval.End && newInterval.End < intervals[i].End {
			if intervals[i].Start <= newInterval.Start {
				result = append(result, intervals[i])
				newInterval = nil
			} else {
				newInterval.End = intervals[i].End
			}
		}
	}
	if newInterval != nil {
		result = append(result, newInterval)
	}
	return result
}

func insertNew(intervals []*Interval, newInterval *Interval) []*Interval {
	result := make([]*Interval, 0)

	idx := 0
	for idx < len(intervals) && intervals[idx].Start < newInterval.Start {
		idx++
	}
	it := append([]*Interval{}, intervals[:idx]...)
	it = append(it, newInterval)
	it = append(it, intervals[idx:]...)
	var last *Interval
	for _, il := range it {
		if last == nil || last.End < il.Start {
			result = append(result, il)
			last = il
		} else {
			last.End = max(last.End, il.End)
		}
	}
	return result
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
