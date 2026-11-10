package users_postgres_repository

import "github.com/WaveSilent/golang-todoapp/internal/core/domain"

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func userDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))

	for k, v := range users {
		userDomains[k] = domain.NewUser(
			v.ID,
			v.Version,
			v.FullName,
			v.PhoneNumber,
		)
	}

	return userDomains
}
