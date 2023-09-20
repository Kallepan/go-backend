package repository

import "github.com/kallepan/go-backend/app/domain/dco"

type SystemRepository interface {
	GetVersion() (*dco.Version, error)
}

type SystemRepositoryImpl struct {
}

func (s SystemRepositoryImpl) GetVersion() (*dco.Version, error) {
	return &dco.Version{
		Major: 1,
		Minor: 0,
		Patch: 0,
	}, nil
}

func SystemRepositoryInit() *SystemRepositoryImpl {
	return &SystemRepositoryImpl{}
}
