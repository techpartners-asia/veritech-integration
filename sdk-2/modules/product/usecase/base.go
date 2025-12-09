package usecaseProduct

import (
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	veritechSdkModels "github.com/techpartners-asia/veritech-integration/sdk-2/models"
	usecaseProductDTO "github.com/techpartners-asia/veritech-integration/sdk-2/modules/product/dto"
)

type ProductUsecase interface {
	IsExist(code string) (bool, error)
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

func (p *productUsecase) IsExist(code string) (bool, error) {

	product, err := p.GetByCode(code)
	if err != nil {
		return false, err
	}

	return len(product.ItemID) > 0, nil
}

func (p *productUsecase) GetByCode(code string) (*usecaseProductDTO.ProductResponseDTO, error) {

	var response veritechSdkModels.BaseServiceResponse[usecaseProductDTO.ProductResponseDTO]

	res, err := p.client.R().
		SetBody(map[string]interface{}{
			"username": p.input.Username,
			"password": p.input.Password,
			"command":  "MTM_ITEM_004",
			"parameters": map[string]interface{}{
				"systemmetagroupcode": "PL_MDVIEW_005",
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

	fmt.Println(res.Request.Body)

	if response.Response.Status != veritechSdkModels.ResponseStatusSuccess {
		return nil, errors.New(response.Response.Text)
	}

	if len(response.Response.Result.ItemID) == 0 {
		return nil, errors.New("product not found")
	}

	return &response.Response.Result, nil
}
