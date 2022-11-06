package use_cases

import "agedito/udemy/rest_api_jwt/repository"

type UseCases struct {
	Repo repository.Repository
}

func New(repo repository.Repository) UseCases {
	return UseCases{repo}
}
