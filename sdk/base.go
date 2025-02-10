package veritechSdk

type BaseService interface {
	StaffService() StaffManagerService
	ProductService() ProductService
}

type Sdk struct {
	BaseServiceModel
}

func New(input BaseServiceModel) BaseService {
	return &Sdk{
		input,
	}
}

func (co *Sdk) StaffService() StaffManagerService {
	return co
}

func (co *Sdk) ProductService() ProductService {
	// co.Command = "MTM_ITEM_004"
	return co
}
