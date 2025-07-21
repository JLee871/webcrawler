package main

import (
	"net/url"
	"reflect"
	"testing"
)

func Test_getURLsFromHTML(t *testing.T) {
	type args struct {
		htmlBody   string
		rawBaseURL string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "absolute and relative URLs",
			args: args{
				htmlBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
				rawBaseURL: "https://blog.boot.dev",
			},
			want: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},

		{
			name: "empty htmlBody",
			args: args{
				htmlBody: `
<html>
	<body>
	</body>
</html>
`,
				rawBaseURL: "https://blog.boot.dev",
			},
			want: nil,
		},

		{
			name: "contains invalid paths",
			args: args{
				htmlBody: `
<html>
	<body>
		<a href="path/one">
			<span>Boot.dev</span>
		</a>
		<a href=":\\invalidURL">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
				rawBaseURL: "https://blog.boot.dev",
			},
			want: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},

		{
			name: "malformed base url",
			args: args{
				htmlBody: `
<html>
	<body>
	</body>
</html>
`,
				rawBaseURL: "://invalidURL",
			},
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.args.rawBaseURL)

			got, err := getURLsFromHTML(tt.args.htmlBody, baseURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("getURLsFromHTML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getURLsFromHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
