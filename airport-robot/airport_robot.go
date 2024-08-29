package airportrobot

type Language interface {
	SayHello() string
}

type Italian struct {
}

func (s Italian) SayHello() string {
	return "I can speak Italian: Ciao "
}

type Portuguese struct {
}

func (p Portuguese) SayHello() string {
	return "I can speak Portuguese: Olá "
}

func SayHello(name string, lang Language) string {
	return lang.SayHello() + name + "!"
}

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
