package db

import (
	"context"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Conversation struct {
	_id        primitive.ObjectID
	Members    []string        // slice of usernames
	NewMessage map[string]bool // username -> true/false    true when an unseen message exists, false when all messages fetched
}

func (conv *Conversation) Create(context context.Context, sender string, recipient string) (cid primitive.ObjectID, err error) {
	NewMessageMap := make(map[string]bool)
	NewMessageMap[sender] = false
	NewMessageMap[recipient] = false
	conv.NewMessage = NewMessageMap
	result, err := db.Collection("conversations").InsertOne(context, conv)
	if err != nil {
		glog.Error(err)
		return
	}
	cid = result.InsertedID.(primitive.ObjectID)
	return
}