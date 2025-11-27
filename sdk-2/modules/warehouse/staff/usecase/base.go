package usecaseWarehouseStaff

import (
	"errors"

	"github.com/go-resty/resty/v2"
	veritechSdkModels "github.com/techpartners-asia/veritech-integration/sdk-2/models"
	warehouseStaffDTO "github.com/techpartners-asia/veritech-integration/sdk-2/modules/warehouse/staff/dto"
)

type WarehouseStaffUsecase interface {
	IsExist(code string) bool
	GetByCode(code string) (*warehouseStaffDTO.StaffResponseDTO, error)
}

type warehouseStaffUsecase struct {
	client *resty.Client
	input  veritechSdkModels.BaseServiceInput
}

func New(client *resty.Client, input veritechSdkModels.BaseServiceInput) WarehouseStaffUsecase {
	return &warehouseStaffUsecase{
		client: client,
		input:  input,
	}
}

func (w *warehouseStaffUsecase) IsExist(code string) bool {

	var response veritechSdkModels.BaseServiceResponse[warehouseStaffDTO.IsExistResponseDTO]

	res, err := w.client.R().
		SetBody(map[string]interface{}{
			"username": w.input.Username,
			"password": w.input.Password,
			"command":  "mtm_sales_skk_code_004",
			"parameters": map[string]interface{}{
				"storeKeeperKeyCode": code,
			},
		}).SetResult(&response).Post("")
	if err != nil {
		return false
	}

	if res.StatusCode() != 200 {
		return false
	}

	return true
}

func (w *warehouseStaffUsecase) GetByCode(code string) (*warehouseStaffDTO.StaffResponseDTO, error) {

	var response veritechSdkModels.BaseServiceResponse[warehouseStaffDTO.StaffResponseDTO]

	res, err := w.client.R().
		SetBody(map[string]interface{}{
			"username": w.input.Username,
			"password": w.input.Password,
			"command":  "mtm_sales_skk_code_004",
			"parameters": map[string]interface{}{
				"storeKeeperKeyCode": code,
			},
		}).SetResult(&response).Post("")
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		return nil, errors.New("status code is not 200")
	}

	if response.Response.Status != veritechSdkModels.ResponseStatusSuccess {
		return nil, errors.New("status is not success")
	}

	return &response.Response.Result, nil
}
