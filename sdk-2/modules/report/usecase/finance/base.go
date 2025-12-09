package usecaseReportFinance

import (
	"errors"
	"time"

	"github.com/go-resty/resty/v2"
	veritechSdkModels "github.com/techpartners-asia/veritech-integration/sdk-2/models"
	reportFinanceDTO "github.com/techpartners-asia/veritech-integration/sdk-2/modules/report/dto/finance"
	"github.com/techpartners-asia/veritech-integration/utils"
)

type ReportFinanceUsecase interface {
	List(code string, bookDate *time.Time) (*reportFinanceDTO.ListResponseDTO, error)
}

type reportFinanceUsecase struct {
	client *resty.Client
	input  veritechSdkModels.BaseServiceInput
}

func New(client *resty.Client, input veritechSdkModels.BaseServiceInput) ReportFinanceUsecase {
	return &reportFinanceUsecase{
		client: client,
		input:  input,
	}
}

func (r *reportFinanceUsecase) List(code string, bookDate *time.Time) (*reportFinanceDTO.ListResponseDTO, error) {

	var response veritechSdkModels.BaseServiceResponse[reportFinanceDTO.ListResponseDTO]

	res, err := r.client.R().
		SetBody(map[string]interface{}{
			"username":     r.input.Username,
			"password":     r.input.Password,
			"languageCode": r.input.LanguageCode,
			"command":      "sales_QTY_MTM_004",
			"parameters": map[string]interface{}{
				"filterStoreKeeperKeyCode": code,
				"filterbookdate":           utils.FormatDate(*bookDate),
			},
		}).SetResult(&response).Post("")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		return nil, errors.New(res.String())
	}
	if response.Response.Status != veritechSdkModels.ResponseStatusSuccess {
		return nil, errors.New(response.Response.Text)
	}

	return &response.Response.Result, nil
}

func (r *reportFinanceUsecase) GetInitialQty(code string) (*reportFinanceDTO.GetInitialQtyResponseDTO, error) {

	var response veritechSdkModels.BaseServiceResponse[reportFinanceDTO.GetInitialQtyResponseDTO]

	res, err := r.client.R().
		SetBody(map[string]interface{}{
			"username":     r.input.Username,
			"password":     r.input.Password,
			"languageCode": r.input.LanguageCode,
			"command":      "sales_BEGIN_QTY_MTM_004",
			"parameters": map[string]interface{}{
				"filterStoreKeeperKeyCode": code,
			},
		}).SetResult(&response).Post("")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		return nil, errors.New(res.String())
	}
	if response.Response.Status != veritechSdkModels.ResponseStatusSuccess {
		return nil, errors.New(response.Response.Text)
	}

	return &response.Response.Result, nil
}
