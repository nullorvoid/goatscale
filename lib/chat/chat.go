package chat

// Provider interface that allow subscribing and publishing to a channel
type Provider interface {
	Publish(string, string) error
	Subscribe(string) error
	GetNextMessage() ([]byte, error)
}

// Client for chat module that provides functionality to send and receive messages
type Client struct {
	provider Provider
}

// NewChatClient returns a client that allows for sending and receiving messages
func NewChatClient(provider Provider) (*Client, error) {
	err := provider.Subscribe("message")

	if err != nil {
		return nil, err
	}

	return &Client{provider: provider}, nil
}

// SendMessage to all clients connected to the chat application
func (c *Client) SendMessage(msg string) error {
	return c.provider.Publish("message", msg)
}

// GetNextMessage gets the next message from the provider
func (c *Client) GetNextMessage() (string, error) {
	msg, err := c.provider.GetNextMessage()
	if err != nil {
		return "", err
	}

	// Convert byte array to string
	return string(msg), nil
}
