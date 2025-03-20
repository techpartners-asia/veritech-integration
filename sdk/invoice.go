package veritechSdk

import (
	"fmt"

	veritechModels "github.com/techpartners-asia/veritech-integration/sdk/models"
	veritechUtils "github.com/techpartners-asia/veritech-integration/sdk/utils"
)

type (
	SaleReportsService interface {
		SendSaleReport(input SaleReportInput) (*SaleReportResponse, error)
	}
)

type SaleReportProductInputItem struct {
	ProductID       string  `json:"productId"`
	Quantity        int     `json:"invoiceQty"`
	UnitPrice       float64 `json:"unitPrice"`
	LineTotalPrice  float64 `json:"lineTotalPrice"`
	UnitVat         float64 `json:"unitVat"`
	LineTotalVat    float64 `json:"lineTotalVat"`
	UnitAmount      float64 `json:"unitAmount"`
	LineTotalAmount float64 `json:"lineTotalAmount"`
	Description     string  `json:"description"`
}

type SaleReportInput struct {
	StoreCode   string                       `json:"storeCode"`
	SectionCode string                       `json:"sectionCode"`
	InvoiceDate string                       `json:"invoiceDate"`
	TotalAmount float64                      `json:"totalAmount"`
	TotalVat    float64                      `json:"totalVat"`
	Items       []SaleReportProductInputItem `json:"SDM_SM_SALES_INVOICE_DETAIL_DV"`
}

type SaleReportResponse struct {
	StoreCode                      string                               `json:"storecode"`
	SectionCode                    string                               `json:"sectioncode"`
	InvoiceDate                    string                               `json:"invoicedate"`
	Vat                            float64                              `json:"vat"`
	Total                          float64                              `json:"total"`
	SDM_SM_SALES_INVOICE_DETAIL_DV map[string]SaleReportProductResponse `json:"sdm_sm_sales_invoice_detail_dv"`
	CommandID                      string                               `json:"commandid"`
	ParentCommandID                string                               `json:"parentcommandid"`
	IsApproved                     int                                  `json:"isapproved"`
	IsRemoved                      int                                  `json:"isremoved"`
	SectionID                      string                               `json:"sectionid"`
	BookTypeID                     int                                  `json:"booktypeid"`
	CustomerName                   string                               `json:"customername"`
	StoreID                        string                               `json:"storeid"`
	InvoiceNumber                  string                               `json:"invoicenumber"`
	RefNumber                      string                               `json:"refnumber"`
	CashRegisterID                 string                               `json:"cashregisterid"`
	CreatedCashierID               string                               `json:"createdcashierid"`
	CustomerID                     string                               `json:"customerid"`
	SystemMetaGroupID              string                               `json:"systemmetagroupid"`
	RefMetaGroupID                 string                               `json:"refmetagroupid"`
	RowState                       string                               `json:"rowstate"`
	CompanyDepartmentID            string                               `json:"companydepartmentid"`
	ConnectionManagerID            string                               `json:"_connectionmanagerid"`
	DepartmentID                   string                               `json:"_departmentid"`
	UCompanyDepartmentID           string                               `json:"_companydepartmentid"`
	SubTotal                       float64                              `json:"subtotal"`
	ID                             string                               `json:"id"`
	CreatedDate                    string                               `json:"createddate"`
	CreatedUserID                  string                               `json:"createduserid"`
	CreatedCommandID               string                               `json:"createdcommandid"`
	CreatedUsername                string                               `json:"createdusername"`
}

type SaleReportProductResponse struct {
	ProductID            int     `json:"productid"`
	Quantity             int     `json:"invoiceqty"`
	UnitPrice            float64 `json:"unitprice"`
	LineTotalPrice       float64 `json:"linetotalprice"`
	UnitVat              float64 `json:"unitvat"`
	LineTotalVat         float64 `json:"linetotalvat"`
	UnitAmount           float64 `json:"unitamount"`
	LineTotalAmount      float64 `json:"linetotalamount"`
	Description          string  `json:"description"`
	PercentDiscount      float64 `json:"percentdiscount"`
	PercentVat           float64 `json:"percentvat"`
	IsRemoved            int     `json:"isremoved"`
	Rate                 float64 `json:"rate"`
	UnitCost             float64 `json:"unitcost"`
	TotalCost            float64 `json:"totalcost"`
	RefSalePrice         float64 `json:"refsaleprice"`
	RefWholesalePrice    float64 `json:"refwholesaleprice"`
	RefPurchasePrice     float64 `json:"refpurchaseprice"`
	RefContractPrice     float64 `json:"refcontractprice"`
	CompanyDepartmentID  string  `json:"companydepartmentid"`
	ConnectionManagerID  string  `json:"_connectionmanagerid"`
	DepartmentID         string  `json:"_departmentid"`
	UCompanyDepartmentID string  `json:"_companydepartmentid"`
	SalesInvoiceID       string  `json:"salesinvoiceid"`
	SectionID            string  `json:"sectionid"`
	StoreID              string  `json:"storeid"`
	ID                   string  `json:"id"`
	CreatedDate          string  `json:"createddate"`
	CreatedUserID        string  `json:"createduserid"`
	CreatedCommandID     string  `json:"createdcommandid"`
	CreatedUsername      string  `json:"createdusername"`
}

func (co *sdk) SendSaleReport(input SaleReportInput) (*SaleReportResponse, error) {

	co.Command = "smSalesInvoiceHeaderStoreVMSteso"

	body := map[string]interface{}{
		"parameters": map[string]interface{}{
			"storeCode":                      input.StoreCode,
			"sectionCode":                    input.SectionCode,
			"invoiceDate":                    input.InvoiceDate,
			"vat":                            input.TotalVat,
			"total":                          input.TotalAmount,
			"SDM_SM_SALES_INVOICE_DETAIL_DV": input.Items,
		},
	}

	res, err := veritechUtils.GetResponse[SaleReportResponse](co.Send, body)
	if err != nil {
		return nil, err
	}

	if res.Response.Status != veritechModels.VERITECH_SUCCESS {
		return nil, fmt.Errorf("Failed to send sale report: %v", res)
	}

	return &res.Response.Result, nil
}
