package match

var (
	Interests = []string{"movie", "music", "sport",
		"game", "travel", "food", "book",
		"art", "tech", "science", "animal",
	}
)

// MatchService is the interface that provides match methods.
type Service interface {
	GenerateMatch(userId string) (string, error)
	GenerateMatches(userId string, count int) ([]string, error)
	GetTodayMatches(userId string) ([]string, error)
	Like(userId string, likedId string) error
}

type service struct {
	repo Repository
}

// NewService creates a match service with necessary dependencies.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GenerateMatch(userId string) (string, error) {
	panic("not implemented") // TODO: Implement
}

func (s *service) GenerateMatches(userId string, count int) ([]string, error) {
	panic("not implemented") // TODO: Implement
}

func (s *service) GetTodayMatches(userId string) ([]string, error) {

	panic("not implemented") // TODO: Implement
}

func (s *service) Like(userId string, likedId string) error {
	panic("not implemented") // TODO: Implement
}
