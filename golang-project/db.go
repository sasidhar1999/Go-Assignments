package main

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type DB struct {
	conn *pgx.Conn
}

// constructor function
func NewDB(dsn string) (*DB, error) {
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return &DB{conn: conn}, nil
}

func (db *DB) Close() {
	db.conn.Close(context.Background())
}

// Create table
func (db *DB) CreateTable() error {
	_, err := db.conn.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS subscriptions (
		id SERIAL PRIMARY KEY,
		user_id VARCHAR(255) NOT NULL,
		topic VARCHAR(255) NOT NULL,
		email VARCHAR(255),
		sms VARCHAR(20),
		push_notifications BOOLEAN,
		UNIQUE(user_id, topic)
	);`)
	return err
}

func (db *DB) Subscribe(sub Subscription) error {
	for _, topic := range sub.Topics {
		_, err := db.conn.Exec(context.Background(),
			"INSERT INTO subscriptions (user_id, topic, email, sms, push_notifications) VALUES ($1, $2, $3, $4, $5)",
			sub.UserID, topic, sub.NotificationChannels.Email, sub.NotificationChannels.SMS, sub.NotificationChannels.PushNotifications)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *DB) Unsubscribe(userID string, topics []string) error {
	for _, topic := range topics {
		_, err := db.conn.Exec(context.Background(),
			"DELETE FROM subscriptions WHERE user_id = $1 AND topic = $2", userID, topic)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *DB) FetchSubscriptions(userID string) ([]Subscription, error) {
	rows, err := db.conn.Query(context.Background(), "SELECT topic, email, sms, push_notifications FROM subscriptions WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []Subscription
	for rows.Next() {
		var sub Subscription
		var channels NotificationChannels
		if err := rows.Scan(&sub.Topics, &channels.Email, &channels.SMS, &channels.PushNotifications); err != nil {
			return nil, err
		}
		sub.UserID = userID
		sub.NotificationChannels = channels
		subscriptions = append(subscriptions, sub)
	}

	return subscriptions, nil
}
