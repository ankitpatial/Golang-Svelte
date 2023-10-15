/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package account

import (
	"context"
	"fmt"
	"roofix/ent"
	"roofix/ent/notifysetting"
	"roofix/ent/user"
	"roofix/pkg/util/crypt"
	"roofix/pkg/util/log"
)

func ToNotifyEmails(ctx context.Context, topic string) []string {
	users := byTopic(ctx, topic)

	var to []string
	for _, u := range users {
		to = append(to, u.Email)
	}
	return to
}

func ToNotify(ctx context.Context, topic string) []*EmailAddress {
	users := byTopic(ctx, topic)

	var to []*EmailAddress
	for _, u := range users {
		name := u.FirstName + " " + u.LastName
		to = append(to, &EmailAddress{
			Name:  name,
			Email: fmt.Sprintf("%s<%s>", name, u.Email),
		})
	}
	return to
}

// byTopic user's name ane mail info
func byTopic(ctx context.Context, topic string) []*ent.User {
	db := ent.GetClient()
	defer db.CloseClient()

	id := crypt.MD5Hash(topic)
	// let admin know about the process is done
	users, err := db.User.Query().
		Where(user.HasNotifyWith(notifysetting.TopicID(id), notifysetting.ReceiveEmail(true))).
		Select(user.FieldID, user.FieldEmail, user.FieldFirstName, user.FieldLastName).
		All(ctx)

	if err != nil {
		log.Error(err)
		return nil
	}

	return users
}
