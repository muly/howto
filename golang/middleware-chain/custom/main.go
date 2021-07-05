package main

import "fmt"

type Service interface {
	call()
}

type service struct {
	name string
}

func (s service) call() {
	fmt.Println("calling ... ", s.name)
}

////////
type Middleware func(s Service) Service

type loggingMiddleware struct {
	name string
	next Service
}

func (m loggingMiddleware) call() {
	fmt.Println("calling ... ", m.name)
	m.next.call()
}

func main() {
	var svc = NewService("service-abc")
	{
		svc = NewLoggingMiddleware("logging-abc")(svc)
	}
	handler := makeHandler(svc)

	handler()
}

func NewService(name string) Service {
	return service{
		name: name,
	}
}

func NewLoggingMiddleware(name string) Middleware {
	return func(s Service) Service {
		return loggingMiddleware{
			name: name,
			next: s,
		}
	}
}

func makeHandler(s Service) func() {
	return func() { s.call() }
}
