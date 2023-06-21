package usecase

type AuthUseCase interface {
	Authenticate(email string, password string) bool
	GenerateToken(email string, status string) (string, error)
}
