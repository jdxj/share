package handler

import (
	"context"
	"errors"
	"fmt"
	"net"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/jdxj/logger"
	"github.com/jdxj/share/config"
	email "github.com/jdxj/share/email/proto"
)

func TestEmail_Send(t *testing.T) {
	config.Init("/home/jdxj/workspace/share/config.yaml")
	logger.NewPathMode(config.Log().Path, config.Mode())
	e := new(Email)

	req := &email.RequestEmail{
		Token:      "123",
		Subject:    "test",
		Recipients: []string{"985759262@qq.com"},
		Type:       1,
		Content:    []byte("hello world"),
	}
	e.Send(context.TODO(), req, &email.ResponseEmail{})
}

func TestInterface(t *testing.T) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		// no net connectivity maybe so fallback
		t.Fatalf("%s\n", err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	fmt.Printf("%s\n", localAddr.IP)
}

func deadloop() {
	for {
	}
}

func TestGoroutineScheduler(t *testing.T) {
	go deadloop()
	go deadloop()
	go deadloop()
	go deadloop()
	go deadloop()
	go deadloop()
	go deadloop()
	go deadloop()
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("I got scheduled!")
	}
}

func shutdownMaker(processTm int) func(time.Duration) error {
	return func(time.Duration) error { // 参数没有使用
		time.Sleep(time.Second * time.Duration(processTm))
		return nil
	}
}

func TestConcurrentShutdown(t *testing.T) {
	f1 := shutdownMaker(2)
	f2 := shutdownMaker(6)

	err := ConcurrentShutdown(10*time.Second, ShutdownerFunc(f1), ShutdownerFunc(f2))
	if err != nil {
		t.Errorf("want nil, actual: %s", err)
		return
	}

	err = ConcurrentShutdown(4*time.Second, ShutdownerFunc(f1), ShutdownerFunc(f2))
	if err == nil {
		t.Error("want timeout, actual nil")
		return
	}
}

func ConcurrentShutdown(waitTimeout time.Duration, shutdowners ...GracefullyShutdowner) error {
	c := make(chan struct{})

	go func() {
		var wg sync.WaitGroup
		for _, g := range shutdowners {
			wg.Add(1)
			go func(shutdowner GracefullyShutdowner) {
				defer wg.Done()
				shutdowner.Shutdown(waitTimeout)
			}(g)
		}
		wg.Wait()
		c <- struct{}{}
	}()

	timer := time.NewTimer(waitTimeout)
	defer timer.Stop()

	select {
	case <-c:
		return nil
	case <-timer.C:
		return errors.New("wait timeout")
	}
}

type ShutdownerFunc func(time.Duration) error

func (f ShutdownerFunc) Shutdown(waitTimeout time.Duration) error {
	return f(waitTimeout)
}

type GracefullyShutdowner interface {
	Shutdown(waitTimeout time.Duration) error
}

type MyError struct {
	e string
}

func (e *MyError) Error() string {
	return e.e
}

func TestErrAs(t *testing.T) {
	var err = &MyError{"my error type"}
	err1 := fmt.Errorf("wrap err1: %w", err)
	err2 := fmt.Errorf("wrap err2: %w", err1)
	var e *MyError
	if errors.As(err2, &e) {
		println("err is a variable of MyError type ")
		println(e == err)
		return
	}

	println("err is not a variable of the MyError type ")
}

type A struct {
}

func (a *A) MethodA() {

}

func TestReflectIndirect(t *testing.T) {
	n := 3
	nValue := reflect.ValueOf(&n)
	fmt.Printf("%v\n", nValue.CanSet())
	fmt.Printf("%v\n", nValue.Elem().CanSet())

	s := []int{1, 2, 3}
	sValue := reflect.ValueOf(s)
	sValue.Index(0).Set(reflect.ValueOf(4))
	fmt.Printf("%v\n", s)

	a := [...]int{1, 2, 3}
	fmt.Printf("a's addr: %p\n", &a)

}
