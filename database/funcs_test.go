package database

import (
	"context"
	"reflect"
	"testing"
)

func TestIsQueryableContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "regular context",
			args: args{ctx: context.Background()},
			want: false,
		},
		{
			name: "queryable context",
			args: args{ctx: Context(context.Background(), nil)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsQueryableContext(tt.args.ctx); got != tt.want {
				t.Errorf("IsQueryableContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContext(t *testing.T) {
	ctxBackground := context.Background()
	ctx := Context(ctxBackground, nil)

	if !reflect.DeepEqual(ctx, QueryableContext{Context: ctxBackground, queryable: nil}) {
		t.Errorf("Context() = %v, want %v", ctx, QueryableContext{Context: ctxBackground, queryable: nil})
	}

	if !IsQueryableContext(ctx) {
		t.Error(`IsQueryableContext() = `, IsQueryableContext(ctx), `, want `, true)
	}
}
