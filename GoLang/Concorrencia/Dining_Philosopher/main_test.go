package main

import (
	"testing"
	"time"
)

// func TestDine(testing *testing.T) {
// 	eatTime = 0 * time.Second
// 	sleepTime  = 0 * time.Second
// 	thinkTime = 0 * time.Second

// 	for i := 0; i < 10; i++{
// 		OrderFinished = []string{}
// 		dine()
// 		if len(OrderFinished) != len(philosophers) {
// 			testing.Errorf("Expected %d orders finished, got %d", len(philosophers), len(OrderFinished))
// 		}
// 	}
// }

func TestDineWithTiming(testing *testing.T){
	var tests = []struct{
		name string
		delay time.Duration
	}{
		{name: "Zero Delay", delay: 0 * time.Second},
		{name: "One Second Delay", delay: 1 * time.Second},
		{name: "Two Second Delay", delay: 2 * time.Second},
		{name: "Three Second Delay", delay: 3 * time.Second},
	}

	for i := 0; i<len(tests); i++{
		eatTime = tests[i].delay
		sleepTime  = tests[i].delay
		thinkTime = tests[i].delay
		OrderFinished = []string{}
		dine()
		if len(OrderFinished) != len(philosophers) {
			testing.Errorf("Expected %d orders finished, got %d", len(philosophers), len(OrderFinished))
		}
	}
}