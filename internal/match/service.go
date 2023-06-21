package match

import (
	"errors"
	"math/rand"

	"github.com/krissukoco/deall-technical-test-go/internal/models"
	"github.com/krissukoco/deall-technical-test-go/internal/subscription"
	"github.com/krissukoco/deall-technical-test-go/internal/user"
)

var (
	Interests = []string{"movie", "music", "sport",
		"game", "travel", "food", "book",
		"art", "tech", "science", "animal",
	}
	ErrMaxMatchPerDay    = errors.New("max match per day reached")
	ErrNoMatchAvailable  = errors.New("no match available")
	ErrMatchAlreadyLiked = errors.New("match already liked")
)

const (
	MaxMatchPerDayUnsubscribed = 10
)

// MatchService is the interface that provides match methods.
type Service interface {
	GenerateMatch(userId string) (*MatchData, error)
	// GenerateMatches(userId string, count int) ([]string, error)
	GetTodayMatches(userId string, limit int) ([]string, error)
	Like(userId string, id int64) (*models.Match, error)
}

type service struct {
	subsService subscription.Service
	userService user.Service
	repo        Repository
}

// NewService creates a match service with necessary dependencies.
func NewService(repo Repository, subService subscription.Service, userService user.Service) Service {
	return &service{
		subsService: subService,
		userService: userService,
		repo:        repo,
	}
}

type MatchData struct {
	Id             int64    `json:"id"`
	UserId         string   `json:"user_id"`
	Name           string   `json:"name"`
	ProfilePicture string   `json:"profile_picture"`
	Images         []string `json:"images"`
	CreatedAt      int64    `json:"created_at"`
}

// Generate match for a user
// Returns a user id
func (s *service) GenerateMatch(userId string) (*MatchData, error) {
	// Get user
	user, err := s.userService.GetById(userId)
	if err != nil {
		return nil, err
	}
	gender := "male"
	if user.Gender == "male" {
		gender = "female"
	}

	// Get todays matches
	todayMatches, err := s.GetTodayMatches(userId, MaxMatchPerDayUnsubscribed)
	if err != nil {
		return nil, err
	}

	// Get user's subscription
	sub, err := s.subsService.Get(userId)
	if err != nil && len(todayMatches) >= MaxMatchPerDayUnsubscribed {
		// Not subscribed yet
		return nil, ErrMaxMatchPerDay
	}
	if len(todayMatches) >= MaxMatchPerDayUnsubscribed {
		if sub == nil {
			// Subscription is inactive
			return nil, ErrMaxMatchPerDay
		} else if !sub.IsActive() {
			return nil, ErrMaxMatchPerDay
		}
	}

	// Generate match
	// log.Printf("gender: %s, limit: %d, matches today: %v", gender, 10, todayMatches)
	availableUsers, err := s.userService.FindByGenderExcludeIds(gender, 10, todayMatches)
	if err != nil {
		return nil, err
	}
	if len(availableUsers) == 0 {
		return nil, ErrNoMatchAvailable
	}
	// Choose random user id
	// TODO: can be improved by using a better algorithm
	matchedUser := availableUsers[rand.Intn(len(availableUsers))]

	match, err := s.createMatch(userId, matchedUser.Id)
	if err != nil {
		return nil, err
	}

	return &MatchData{
		Id:             match.Id,
		UserId:         matchedUser.Id,
		Name:           matchedUser.Name,
		ProfilePicture: matchedUser.ProfilePicture,
		// TODO: Get images from user
		Images:    make([]string, 0),
		CreatedAt: match.CreatedAt,
	}, nil
}

func (s *service) createMatch(userId, matcheeId string) (*models.Match, error) {
	return s.repo.Create(userId, matcheeId)
}

func (s *service) GenerateMatches(userId string, count int) ([]string, error) {
	panic("not implemented") // TODO: Implement
}

func (s *service) GetTodayMatches(userId string, limit int) ([]string, error) {
	userIds := make([]string, 0)
	matches, err := s.repo.GetLast24Hours(userId, limit)
	if err != nil {
		return nil, err
	}
	for _, match := range matches {
		userIds = append(userIds, match.MatcheeId)
	}
	return userIds, nil
}

func (s *service) Like(userId string, id int64) (*models.Match, error) {
	m, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	if m.UserId != userId {
		return nil, ErrMatchNotFound
	}
	if m.Liked {
		return nil, ErrMatchAlreadyLiked
	}

	return s.repo.Like(id)
}
