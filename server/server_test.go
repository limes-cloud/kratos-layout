/**
 * @Author: 1280291001@qq.com
 * @Description:
 * @File: server_test.go
 * @Version: 1.0.0
 * @Date: 2023/9/5 10:48
 */

package server

import (
	"context"
	"fmt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	v1 "github.com/limes-cloud/kratos-layout/api/v1"
	"github.com/limes-cloud/kratos/log"
	"github.com/limes-cloud/kratos/middleware/auth/jwt"
	"github.com/limes-cloud/kratos/transport/grpc"
	"testing"
)

func TestClient(t *testing.T) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			jwt.Client(func(token *jwtv4.Token) (interface{}, error) {
				return []byte("hello world"), nil
			}),
		),
	)
	if err != nil {
		t.Fatal(err)
	}

	client := v1.NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &v1.HelloRequest{Name: "ddd1"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
