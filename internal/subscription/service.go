package subscription

import (
	"errors"

	"github.com/krissukoco/deall-technical-test-go/internal/models"
)

const (
	SubscriptionMonthly string = "monthly"
	SubscriptionYearly  string = "yearly"
	OneMonthInt         int64  = 24 * 30 * 60 * 60 * 1000
	OneYearInt          int64  = 24 * 365 * 60 * 60 * 1000
)

var (
	ErrNoSubscription          = errors.New("no subscription")
	ErrSubscriptionTypeInvalid = errors.New("invalid subscription type")
	ErrAlreadySubscribed       = errors.New("already subscribed")
)

type subscriptionPackage struct {
	Type  string `json:"type"`
	Price int64  `json:"price"`
	Title string `json:"title"`
}

type Service interface {
	Get(userId string) (*models.Subscription, error)
	Packages() []*subscriptionPackage
	Buy(userId string, subType string) (*models.Subscription, error)
	Renew(userId string, subType string) (*models.Subscription, error)
}

type service struct {
	repo Repository
}

func validSubType(subType string) bool {
	return subType == SubscriptionMonthly || subType == SubscriptionYearly
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Get(userId string) (*models.Subscription, error) {
	return s.repo.Get(userId)
}

func (s *service) Packages() []*subscriptionPackage {
	return []*subscriptionPackage{
		{"monthly", 100000, "One Month Subscription for Unlimited Swipes"},
		{"yearly", 1000000, "One Year Subscription for Unlimited Swipes"},
	}
}

func (s *service) Buy(userId string, subType string) (*models.Subscription, error) {
	if !validSubType(subType) {
		return nil, ErrSubscriptionTypeInvalid
	}
	sub, err := s.repo.Get(userId)
	if err == nil {
		if sub.IsActive() {
			return nil, ErrAlreadySubscribed
		}
		return s.Renew(userId, subType)
	}
	if !errors.Is(err, ErrNoSubscription) {
		return nil, err
	}
	add := OneMonthInt
	if subType == SubscriptionYearly {
		add = OneYearInt
	}
	return s.repo.Create(userId, add)
}

func (s *service) Renew(userId string, subType string) (*models.Subscription, error) {
	if !validSubType(subType) {
		return nil, ErrSubscriptionTypeInvalid
	}
	add := OneMonthInt
	if subType == SubscriptionYearly {
		add = OneYearInt
	}
	return s.repo.Renew(userId, add)
}
