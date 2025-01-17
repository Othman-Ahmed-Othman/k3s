package config

import (
	"reflect"
	"testing"
)

func Test_UnitGetArgsList(t *testing.T) {
	type args struct {
		argsMap   map[string]string
		extraArgs []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Default Test",
			args: args{
				argsMap: map[string]string{
					"aaa": "A",
					"bbb": "B",
					"ccc": "C",
					"ddd": "d",
					"eee": "e",
					"fff": "f",
					"ggg": "g",
					"hhh": "h",
				},
				extraArgs: []string{
					"bbb=BB",
					"ddd=DD",
					"iii=II",
				},
			},

			want: []string{
				"--aaa=A",
				"--bbb=BB",
				"--ccc=C",
				"--ddd=DD",
				"--eee=e",
				"--fff=f",
				"--ggg=g",
				"--hhh=h",
				"--iii=II",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetArgsList(tt.args.argsMap, tt.args.extraArgs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArgsList() = %+v\nWant = %+v", got, tt.want)
			}
		})
	}
}
