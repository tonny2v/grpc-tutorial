package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
	pb "tonny2v/proto"
)

type Result struct {
	First  uint64 `json:"firstnumber"`
	Second uint64 `json:"lastnumber"`
	Result uint64 `json:"result"`
}

func main() {

	r := gin.Default()
	r.GET("/add/:a/:b", Add)
	r.Run(":8080")
}

func Add(c *gin.Context) {
	a := c.Params.ByName("a")
	b := c.Params.ByName("b")
	first, _ := strconv.ParseUint(a, 10, 64)
	second, _ := strconv.ParseUint(b, 10, 64)
	r := Result{First: first, Second: second, Result: CallRpcAdd(first, second)}
	c.JSON(200, r)
}

func CallRpcAdd(a uint64, b uint64) uint64 {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
	defer cancel()
	conn, err := grpc.Dial("add-service:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial Failed: %v", err)
	}
	addClient := pb.NewAddServiceClient(conn)
	req := &pb.AddRequest{A: a, B: b}
	if resp, err := addClient.Compute(ctx, req); err == nil {
		return resp.Result
	} else {
		log.Printf("Error: %v", err)
		return 0
	}
}
