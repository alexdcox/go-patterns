package main

import "time"

type message struct {
	content string
	offset time.Duration
}

type replayQueue struct {
	first    *time.Time
	messages []*message
}

func (r *replayQueue) Send(content string) error {
	if r.first == nil {
		now := time.Now()
		r.first = &now
	}

	r.messages = append(r.messages, &message{
		content: content,
		offset: time.Now().Sub(*r.first),
	})

	return nil
}

func (r *replayQueue) Replay() <-chan *message {
	replayChannel := make(chan *message)

	tickerStart := time.Now()
	ticker := time.NewTicker(time.Nanosecond)
	indexCurrent := 0
	indexEnd := len(r.messages) - 1

	go func() {
		for {
			if indexCurrent == indexEnd {
				close(replayChannel)
				return
			} else {
				<-ticker.C
				currentOffset := time.Now().Sub(tickerStart)
				currentMessage := r.messages[indexCurrent]
				if currentOffset >= currentMessage.offset {
					replayChannel <- currentMessage
					indexCurrent++
				}
			}
		}
	}()

	return replayChannel
}
