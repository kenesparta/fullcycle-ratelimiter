package entity

//go:generate mockgen -source repository.go -destination mock/repository_mock.go -package mock
type APITokenRepository interface {
	Save(token APIToken) error
}

type IPRepository interface {
	Save(ip IP) error
}
