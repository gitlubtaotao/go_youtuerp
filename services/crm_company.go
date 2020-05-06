package services
type ICrmCompanyService interface {
}

type CrmCompanyService struct {
	//repo repositories.ICrmCompanyRepository
	BaseService
}

func NewCrmCompanyService() ICrmCompanyService {
	return &CrmCompanyService{}
}
