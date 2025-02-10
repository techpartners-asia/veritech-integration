package structs

type InvoiceCreateRequest struct {
	Base
	Parameters InvoiceCreateParameters `json:"parameters"`
}

type InvoiceCreateParameters struct {
	StoreCode                 string          `json:"storeCode"`                      // код 888
	SectionCode               string          `json:"sectionCode"`                    // Машины код
	InvoiceDate               string          `json:"invoiceDate"`                    //
	Vat                       float64         `json:"vat"`                            // DTL-н lineTotalVat талбарын нийт дүн байна
	Total                     float64         `json:"total"`                          // DTL-н lineTotalPrice талбарын нийт дүн байна
	SdmSmSalesInvoiceDetailDv []InvoiceDetail `json:"SDM_SM_SALES_INVOICE_DETAIL_DV"` //
	SmSalesPaymentDV          []PaymentDV     `json:"sm_sales_payment_dv"`
}

type InvoiceDetail struct {
	ProductID       string  `json:"productId"`       // Барааны систем дугаар
	InvoiceQty      float64 `json:"invoiceQty"`      // Тоо ширхэг
	UnitPrice       float64 `json:"unitPrice"`       // Барааны үнэ байна
	LineTotalPrice  float64 `json:"lineTotalPrice"`  // Барааны тоо хэмжээгээр үржсэн дүн байна
	UnitVat         float64 `json:"unitVat"`         // Барааны нэгжийн НӨАТ
	LineTotalVat    float64 `json:"lineTotalVat"`    // Барааны тоо хэмжээгээр үржсэн НӨАТ байна
	UnitAmount      float64 `json:"unitAmount"`      // Барааны НӨАТгүй дүн байна
	LineTotalAmount float64 `json:"lineTotalAmount"` // Барааны тоо хэмжээгээр үржсэн НӨАТгүй нийт дүн байна
	Description     string  `json:"description"`     // Шаардлагатай тайлбарыг барааны мөр бүрт явуулж болно
}

type PaymentDV struct {
	Amount        float64 `json:"amount"`
	PaymentTypeID uint    `json:"paymenttypeid"`
}

type BaseResponse struct {
	Status string      `json:"status"` // error, success
	Text   string      `json:"text"`
	Code   string      `json:"code"`
	Result interface{} `json:"result"`
}

// INT_SM_EXTERNAL_HEADER_001- Гүйлгээ хүлээн авах
type SellResponce struct {
	Response BaseResponse `json:"response"`
}
