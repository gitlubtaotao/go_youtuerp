package repositories

type IOrderMaster interface {

}
type OrderMasterRepository struct {
	BaseRepository
}

func NewOrderMasterRepository() IOrderMaster {
	return 	OrderMasterRepository{}
}
