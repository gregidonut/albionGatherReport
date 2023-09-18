package gatherReport

import (
	"reflect"
	"testing"
	"time"
)

func Test_newCost(t *testing.T) {
	type args struct {
		kwargs map[string]int
	}
	tests := []struct {
		name string
		args args
		want *Cost
	}{
		{
			name: "empty",
			args: args{
				kwargs: nil,
			},
			want: &Cost{
				JourneyBack: 0,
				PorkPie:     0,
				Repair:      0,
				total:       0,
			},
		},
		{
			name: "all kwargs filled",
			args: args{
				kwargs: map[string]int{
					"journeyBack": 2_000,
					"porkPie":     3,
					"repair":      4_000,
				},
			},
			want: &Cost{
				JourneyBack: 2_000,
				PorkPie:     3, //15_000
				Repair:      4_000,
				total:       21_000,
			},
		},
		{
			name: "only one of the kwargs filled",
			args: args{
				kwargs: map[string]int{
					"repair": 4_000,
				},
			},
			want: &Cost{
				JourneyBack: 0,
				PorkPie:     0, //15_000
				Repair:      4_000,
				total:       4_000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newCost(tt.args.kwargs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCost() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGatherReportRow(t *testing.T) {
	type args struct {
		url    string
		kwargs map[string]int
	}
	tests := []struct {
		name string
		args args
		want *GatherReport
	}{
		{
			name: "empty",
			args: args{
				url:    "hello",
				kwargs: nil,
			},
			want: &GatherReport{
				Date: time.Time{},
				Cost: &Cost{
					JourneyBack: 0,
					PorkPie:     0,
					Repair:      0,
					total:       0,
				},
				Profit:  0,
				Revenue: 0,
				URL:     "hello",
			},
		},
		{
			name: "all fields used",
			args: args{
				url: "hello",
				kwargs: map[string]int{
					"journeyBack": 2_000,
					"porkPie":     3,
					"repair":      4_000,
					"revenue":     40_000,
				},
			},
			want: &GatherReport{
				Date: time.Time{},
				Cost: &Cost{
					JourneyBack: 2_000,
					PorkPie:     3, //15_000
					Repair:      4_000,
					total:       21_000,
				},
				Profit:  19_000,
				Revenue: 40_000,
				URL:     "hello",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGatherReportRow(tt.args.url, tt.args.kwargs)
			tt.want.Date = got.Date // stubbing this field out
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGatherReportRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
