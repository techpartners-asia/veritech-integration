package usecaseReportSale

import (
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	veritechSdkModels "github.com/techpartners-asia/veritech-integration/sdk-2/models"
	usecaseProduct "github.com/techpartners-asia/veritech-integration/sdk-2/modules/product/usecase"
	reportSaleDTO "github.com/techpartners-asia/veritech-integration/sdk-2/modules/report/dto/sale"
	"github.com/techpartners-asia/veritech-integration/utils"
)

type ReportSaleUsecase interface {
	Send(input reportSaleDTO.SendRequestDTO) (*reportSaleDTO.SendResponseDTO, error)
}

type reportSaleUsecase struct {
	client         *resty.Client
	productUsecase usecaseProduct.ProductUsecase
}

func New(client *resty.Client, input veritechSdkModels.BaseServiceInput) ReportSaleUsecase {
	return &reportSaleUsecase{
		client: client,
	}
}

func (r *reportSaleUsecase) Send(input reportSaleDTO.SendRequestDTO) (*reportSaleDTO.SendResponseDTO, error) {

	for _, item := range input.Items {
		isExist, err := r.productUsecase.IsExist(item.ProductID)
		if err != nil {
			return nil, err
		}

		if !isExist {
			return nil, fmt.Errorf("%v product is not exist", item.ProductID)
		}
	}

	var response veritechSdkModels.BaseServiceResponse[reportSaleDTO.SendResponseDTO]

	totalVat, totalAmount, actualItems := r.Format(input.Items)

	res, err := r.client.R().
		SetBody(map[string]interface{}{
			"command": "smSalesInvoiceHeaderStoreVMSteso",
			"parameters": map[string]interface{}{
				"storeCode":                      input.Code,
				"invoiceDate":                    utils.FormatDate(input.Date),
				"sectionCode":                    input.SectionCode,
				"vat":                            totalVat,
				"total":                          totalAmount,
				"SDM_SM_SALES_INVOICE_DETAIL_DV": actualItems,
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

func (r *reportSaleUsecase) Format(items []reportSaleDTO.SendRequestItemDTO) (float64, float64, []reportSaleDTO.SendRequestActualItemDTO) {

	var groupedItems = make(map[string]reportSaleDTO.SendRequestActualItemDTO)
	totalVat := 0.0
	totalAmount := 0.0

	for _, item := range items {

		unitPrice := utils.NumberPrecision(item.TotalPrice / float64(item.Qty))
		unitVat := utils.CalculateVat(unitPrice)
		unitWithNoVat := utils.CalculateWithNoVat(unitPrice)

		if _, ok := groupedItems[item.ProductID]; ok {
			grouped := groupedItems[item.ProductID]
			grouped.UnitPrice += unitPrice
			grouped.UnitVat += unitVat
			grouped.LineTotalVat += utils.NumberPrecision(unitVat * float64(item.Qty))
			grouped.UnitAmount += unitWithNoVat
			grouped.LineTotalAmount += utils.NumberPrecision(unitWithNoVat * float64(item.Qty))
			grouped.Description += item.Description
			grouped.Qty += item.Qty
			grouped.TotalPrice += item.TotalPrice

			groupedItems[item.ProductID] = grouped
		} else {
			groupedItems[item.ProductID] = reportSaleDTO.SendRequestActualItemDTO{
				ProductID:       item.ProductID,
				Qty:             item.Qty,
				TotalPrice:      item.TotalPrice,
				Description:     item.Description,
				UnitPrice:       unitPrice,
				UnitVat:         unitVat,
				LineTotalVat:    utils.NumberPrecision(unitVat * float64(item.Qty)),
				UnitAmount:      unitWithNoVat,
				LineTotalAmount: utils.NumberPrecision(unitWithNoVat * float64(item.Qty)),
			}
		}
	}

	var actualItems []reportSaleDTO.SendRequestActualItemDTO
	for _, item := range groupedItems {
		actualItems = append(actualItems, item)
		totalVat += item.UnitVat
		totalAmount += item.TotalPrice
	}

	return totalVat, totalAmount, actualItems
}
