package utils_test

import (
	"testing"
	"time"

	"github.com/corentings/kafejo-bot/utils"
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

func TestFormatTimeSince(t *testing.T) {
	// Define test cases with input and expected output
	testCases := []struct {
		name           string
		oldTime        time.Time
		expectedOutput string
	}{
		{
			name:           "Less than a minute ago",
			oldTime:        time.Now().Add(-30 * time.Second),
			expectedOutput: "30 seconds ago",
		},
		{
			name:           "Less than an hour ago",
			oldTime:        time.Now().Add(-45 * time.Minute),
			expectedOutput: "45 minutes ago",
		},
		{
			name:           "Less than a day ago",
			oldTime:        time.Now().Add(-10 * time.Hour),
			expectedOutput: "10 hours ago",
		},
		{
			name:           "Less than a month ago",
			oldTime:        time.Now().Add(-2 * 24 * time.Hour),
			expectedOutput: "0m 2d 0h 0m 0s",
		},
		{
			name:           "Less than a year ago",
			oldTime:        time.Now().Add(-3 * 30 * 24 * time.Hour),
			expectedOutput: "3m 0d 0h 0m 0s",
		},
		{
			name:           "More than a year ago",
			oldTime:        time.Now().Add(-2 * 365 * 24 * time.Hour),
			expectedOutput: "2y 0m 0d 0h 0m 0s",
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := utils.FormatTimeSince(tc.oldTime)
			if output != tc.expectedOutput {
				t.Errorf("Expected output: %s, but got: %s", tc.expectedOutput, output)
			}
		})
	}
}
