package pubsubapi

import (
	"github.com/garyburd/redigo/redis"
)

// Client for pub sub
type Client struct {
	pool *redis.Pool
	conn redis.PubSubConn
}

// NewPubSubClient returns a client which allows pubsub with Redis
func NewPubSubClient(addr string) (*Client, error) {
	var redispool *redis.Pool
	redispool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}

	return &Client{pool: redispool}, nil
}

// Publish to a channel using key as channel name
func (c *Client) Publish(key string, value string) error {
	conn := c.pool.Get()
	_, err := conn.Do("PUBLISH", key, value)

	return err
}

// Subscribe to a channel key as channel name
func (c *Client) Subscribe(key string) error {
	rc := c.pool.Get()
	c.conn = redis.PubSubConn{Conn: rc}
	return c.conn.PSubscribe(key)
}

// GetNextMessage reads the next message from redis channel
func (c *Client) GetNextMessage() ([]byte, error) {
	for {
		switch v := c.conn.Receive().(type) {
		case redis.PMessage:
			return v.Data, nil
		case redis.Error:
			return nil, v
		}
	}
}
