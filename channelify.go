package channelify

import (
	"fmt"
	"reflect"
)

// Channelify receives whatever type of function and transforms it in the same function returning channels for each of the returned types.
func Channelify(fn interface{}) interface{} {

	funcValue := reflect.ValueOf(fn)
	funcType := reflect.TypeOf(fn)
	numIn := funcType.NumIn()
	numOut := funcType.NumOut()

	var ins []reflect.Type
	for i := 0; i < numIn; i++ {
		inV := funcType.In(i)
		ins = append(ins, inV)
	}

	var outs []reflect.Type
	for i := 0; i < numOut; i++ {
		outV := funcType.Out(i)
		outs = append(outs, reflect.ChanOf(reflect.BothDir, outV))
	}

	newFuncType := reflect.FuncOf(ins, outs, funcType.IsVariadic())

	makeFun := reflect.MakeFunc(newFuncType, func(incoming []reflect.Value) (outgoing []reflect.Value) {

		fmt.Println(len(incoming), len(outs))
		var channels []reflect.Value
		for _, o := range outs {
			channels = append(channels, reflect.MakeChan(o, 0))
		}

		go func() {
			for _, c := range channels {
				defer c.Close()
			}

			results := funcValue.Call(incoming)

			for i, result := range results {
				channels[i].Send(result)
			}

		}()

		return channels
	})

	return makeFun.Interface()
}
