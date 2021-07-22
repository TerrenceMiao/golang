package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	pbh "github.com/terrencemiao/golang/protos/hello"
)

func TestHello(t *testing.T) {
	srv := Server{}

	parameters := []struct{
		name string
		want string
	} {
		{
			name: "world",
			want: "Hello world",
		},
		{
			name: "Mark",
			want: "Hello Mark",
		},
	}

	for _, parameter := range parameters {
		req := &pbh.GreetingRequest{Name: parameter.name}

		resp, err := srv.Greet(context.Background(), req)

		if err != nil {
			t.Errorf("Greet(%v) got unexpected error", parameter.name)
		}
		
		fmt.Println(resp.Greeting)
		require.EqualValues(t, parameter.want, resp.Greeting)
		if parameter.want != resp.Greeting {
			t.Errorf("Great(%v) = %v, wanted %v", parameter.name, resp.Greeting, parameter.want)
		}
	}
}
