package veritechSdk

type BaseService interface {
	StaffService() StaffManagerService
	ProductService() ProductService
}

type sdk struct {
	BaseServiceModel
}

func New(input BaseServiceModel) BaseService {
	return &sdk{
		input,
	}
}

func (co *sdk) StaffService() StaffManagerService {
	return co
}

func (co *sdk) ProductService() ProductService {
	// co.Command = "MTM_ITEM_004"
	return co
}
