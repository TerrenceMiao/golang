package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"github.com/stretchr/testify/require"

	pbh "github.com/terrencemiao/golang/protos/hello"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
    lis = bufconn.Listen(bufSize)
    s := grpc.NewServer()
    pbh.RegisterHelloServer(s, &Server{})
	
    go func() {
        if err := s.Serve(lis); err != nil {
            log.Fatalf("Server exited with error: %v", err)
        }
    }()
}

func bufDialer(context.Context, string) (net.Conn, error) {
    return lis.Dial()
}

func TestGreetGrpc(t *testing.T) {
    ctx := context.Background()
 
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
    
	if err != nil {
        t.Fatalf("Failed to dial bufnet: %v", err)
    }
    
	defer conn.Close()
    
	client := pbh.NewHelloClient(conn)
    resp, err := client.Greet(ctx, &pbh.GreetingRequest{Name: "Dr. Seuss"})
    
	if err != nil {
        t.Fatalf("SayHello failed: %v", err)
    }
    
	log.Printf(resp.Greeting)
	fmt.Println(resp.Greeting)
}

func TestGreetNoGrpc(t *testing.T) {
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
		
		require.EqualValues(t, parameter.want, resp.Greeting)

		if parameter.want != resp.Greeting {
			t.Errorf("Great(%v) = %v, wanted %v", parameter.name, resp.Greeting, parameter.want)
		}

		log.Printf(resp.Greeting)
		fmt.Println(resp.Greeting)
	}
}
