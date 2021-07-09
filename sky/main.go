package main

import (
	"context"
	"log"
	"time"

	"github.com/SkyAPM/go2sky"
	shttp "github.com/SkyAPM/go2sky/plugins/http"
	// "github.com/SkyAPM/go2sky/propagation"
	"github.com/SkyAPM/go2sky/reporter"

	"net/http"

	"bufio"
	"fmt"
)

func main() {
	r, err := reporter.NewGRPCReporter("localhost:11800")
	if err != nil {
		log.Fatalf("new reporter error %v\n", err)
	}
	defer r.Close()

	tracer, err := go2sky.NewTracer("example", go2sky.WithReporter(r), go2sky.WithSampler(1.1))
	if err != nil {
		log.Fatalf("create new tracer error %v\n", err)
	}

	span, ctx, err := tracer.CreateLocalSpan(context.Background())
	if err != nil {
		log.Fatalf("create new local span error %v\n", err)
	}

	client, err := shttp.NewClient(tracer)
	if err != nil {
		log.Fatalf("create client error %v \n", err)
	}

	span.SetOperationName("TOP GUN")
	span.Log(time.Now(), "an event for the span")

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8081/greeting/hello", nil)
	if err != nil {
		log.Fatalf("failed to create a request")
	}

	// Inject context into HTTP request header `sw8`
	// subSpan, err := tracer.CreateExitSpan(req.Context(), "NET INVOKE", "mac:8080", func(header string) error {
	//   req.Header.Set(propagation.Header, header)
	//   return nil
	// })

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	// subSpan.End()
	span.End()
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
func one() {
	// Use gRPC reporter for production
	// r, err := reporter.NewLogReporter()
	r, err := reporter.NewGRPCReporter("localhost:11800")
	if err != nil {
		log.Fatalf("new reporter error %v \n", err)
	}
	defer r.Close()
	tracer, err := go2sky.NewTracer("example", go2sky.WithReporter(r))
	if err != nil {
		log.Fatalf("create tracer error %v \n", err)
	}
	// This for test
	span, ctx, err := tracer.CreateLocalSpan(context.Background())
	if err != nil {
		log.Fatalf("create new local span error %v \n", err)
	}
	span.SetOperationName("invoke data")
	span.Tag("kind", "outer")
	time.Sleep(time.Second)
	subSpan, _, err := tracer.CreateLocalSpan(ctx)
	if err != nil {
		log.Fatalf("create new sub local span error %v \n", err)
	}
	subSpan.SetOperationName("invoke inner")
	subSpan.Log(time.Now(), "inner", "this is right")
	time.Sleep(time.Second)
	subSpan.End()
	time.Sleep(500 * time.Millisecond)
	span.End()
	time.Sleep(time.Second)
	// Output:

}
