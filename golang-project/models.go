package main

type Subscription struct {
	UserID               string               `json:"user_id"`
	Topics               []string             `json:"topics"`
	NotificationChannels NotificationChannels `json:"notification_channels"`
}

type NotificationChannels struct {
	Email             string `json:"email"`
	SMS               string `json:"sms"`
	PushNotifications bool   `json:"push_notifications"`
}

type Notification struct {
	Topic   string        `json:"topic"`
	Event   EventDetail   `json:"event"`
	Message MessageDetail `json:"message"`
}

type EventDetail struct {
	EventID   string      `json:"event_id"`
	Timestamp string      `json:"timestamp"`
	Details   UserDetails `json:"details"`
}

type UserDetails struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type MessageDetail struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type SubscriptionResponse struct {
	UserID        string         `json:"user_id"`
	Subscriptions []Subscription `json:"subscriptions"`
}
