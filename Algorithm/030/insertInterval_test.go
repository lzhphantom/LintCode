package _30

import (
	"reflect"
	"testing"
)

func TestInsertInterval(t *testing.T) {
	tests := []struct {
		intervals   []*Interval
		newInterval *Interval
		want        []*Interval
	}{
		{
			[]*Interval{
				{1, 2},
				{5, 9},
			},
			&Interval{
				2,
				5,
			},
			[]*Interval{
				{1, 9},
			},
		},
		{
			[]*Interval{
				{1, 5},
				{7, 8},
				{10, 13},
			},
			&Interval{
				6, 6,
			},
			[]*Interval{
				{1, 5},
				{6, 6},
				{7, 8},
				{10, 13},
			},
		},
		{
			[]*Interval{
				{1, 5},
				{7, 8},
				{10, 13},
				{16, 19},
				{21, 23},
			},
			&Interval{
				5, 7,
			},
			[]*Interval{
				{1, 8},
				{10, 13},
				{16, 19},
				{21, 23},
			},
		},
	}

	for _, test := range tests {
		result := insert(test.intervals, test.newInterval)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("insert(%v,%v) => %v,want %v", test.intervals, test.newInterval, result, test.want)
		}
	}
}

func TestInsertInterval1(t *testing.T) {
	tests := []struct {
		intervals   []*Interval
		newInterval *Interval
		want        []*Interval
	}{
		{
			[]*Interval{
				{1, 2},
				{5, 9},
			},
			&Interval{
				2,
				5,
			},
			[]*Interval{
				{1, 9},
			},
		},
		{
			[]*Interval{
				{1, 5},
				{7, 8},
				{10, 13},
			},
			&Interval{
				6, 6,
			},
			[]*Interval{
				{1, 5},
				{6, 6},
				{7, 8},
				{10, 13},
			},
		},
		{
			[]*Interval{
				{1, 5},
				{7, 8},
				{10, 13},
				{16, 19},
				{21, 23},
			},
			&Interval{
				5, 7,
			},
			[]*Interval{
				{1, 8},
				{10, 13},
				{16, 19},
				{21, 23},
			},
		},
	}

	for _, test := range tests {
		result := insertNew(test.intervals, test.newInterval)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("insert(%v,%v) => %v,want %v", test.intervals, test.newInterval, result, test.want)
		}
	}
}
