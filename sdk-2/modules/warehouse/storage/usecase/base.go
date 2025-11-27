package usecaseWarehouseStorage

import (
	"errors"

	"github.com/go-resty/resty/v2"
	veritechSdkModels "github.com/techpartners-asia/veritech-integration/sdk-2/models"
	warehouseStorageDTO "github.com/techpartners-asia/veritech-integration/sdk-2/modules/warehouse/storage/dto"
)

type WarehouseStorageUsecase interface {
	GetByCode(command string, code string) (*warehouseStorageDTO.WarehouseStorageResponseDTO, error)
}

type warehouseStorageUsecase struct {
	client *resty.Client
	input  veritechSdkModels.BaseServiceInput
}

func New(client *resty.Client, input veritechSdkModels.BaseServiceInput) WarehouseStorageUsecase {
	return &warehouseStorageUsecase{
		client: client,
		input:  input,
	}
}

func (w *warehouseStorageUsecase) GetByCode(command string, code string) (*warehouseStorageDTO.WarehouseStorageResponseDTO, error) {

	var response veritechSdkModels.BaseServiceResponse[warehouseStorageDTO.WarehouseStorageResponseDTO]

	res, err := w.client.R().
		SetBody(map[string]interface{}{
			"username": w.input.Username,
			"password": w.input.Password,
			"command":  command,
			"parameters": map[string]interface{}{
				"storeKeeperKeyCode": code,
			},
		}).SetResult(&response).Post("")
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		return nil, errors.New(res.String())
	}

	if response.Response.Status != veritechSdkModels.ResponseStatusSuccess {
		return nil, errors.New(res.String())
	}

	return &response.Response.Result, nil
}
