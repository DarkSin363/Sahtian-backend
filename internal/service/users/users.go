package users

import (
	"context"
	"time"

	"github.com/BigDwarf/sahtian/internal/model"
	"github.com/patrickmn/go-cache"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	unlockProfilePropertyCost  int64         = 120
	referralPremiumReward      int64         = 2000
	referralReward             int64         = 1000
	sahtianPerVote             int64         = 100
	sahtianPerVoteSelf         int64         = 20
	initialClaimableVotes      int64         = 0
	initialRemainingVotes      int64         = 10
	endFarmClaimableVotes      int64         = 10
	farmingPeriod              time.Duration = 8 * time.Hour //8 * time.Hour
	initialStreakDays          int64         = 1
	firstLevelsahtiansPercent  float64       = 0.03
	secondLevelsahtiansPercent float64       = 0.015
)

var streaksahtians = map[int64]int64{
	1: 100,
	2: 200,
	3: 400,
	4: 800,
	5: 1600,
	6: 3200,
	7: 6400,
}

type Repository interface {
	GetUser(ctx context.Context, id int64) (*model.User, error)
	UpsertDefaultUserData(ctx context.Context, user *model.User) (bool, error)
	SetDisplayName(ctx context.Context, userId int64, displayName string) error
	SetAvatarURL(ctx context.Context, userId int64, url string) error
}

type Service struct {
	client     *mongo.Client
	rep        Repository
	bucketName string
	cache      *cache.Cache
}

func NewUsersService(rep Repository, session *mongo.Client, bucket string) *Service {
	return &Service{
		rep:        rep,
		client:     session,
		bucketName: bucket,
		cache:      cache.New(10*time.Second, 10*time.Second),
	}
}

func (s *Service) SetAvatarURL(ctx context.Context, userId int64, url string) error {
	return s.rep.SetAvatarURL(ctx, userId, url)
}

func (s *Service) SetDisplayName(ctx context.Context, userId int64, displayName string) error {
	return s.rep.SetDisplayName(ctx, userId, displayName)
}

func (s *Service) GetExistingUser(ctx context.Context, id, requestedId int64) (*model.User, error) {
	user, err := s.rep.GetUser(ctx, requestedId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
