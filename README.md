# channelify

[![GitHub stars](https://img.shields.io/github/stars/ddelizia/channelify.svg?style=social&label=Star&maxAge=2592000)](https://GitHub.com/ddelizia/channelify/stargazers/) 
[![GitHub forks](https://img.shields.io/github/forks/ddelizia/channelify.svg?style=social&label=Fork&maxAge=2592000)](https://GitHub.com/ddelizia/channelify/network/) 

[![Test Actions Status](https://github.com/ddelizia/channelify/workflows/Test/badge.svg)](https://github.com/ddelizia/channelify/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/ddelizia/channelify)](https://goreportcard.com/report/github.com/ddelizia/channelify)
[![Go Doc](https://godoc.org/github.com/channelify/channelify?status.svg)](https://godoc.org/github.com/ddelizia/channelify) 
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://lbesson.mit-license.org/)


This library helps you to transform any function into a function that returns a any type to a function that return such types within a channel. This is useful to run in parallel multiple functions and have control on the returned values.

Channelify uses go routines to parallelize the execution of the functions. 

The idea comes from Javascript Promisify utility that transforms a callback into a promise.

## Installation

```
go get github.com/ddelizia/channelify
```


## Usage example

Here an example of transforming a simple function in channel so you can execute multiple functions in parallel:

```go
fn := func () string {
    time.Sleep(time.Second * 3)
    return "hello"
}

ch1 := Channelify(fn)
ch2 := Channelify(fn)
chV1 := ch1.(func () chan string)()
chV2 := ch2.(func () chan string)()

v1, v2 := <- chV1, <- chV2
```

If your functions returns multiple values you can use as follow:

```go
fn1 := func (hello string) (string, error)  {
    time.Sleep(time.Second * 2)
    fmt.Println(hello)
    return hello, nil
}

fn2 := func (hello string) (string, error)  {
    time.Sleep(time.Second * 3)
    fmt.Println(hello)
    return hello, nil
}

ch1 := Channelify(fn1)
ch2 := Channelify(fn2)
chV1, chE1 := ch1.(func (string) (chan string, chan error))("hello1")
chV2, chE2 := ch2.(func (string) (chan string, chan error))("hello2")

v1, e1, v2, e2 := <- chV1, <- chE1, <- chV2, <- chE2

fmt.Print(v1, e1, v2, e2)
```

# Contributing

 * Fork it
 * Create your feature branch (`git checkout -b my-new-feature`)
 * Commit your changes (`git commit -am 'Add some feature'`)
 * Push to the branch (`git push origin my-new-feature`)
 * Create new Pull Request


