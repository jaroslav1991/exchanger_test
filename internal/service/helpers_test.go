package service

import (
	"reflect"
	"testing"
)

func Test_findCombinations(t *testing.T) {
	notes := []int{5000, 2000, 1000, 500, 200, 100, 50}
	notes2 := []int{500, 200, 100, 50, 25}

	type args struct {
		banknotes []int
		amount    int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			name: "success_1",
			args: args{
				banknotes: notes,
				amount:    300,
			},
			want: [][]int{
				{200, 100},
				{200, 50, 50},
				{100, 100, 100},
				{100, 100, 50, 50},
				{100, 50, 50, 50, 50},
				{50, 50, 50, 50, 50, 50},
			},
			wantErr: false,
		}, {
			name: "incorrect_amount",
			args: args{
				banknotes: notes,
				amount:    240,
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "incorrect_amount_2",
			args: args{
				banknotes: notes2,
				amount:    223,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "zero amount",
			args: args{
				banknotes: notes,
				amount:    0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "negative amount",
			args: args{
				banknotes: notes,
				amount:    -200,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findCombinations(tt.args.banknotes, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("findCombinations error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findCombinations got = %v, want %v", got, tt.want)
			}
		})
	}
}
