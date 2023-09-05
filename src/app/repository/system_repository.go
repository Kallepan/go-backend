package repository

type SystemRepository interface {
	GetVersion() (string, error)
}

type SystemRepositoryImpl struct {
}

func (s SystemRepositoryImpl) GetVersion() (string, error) {
	return "1.0.0", nil
}

func SystemRepositoryInit() *SystemRepositoryImpl {
	return &SystemRepositoryImpl{}
}
