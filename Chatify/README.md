# Chatify - Real-time Chat Application

Chatify is a modern real-time chat application built with Go, featuring WebSocket communication for instant messaging.

## Features

### Authentication
- User registration with email verification
- OTP-based login system
- JWT token authentication

### Messaging
- Real-time messaging using WebSockets
- Direct messaging between users
- Group chat capabilities
- Message history storage in MongoDB

### Enhanced WebSocket Features
- Typing indicators
- Read receipts
- Online/offline status updates
- Direct messaging with targeted delivery

## WebSocket Protocol

### Connection
Connect to the WebSocket server at `/ws` with a valid JWT token:

```
ws://localhost:8003/ws?token=your_jwt_token
```

Alternatively, you can provide the token in the Authorization header:

```
Authorization: Bearer your_jwt_token
```

### Message Format

All WebSocket messages use the following JSON format:

```json
{
  "type": "message",        // message, typing, read_receipt, status
  "sender_id": "user123",   // ID of the sender
  "receiver_id": "user456", // ID of the receiver (for direct messages)
  "content": "Hello!",      // Message content
  "message_id": "msg123",   // Message ID (for updates/deletes/receipts)
  "timestamp": "2023-01-01T12:00:00Z"
}
```

### Message Types

#### Regular Message
```json
{
  "type": "message",
  "receiver_id": "user456",
  "content": "Hello, how are you?"
}
```

#### Typing Indicator
```json
{
  "type": "typing",
  "receiver_id": "user456",
  "content": "true"  // or "false" when stopped typing
}
```

#### Read Receipt
```json
{
  "type": "read_receipt",
  "receiver_id": "user456",
  "message_id": "msg123"
}
```

### Status Updates

The server automatically sends status updates when users connect or disconnect:

```json
{
  "type": "status",
  "user_id": "user123",
  "status": "online",  // or "offline"
  "timestamp": "2023-01-01T12:00:00Z"
}
```

## Running the Application

### Prerequisites
- Go 1.16+
- MongoDB
- Firebase account (for authentication)

### Environment Variables
Create a `.env` file with the following variables:

```
MONGO_URI=your_mongodb_connection_string
JWT_SECRET=your_jwt_secret
FIREBASE_KEY_PATH=path_to_firebase_credentials.json
```

### Starting the Servers

1. Start the HTTP API server:
```
go run cm/main.go
```

2. Start the WebSocket server:
```
go run cmd/main.go
```

## Architecture

The application follows a modular architecture:

- **API Layer**: RESTful endpoints for authentication and message operations
- **WebSocket Layer**: Real-time communication
- **Service Layer**: Business logic
- **Database Layer**: Data persistence with MongoDB and Firebase

## WebSocket Implementation Details

### Hub
The central manager for all WebSocket connections. It handles:
- Client registration and unregistration
- Broadcasting messages to all clients
- Directing messages to specific clients
- Tracking online/offline status

### Client
Represents a connected user. Each client has:
- WebSocket connection
- User ID
- Message channel
- Typing status
- Last seen timestamp

### Message Routing
Messages are routed based on their type and recipient:
- Direct messages are sent only to the specified recipient
- Broadcast messages are sent to all connected clients
- System messages (status updates, etc.) are sent to relevant clients

## Security Considerations

- All WebSocket connections require JWT authentication
- Messages are validated before processing
- Rate limiting is implemented to prevent abuse
- Connection timeouts and pings maintain connection health