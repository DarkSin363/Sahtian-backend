package repository

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	in_errors "github.com/BigDwarf/sahtian/internal/errors"
	"github.com/BigDwarf/sahtian/internal/model"
)

type UsersRepository struct {
	collection *mongo.Collection
}

func NewUsersRepository(client *mongo.Collection) *UsersRepository {
	return &UsersRepository{
		collection: client,
	}
}

func (u *UsersRepository) GetUser(ctx context.Context, id int64) (*model.User, error) {
	var res model.User

	err := u.collection.FindOne(ctx, bson.D{{"id", id}}).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, in_errors.ErrUserNotFound
		}

		return nil, err
	}
	return &res, nil
}

func (u *UsersRepository) UpsertDefaultUserData(ctx context.Context, user *model.User) (bool, error) {
	filter := bson.D{{"id", user.ID}}

	res, err := u.collection.UpdateOne(ctx, filter, bson.D{{
		"$set", bson.D{
			{"id", user.ID},
			{"added_to_attachment_menu", user.AddedToAttachmentMenu},
			{"allows_write_to_pm", user.AllowsWriteToPm},
			{"first_name", user.FirstName},
			{"last_name", user.LastName},
			{"is_bot", user.IsBot},
			{"is_premium", user.IsPremium},
			{"is_admin", false},
			{"username", user.Username},
			{"params.display_name", user.Username},
			{"avatar_url", user.AvatarURL},
			{"visible", true},
			{"language_code", user.LanguageCode},
		},
	}}, options.Update().SetUpsert(true))
	if err != nil {
		return false, err
	}

	isNewUser := res.UpsertedCount == 1

	return isNewUser, nil
}

func (u *UsersRepository) SetDisplayName(ctx context.Context, userId int64, displayName string) error {
	filter := bson.D{{"id", userId}}

	_, err := u.collection.UpdateOne(ctx, filter, bson.D{{
		"$set", bson.D{
			{"params.display_name", displayName},
		},
	}}, options.Update())
	if err != nil {
		return err
	}

	return nil
}

func (u *UsersRepository) SetAvatarURL(ctx context.Context, userId int64, url string) error {
	filter := bson.D{{"id", userId}}

	_, err := u.collection.UpdateOne(ctx, filter, bson.D{{
		"$set", bson.D{
			{"avatar_url", url},
		},
	}}, options.Update())
	if err != nil {
		return err
	}

	return nil
}
