package repositories

type IContactRepository interface {
}

type ContactRepository struct {
	BaseRepository
}

func NewContactRepository() IContactRepository {
	return ContactRepository{}
}
