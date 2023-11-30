package main

import "testing"

func TestHello(t *testing.T) {
	t.Log("hello world")
	t.Log("hello world")

}

func TestHelloRedis(t *testing.T) {
	helloRedis()
	t.Log("hello")
}
