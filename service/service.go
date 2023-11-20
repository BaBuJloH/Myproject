package service

type Service struct {
	Repositories Repositories
}

type Repositories struct {
	Users *users_repo.Repo
}

func New(repos Repositories) *Service {
	return &Service{
		Repositories: repos,
	}
}
