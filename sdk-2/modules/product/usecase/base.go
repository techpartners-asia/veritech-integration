package usecaseProduct

import (
	"errors"

	"github.com/go-resty/resty/v2"
	veritechSdkModels "github.com/techpartners-asia/veritech-integration/sdk-2/models"
	usecaseProductDTO "github.com/techpartners-asia/veritech-integration/sdk-2/modules/product/dto"
)

type ProductUsecase interface {
	IsExist(code string) bool
	GetByCode(code string) (*usecaseProductDTO.ProductResponseDTO, error)
}

type productUsecase struct {
	client *resty.Client
	input  veritechSdkModels.BaseServiceInput
}

func New(client *resty.Client, input veritechSdkModels.BaseServiceInput) ProductUsecase {
	return &productUsecase{
		client: client,
		input:  input,
	}
}

func (p *productUsecase) IsExist(code string) bool {

	var response veritechSdkModels.BaseServiceResponse[usecaseProductDTO.ProductResponseDTO]

	res, err := p.client.R().
		SetBody(map[string]interface{}{
			"username": p.input.Username,
			"password": p.input.Password,
			"command":  "PL_MDVIEW_005",
			"parameters": map[string]interface{}{
				"systemmetagroupcode": "MTM_ITEM_004",
				"ignorepermission":    "1",
				"itemCode":            code,
			},
		}).SetResult(&response).Post("")
	if err != nil {
		return false
	}

	if res.StatusCode() != 200 {
		return false
	}

	if response.Response.Status != veritechSdkModels.ResponseStatusSuccess {
		return false
	}

	return len(response.Response.Result.ItemID) > 0
}

func (p *productUsecase) GetByCode(code string) (*usecaseProductDTO.ProductResponseDTO, error) {

	var response veritechSdkModels.BaseServiceResponse[usecaseProductDTO.ProductResponseDTO]

	res, err := p.client.R().
		SetBody(map[string]interface{}{
			"username": p.input.Username,
			"password": p.input.Password,
			"command":  "PL_MDVIEW_005",
			"parameters": map[string]interface{}{
				"systemmetagroupcode": "MTM_ITEM_004",
				"ignorepermission":    "1",
				"itemCode":            code,
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
