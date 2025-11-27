package reportSaleDTO

import "time"

type (
	SendResponseDTO struct {
		Code                string                         `json:"storecode"`                      // * Агуулах код
		Date                string                         `json:"invoicedate"`                    // * Тухайн өдрийн борлуулалтын огноо
		SectionCode         string                         `json:"sectioncode"`                    // * Машины код
		Vat                 float64                        `json:"vat"`                            // * НӨАТ
		Total               float64                        `json:"total"`                          // * Нийт дүн
		Items               map[string]SendResponseItemDTO `json:"sdm_sm_sales_invoice_detail_dv"` // * Борлуулалтын бүтээгдэхүүний жагсаалт
		CommandID           string                         `json:"commandid"`                      // * Команд ID
		ParentCommandID     string                         `json:"parentcommandid"`                // * Parent command ID
		IsApproved          int                            `json:"isapproved"`                     // * Is approved (0: false, 1: true)
		IsRemoved           int                            `json:"isremoved"`                      // * Is removed (0: false, 1: true)
		SectionID           string                         `json:"sectionid"`                      // * Section ID
		BookTypeID          string                         `json:"booktypeid"`                     // * Book type ID
		CustomerName        string                         `json:"customername"`                   // * Customer name
		StoreID             string                         `json:"storeid"`                        // * Store ID
		InvoiceNumber       string                         `json:"invoicenumber"`                  // * Invoice number
		RefNumber           string                         `json:"refnumber"`                      // * Ref number
		CashRegisterID      string                         `json:"cashregisterid"`                 // * Cash register ID
		CreatedCashierID    string                         `json:"createdcashierid"`               // * Created cashier ID
		CustomerID          string                         `json:"customerid"`                     // * Customer ID
		SystemMetaGroupID   string                         `json:"systemmetagroupid"`              // * System meta group ID
		RefMetaGroupID      string                         `json:"refmetagroupid"`                 // * Ref meta group ID
		RowState            string                         `json:"rowstate"`                       // * Row state
		CompanyDepartmentID string                         `json:"companydepartmentid"`            // * Company department ID
		ConnectionManagerID string                         `json:"_connectionmanagerid"`           // * Connection manager ID
		DepartmentID        string                         `json:"_departmentid"`                  // * Department ID
		SubTotal            float64                        `json:"subtotal"`                       // * Sub total
		ID                  string                         `json:"id"`                             // * ID
		CreatedDate         *time.Time                     `json:"createddate"`                    // * Created date
		CreatedUserID       string                         `json:"createduserid"`                  // * Created user ID
		CreatedCommandID    string                         `json:"createdcommandid"`               // * Created command ID
		CreatedUsername     string                         `json:"createdusername"`                // * Created username
	}

	SendResponseItemDTO struct {
		ProductID           string  `json:"productid"`            // * Бүтээгдэхүүний код
		Qty                 int     `json:"invoiceqty"`           // * Тоо ширхэг
		UnitPrice           float64 `json:"unitprice"`            // * Барааны үнэ байна
		LineTotalPrice      float64 `json:"linetotalprice"`       // * Барааны тоо хэмжээгээр үржсэн дүн байна
		UnitVat             float64 `json:"unitvat"`              // * Барааны нэгжийн НӨАТ
		LineTotalVat        float64 `json:"linetotalvat"`         // * Барааны тоо хэмжээгээр үржсэн НӨАТ дүн байна
		UnitAmount          float64 `json:"unitamount"`           // * Барааны НӨАТгүй дүн байна
		LineTotalAmount     float64 `json:"linetotalamount"`      // * Барааны тоо хэмжээгээр үржсэн НӨАТгүй нийт дүн байна
		Description         string  `json:"description"`          // * Optional : Шаардлагатай тайлбарыг барааны мөр бүрт явуулж болно
		PercentDiscount     float64 `json:"percentdiscount"`      // * Discount percentage
		PercentVat          float64 `json:"percentvat"`           // * VAT percentage
		IsRemoved           int     `json:"isremoved"`            // * Whether the item is removed (0: false, 1: true)
		Rate                float64 `json:"rate"`                 // * Rate
		UnitCost            float64 `json:"unitcost"`             // * Unit cost
		TotalCost           float64 `json:"totalcost"`            // * Total cost
		RefSalePrice        float64 `json:"refsaleprice"`         // * Reference sale price
		RefWholeSalePrice   float64 `json:"refwholesaleprice"`    // * Reference wholesale price
		RefPurchasePrice    float64 `json:"refpurchaseprice"`     // * Reference purchase price
		RefContractPrice    float64 `json:"refcontractprice"`     // * Reference contract price
		CompanyDepartmentID string  `json:"companydepartmentid"`  // * Company department ID
		ConnectionManagerID string  `json:"_connectionmanagerid"` // * Connection manager ID
		DepartmentID        string  `json:"_departmentid"`        // * Department ID
		SalesInvoiceID      string  `json:"salesinvoiceid"`       // * Sales invoice ID
		SectionID           string  `json:"sectionid"`            // * Section ID
		StoreID             string  `json:"storeid"`              // * Store ID
		ID                  string  `json:"id"`                   // * ID
		CreatedDate         string  `json:"createddate"`          // * Created date
		CreatedUserID       string  `json:"createduserid"`        // * Created user ID
		CreatedCommandID    string  `json:"createdcommandid"`     // * Created command ID
		CreatedUsername     string  `json:"createdusername"`      // * Created username
	}
)
