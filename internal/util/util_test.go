package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinifyHTML(t *testing.T) {
	tests := []struct {
		html    string
		want    string
		wantErr error
	}{
		{
			html: `<table>
				<tr><th>header</th></tr>
				<tr><td>body</td></tr>
			</table>`,
			want:    `<table><tr><th>header</th></tr><tr><td>body</td></tr></table>`,
			wantErr: nil,
		},
	}

	for _, test := range tests {
		got, err := MinifyHTML(test.html)
		if test.wantErr != nil {
			assert.EqualError(t, err, test.wantErr.Error())
		} else {
			assert.NoError(t, err)
		}
		assert.Equal(t, test.want, got)
	}
}
