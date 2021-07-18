package repository

type Repository interface {
	Start(out string) (string, error)
}
