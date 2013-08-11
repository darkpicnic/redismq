package redismq

import (
	"github.com/adeven/redis"
)

//Consumers are Watchers that have writing commands
type Consumer struct {
	Broker
}

func (self *Queue) AddConsumer(name string) (c *Consumer, err error) {
	c = &Consumer{Broker{Name: name, Queue: self}}
	//check uniqueness and start heartbeat
	self.redisClient.SAdd(self.WorkerKey(), name)
	return c, nil
}

func (self *Consumer) parseRedisAnswer(answer *redis.StringReq) (*Package, error) {
	if answer.Err() != nil {
		return nil, answer.Err()
	}
	p, err := UnmarshalPackage(answer.Val(), self.Queue, self)
	if err != nil {
		return nil, err
	}
	return p, nil
}