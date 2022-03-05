package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Enable metrics middleware
	p := prometheus.NewPrometheus("echo", urlSkipper)
	p.Use(e)

	// Enable tracing middleware
	c := jaegertracing.New(e, urlSkipper)
	defer c.Close()

	e.GET("/", func(c echo.Context) error {
		// Wrap slowFunc on a new span to trace it's execution passing the function arguments
		jaegertracing.TraceFunction(c, slowFunc, "Test String")
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/child-span", func(c echo.Context) error {
		// Do something before creating the child span
		time.Sleep(40 * time.Millisecond)
		sp := jaegertracing.CreateChildSpan(c, "Child span for additional processing")
		defer sp.Finish()
		sp.LogEvent("Test log")
		sp.SetBaggageItem("Test baggage", "baggage")
		sp.SetTag("Test tag", "New Tag")
		time.Sleep(100 * time.Millisecond)
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

// urlSkipper middleware ignores metrics on some route
func urlSkipper(c echo.Context) bool {
	if strings.HasPrefix(c.Path(), "/testurl") {
		return true
	}
	if strings.HasPrefix(c.Path(), "/testurl2") {
		return true
	}
	return false
}

// A function to be wrapped. No need to change it's arguments due to tracing
func slowFunc(s string) {
	time.Sleep(200 * time.Millisecond)
	return
}
