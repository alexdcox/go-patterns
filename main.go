package main

import (
	"time"
	"fmt"
)

func main() {
	demoReplayQueue()
}

func demoReplayQueue() {
	rq := new(replayQueue)

	fmt.Println("sending messages into replay queue...")

	rq.Send("message 1")
	time.Sleep(time.Millisecond * 300)
	rq.Send("message 2")
	time.Sleep(time.Millisecond * 100)
	rq.Send("message 3")
	time.Sleep(time.Millisecond * 500)
	rq.Send("message 4")
	time.Sleep(time.Millisecond * 100)
	rq.Send("message 5")
	time.Sleep(time.Millisecond * 700)
	rq.Send("message 6")
	time.Sleep(time.Millisecond * 200)
	rq.Send("message 7")

	var replayStart *time.Time

	fmt.Println("replaying messages with the same interval...")

	for msg := range rq.Replay() {
		if replayStart == nil {
			now := time.Now()
			replayStart = &now
		}

		fmt.Printf(
			"content: %s, originalDelay: %v, replayDelay: %v\n",
			msg.content,
			msg.offset,
			time.Now().Sub(*replayStart),
		)
	}
}