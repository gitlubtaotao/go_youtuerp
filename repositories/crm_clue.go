package repositories

type ICrmClue interface {
}
type CrmClue struct {
	BaseRepository
}

func NewCrmClue() ICrmClue {
	return &CrmClue{}
}
