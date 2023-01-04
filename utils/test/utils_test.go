package utils_test

import (
	"github.com/corentings/kafejo-bot/utils"
	"testing"
)

func TestIsDefaultAvatar(t *testing.T) {
	type args struct {
		avatar string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "default avatar",
			args: args{
				avatar: "https://cdn.discordapp.com/embed/avatars/5.png",
			},
			want: true,
		},
		{
			name: "not default avatar",
			args: args{
				avatar: "https://cdn.discordapp.com/avatars/123456789012345678/123456789012345678.png",
			},
		},
		{
			name: "not default avatar",
			args: args{
				avatar: "https://cdn.discordapp.com/embed/avatars/6.png",
			},
			want: false,
		},
		{
			name: "default avatar",
			args: args{
				avatar: "https://cdn.discordapp.com/embed/avatars/0.png",
			},
			want: true,
		},
		{
			name: "default avatar",
			args: args{
				avatar: "https://cdn.discordapp.com/embed/avatars/1.png",
			},
			want: true,
		},
		{
			name: "default avatar",
			args: args{
				avatar: "https://cdn.discordapp.com/embed/avatars/2.png",
			},
			want: true,
		},
		{
			name: "default avatar",
			args: args{
				avatar: "https://cdn.discordapp.com/embed/avatars/3.png",
			},
			want: true,
		},
		{
			name: "default avatar",
			args: args{
				avatar: "https://cdn.discordapp.com/embed/avatars/4.png",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.IsDefaultAvatar(tt.args.avatar); got != tt.want {
				t.Errorf("IsDefaultAvatar() = %v, want %v", got, tt.want)
			}
		})
	}
}
