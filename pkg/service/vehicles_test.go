package vehicles

import (
	"context"
	"testing"
)

func TestStart(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"start",
			args{ctx: context.Background()},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Start(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStop(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"stop",
			args{ctx: context.Background()},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Stop(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Stop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
