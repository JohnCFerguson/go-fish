package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run sentry_error_reporter.go <error_type> <message>")
		fmt.Println("Error types: exception, message, context")
		os.Exit(1)
	}

	errorType := os.Args[1]
	message := os.Args[2]

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "http://d9c5a7d2c52135ff1ce204ba9f4d091e@192.168.68.61:9000/2",
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	switch errorType {
	case "exception":
		sentry.CaptureException(fmt.Errorf(message))
	case "message":
		sentry.CaptureMessage(message)
	case "context":
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetTag("category", "errors")
			scope.SetLevel(sentry.LevelError)
			sentry.CaptureException(fmt.Errorf(message))
		})
	default:
		fmt.Println("Invalid error type. Use 'exception', 'message', or 'context'.")
		os.Exit(1)
	}

	fmt.Println("Error sent to Sentry")
}
