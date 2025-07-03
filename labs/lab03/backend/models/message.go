package models

import "time"

// Message represents a chat message
type Message struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

// CreateMessageRequest represents the request to create a new message
type CreateMessageRequest struct {
	Username string `json:"username" validate:"required"`
    Content  string `json:"content" validate:"required"`
}

// UpdateMessageRequest represents the request to update a message
type UpdateMessageRequest struct {
	Content  string `json:"content" validate:"required"`
}

// HTTPStatusResponse represents the response for HTTP status code endpoint
type HTTPStatusResponse struct {
	// TODO: Add StatusCode field of type int with json tag "status_code"
	// TODO: Add ImageURL field of type string with json tag "image_url"
	// TODO: Add Description field of type string with json tag "description"
}

// APIResponse represents a generic API response
type APIResponse struct {
	// TODO: Add Success field of type bool with json tag "success"
	// TODO: Add Data field of type interface{} with json tag "data,omitempty"
	// TODO: Add Error field of type string with json tag "error,omitempty"
}

// NewMessage creates a new message with the current timestamp
func NewMessage(id int, username, content string) *Message {
	// TODO: Return a new Message instance with provided parameters and current timestamp
	return nil
}

// Validate checks if the create message request is valid
func (r *CreateMessageRequest) Validate() error {
	// TODO: Implement validation logic
	// Check if Username is not empty
	// Check if Content is not empty
	// Return appropriate error messages
	return nil
}

// Validate checks if the update message request is valid
func (r *UpdateMessageRequest) Validate() error {
	// TODO: Implement validation logic
	// Check if Content is not empty
	// Return appropriate error messages
	return nil
}
