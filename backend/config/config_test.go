package config

import (
	"reflect"
	"testing"
)

func Test_loadConfigFile(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{

		{
			name: "Test with test.ini",
			args: args{
				name: "./test.ini",
			},
			want: &Config{
				Server: server{
					Port: "1111",
				},
				Sqlite3: sqlite3{
					File: "spec.db",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loadConfigFile(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadConfigFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
