package channelify

import (
	"reflect"
	"testing"
	"time"
)

func TestChannelify_ShouldReturnCorrectType(t *testing.T) {
	fn := func() string {
		return "hello"
	}

	typeOfResutnFunc := reflect.TypeOf(func() chan string {
		return make(chan string)
	})
	ch1 := Channelify(fn)
	if reflect.TypeOf(ch1) != typeOfResutnFunc {
		t.Fail()
	}
}

func TestChannelify_ShouldReturnCorrectData(t *testing.T) {
	fn := func() string {
		time.Sleep(time.Second * 3)
		return "hello"
	}
	ch1 := Channelify(fn)
	chV1 := ch1.(func() chan string)()

	v1 := <-chV1

	if v1 != "hello" {
		t.Fail()
	}
}

func TestChannelify_ShouldRunInParallel(t *testing.T) {
	fn := func() string {
		time.Sleep(time.Second * 3)
		return "hello"
	}

	start := time.Now().UnixNano() / int64(time.Millisecond)
	ch1 := Channelify(fn)
	ch2 := Channelify(fn)
	chV1 := ch1.(func() chan string)()
	chV2 := ch2.(func() chan string)()

	v1, v2 := <-chV1, <-chV2

	end := time.Now().UnixNano() / int64(time.Millisecond)

	if v1 != "hello" || v2 != "hello" || (end-start) > 4000 {
		t.Fail()
	}
}
