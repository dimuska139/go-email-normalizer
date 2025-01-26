package emailnormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProtonmailUsername(t *testing.T) {
	rule := ProtonmailRule{}
	assert.Equal(t, "t", rule.ProcessUsername("t+.est"))
}

func TestProtonmailDomain(t *testing.T) {
	rule := ProtonmailRule{}
	assert.Equal(t, "protonmail.ch", rule.ProcessDomain("protonmail.ch"))
}

func TestProtonmailRule_ProcessUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "With plus sign",
			args: args{username: "t+est"},
			want: "t",
		}, {
			name: "With dot",
			args: args{username: "t.est"},
			want: "test",
		}, {
			name: "With underscore",
			args: args{username: "t_est"},
			want: "test",
		}, {
			name: "With underscore",
			args: args{username: "t-est"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := &ProtonmailRule{}
			assert.Equal(t, tt.want, rule.ProcessUsername(tt.args.username))
		})
	}
}
