package use_cases

func (cases *UseCases) GetOwnProfile(email string) (string, error) {
	_, exists, repoErr := cases.Repo.FindUser(email)
	if !exists {
		return "", repoErr
	}

	return email, nil
}
