package services

// HelloService provides a simple greeting.
// In a real application, this would interact with the database,
// other APIs, etc.
type HelloService struct {
	// Dependencies like a database connection would go here.
}

// NewHelloService creates a new instance of HelloService.
func NewHelloService() *HelloService {
	return &HelloService{}
}

// GetGreeting returns a simple greeting message.
// This is where your business logic would live.
func (s *HelloService) GetGreeting(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello from the service layer, " + name + "!"
}
