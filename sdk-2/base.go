package veritechSdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	veritechSdkModels "github.com/techpartners-asia/veritech-integration/sdk-2/models"
	usecaseProduct "github.com/techpartners-asia/veritech-integration/sdk-2/modules/product/usecase"
	usecaseReportFinance "github.com/techpartners-asia/veritech-integration/sdk-2/modules/report/usecase/finance"
	usecaseWarehouseStaff "github.com/techpartners-asia/veritech-integration/sdk-2/modules/warehouse/staff/usecase"
	usecaseWarehouseStorage "github.com/techpartners-asia/veritech-integration/sdk-2/modules/warehouse/storage/usecase"
)

type Sdk struct {
	input         veritechSdkModels.BaseServiceInput
	Staff         usecaseWarehouseStaff.WarehouseStaffUsecase
	Storage       usecaseWarehouseStorage.WarehouseStorageUsecase
	Product       usecaseProduct.ProductUsecase
	ReportFinance usecaseReportFinance.ReportFinanceUsecase
}

func New(input veritechSdkModels.BaseServiceInput) *Sdk {

	client := resty.New().SetBaseURL(fmt.Sprintf("%s/erp-services/RestWS/runJson", input.Host))

	return &Sdk{
		input:         input,
		Staff:         usecaseWarehouseStaff.New(client, input),
		Storage:       usecaseWarehouseStorage.New(client, input),
		Product:       usecaseProduct.New(client, input),
		ReportFinance: usecaseReportFinance.New(client, input),
	}
}

func NewTest() *Sdk {
	return New(veritechSdkModels.BaseServiceInput{
		Host:     "http://veritech.zahii.mn",
		Username: "anketIntegration",
		Password: "Teso@123",
	})
}
