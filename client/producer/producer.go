package producer

import "git.nm.flipkart.com/git/infra/kafka-lite/client/core"

type Producer struct {
	accumulator *Accumulator
}

func NewProducer(accumulator *Accumulator) *Producer {
	return &Producer{accumulator: accumulator}
}

func (this *Producer) send(record *core.ProducerRecord) *core.RecordMetadata {
	this.accumulator.append(record)
}
