package reportSaleDTO

import "time"

type (
	SendRequestDTO struct {
		Code        string               `json:"code"`         // * Агуулах код
		Date        time.Time            `json:"date"`         // * Тухайн өдрийн борлуулалтын огноо
		SectionCode string               `json:"section_code"` // * Машины код
		Items       []SendRequestItemDTO `json:"items"`        // * Борлуулалтын бүтээгдэхүүний жагсаалт
	}

	SendRequestItemDTO struct {
		ProductID   string  `json:"productId"`      // * Бүтээгдэхүүний код
		Qty         int     `json:"invoiceQty"`     // * Тоо ширхэг
		TotalPrice  float64 `json:"lineTotalPrice"` // * Барааны тоо хэмжээгээр үржсэн дүн байна . (НӨАТ - шингэсэн дүн байна)
		Description string  `json:"description"`    // * Optional : Шаардлагатай тайлбарыг барааны мөр бүрт явуулж болно
	}

	SendRequestActualItemDTO struct {
		ProductID       string  `json:"productId"`       // * Бүтээгдэхүүний код
		Qty             int     `json:"invoiceQty"`      // * Тоо ширхэг
		TotalPrice      float64 `json:"lineTotalPrice"`  // * Барааны тоо хэмжээгээр үржсэн дүн байна . (НӨАТ - шингэсэн дүн байна)
		Description     string  `json:"description"`     // * Optional : Шаардлагатай тайлбарыг барааны мөр бүрт явуулж болно
		UnitPrice       float64 `json:"unitPrice"`       // * Барааны үнэ байна
		UnitVat         float64 `json:"unitVat"`         // * Барааны нэгжийн НӨАТ
		LineTotalVat    float64 `json:"lineTotalVat"`    // * Барааны тоо хэмжээгээр үржсэн НӨАТ дүн байна
		UnitAmount      float64 `json:"unitAmount"`      // * Барааны НӨАТгүй дүн байна
		LineTotalAmount float64 `json:"lineTotalAmount"` // * Барааны тоо хэмжээгээр үржсэн НӨАТгүй нийт дүн байна
	}

	SendRequestActualDTO struct {
		Code        string                     `json:"storeCode"`                      // * Агуулах код
		Date        string                     `json:"invoiceDate"`                    // * Тухайн өдрийн борлуулалтын огноо
		SectionCode string                     `json:"sectionCode"`                    // * Машины код
		Vat         float64                    `json:"vat"`                            // * НӨАТ
		Total       float64                    `json:"total"`                          // * Нийт дүн
		Items       []SendRequestActualItemDTO `json:"SDM_SM_SALES_INVOICE_DETAIL_DV"` // * Борлуулалтын бүтээгдэхүүний жагсаалт
	}
)
