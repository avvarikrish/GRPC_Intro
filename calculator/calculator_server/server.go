package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"time"

	"github.com/avvarikrish/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum method was invoked with %v\n", req)
	n1 := req.GetFirstNumber()
	n2 := req.GetSecondNumber()
	res := &calculatorpb.SumResponse{
		Sum: n1 + n2,
	}
	return res, nil
}

func (*server) PrimeNumber(req *calculatorpb.PrimeNumberRequest, stream calculatorpb.CalculatorService_PrimeNumberServer) error {
	fmt.Printf("Prime number function was invoked with request: %v\n", req)
	number := req.GetNumber()
	k := int32(2)
	for number > 1 {
		if number%k == 0 {
			res := &calculatorpb.PrimeNumberResponse{
				Number: k,
			}
			number = number / k
			stream.Send(res)
			time.Sleep(1000 * time.Millisecond)
		} else {
			k += 1
		}
	}
	return nil
}

func (*server) Average(stream calculatorpb.CalculatorService_AverageServer) error {
	fmt.Printf("Average function was invoked with streaming request")
	count := 0.0
	sum := 0.0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// stream is done
			return stream.SendAndClose(&calculatorpb.AverageResponse{
				Average: sum / count,
			})
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v", err)
		}
		count++
		sum += float64(req.GetNumber())
	}
}

func (*server) FindMax(stream calculatorpb.CalculatorService_FindMaxServer) error {
	fmt.Printf("FindMax function was invoked with streaming request\n")
	max := int32(math.MinInt32)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		curr := req.GetNumber()
		if curr > max {
			max = curr
			sendErr := stream.Send(&calculatorpb.FindMaxResponse{
				Number: max,
			})
			if sendErr != nil {
				log.Fatalf("Error while sending data to client: %v", sendErr)
				return sendErr
			}
		}
	}
}

func main() {
	fmt.Println("Sup")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
