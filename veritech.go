package veritechSdk

import (
	"encoding/json"
	"errors"
	"fmt"

	veritechModels "github.com/techpartners-asia/veritech-integration/structs"
)

type veritech struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Endpoint string `json:"endpoint"`
	Command  string `json:"command"`
}

type Veritech interface {
	GetProductItemID(code string) (*veritechModels.ProductInfoResponse, error)
	SendFillInformation(input veritechModels.StoreFillInput) (*veritechModels.StoreFillResponse, error)
	CheckStaffAndMachine(code string) error
	TakeSalesPersonBeginQty(code string) (*veritechModels.TakeHTBeginQtyResult, error)
	TakeStaffIncomeOutcome(staffCode string, date string) (*veritechModels.StaffWarehouseIncomeOutcomeResult, error)
	SendInvoice(param veritechModels.InvoiceCreateParameters) (veritechModels.SellResponce, string, error)
}

func New(username, password, endpoint, command string) Veritech {
	return &veritech{
		Username: username,
		Password: password,
		Endpoint: endpoint,
		Command:  command,
	}
}

// * NOTE * - Vertitech-ээс барааны мэдээлэл авах -> Veritech Item Id
func (q *veritech) GetProductItemID(code string) (*veritechModels.ProductInfoResponse, error) {
	input := veritechModels.ProductInfoInput{
		Base: veritechModels.Base{
			Username:     q.Username,
			Password:     q.Password,
			Command:      "MTM_ITEM_004",
			LanguageCode: "mn",
		},
		Paramters: veritechModels.ProductInfoParametersInput{
			ItemCode: code,
		},
	}

	res, err := q.httpRequest(input)
	if err != nil {
		return nil, err
	}

	var response veritechModels.ProductInfoResponse
	json.Unmarshal(res, &response)

	if response.Response.Status != "success" {
		return nil, errors.New(response.Response.Text)
	}

	return &response, nil
}

// * NOTE * - Машин дүүргэлт хийхэд дүүргэсэн барааны мэдээлэл Veritech - рүү илгээх
func (q *veritech) SendFillInformation(input veritechModels.StoreFillInput) (*veritechModels.StoreFillResponse, error) {

	input.Username = q.Username
	input.Password = q.Password
	input.LanguageCode = "mn"
	input.Command = "imCheckMovementItem_001"

	res, err := q.httpRequest(input)
	if err != nil {
		fmt.Println(" Veritech Send Fill Information Request error - %v", err.Error())
		return nil, err
	}

	var response veritechModels.StoreFillResponse
	if err := json.Unmarshal(res, &response); err != nil {
		fmt.Println(" Veritech Send Fill Information Response marshalling error - %v", err.Error())
		return nil, err
	}

	return &response, nil
}

func (q *veritech) CheckStaffAndMachine(code string) error {
	tmp := make(map[string]interface{})

	tmp["storeKeeperKeyCode"] = code

	request := veritechModels.CheckRequest{
		Request: &veritechModels.RequestBody{
			Command:    "mtm_sales_skk_code_004",
			Username:   q.Username,
			Password:   q.Password,
			Parameters: tmp,
		},
	}

	res, err := q.httpRequest(request)
	if err != nil {
		return err
	}

	var resCheck veritechModels.StaffMachineCheckResponse
	json.Unmarshal(res, &resCheck)

	if resCheck.Response.Status == "success" {
		if str, ok := resCheck.Response.Result.(string); ok && str == "" {
			// fmt.Println("Veritech дээр байхгүй байна.")
			return errors.New("veritech дээр байхгүй байна")
		} else {
			return nil
		}
	}

	return errors.New("veritech алдаа")
}

// * NOTE * - ХТ эхний үлдэгдэл авах
func (q *veritech) TakeSalesPersonBeginQty(code string) (*veritechModels.TakeHTBeginQtyResult, error) {

	tmp := make(map[string]interface{})
	tmp["filterStoreKeeperKeyCode"] = code

	request := veritechModels.CheckRequest{
		Request: &veritechModels.RequestBody{
			Command:    "sales_BEGIN_QTY_MTM_004",
			Username:   q.Username,
			Password:   q.Password,
			Parameters: tmp,
		},
	}

	res, err := q.httpRequest(request)
	if err != nil {
		return nil, err
	}

	var resCheck veritechModels.TakeHTBeginQtyResponse
	json.Unmarshal(res, &resCheck)

	if resCheck.Response.Status == "success" {
		if resCheck.Response.Result != nil {
			return resCheck.Response.Result, nil
		}

		return nil, errors.New("veritech дээр байхгүй байна")
	}

	return nil, errors.New("veritech алдаа")
}

// * NOTE * - Ажилтаны орлого, зарлагыг авах odortoi
func (q *veritech) TakeStaffIncomeOutcome(staffCode string, date string) (*veritechModels.StaffWarehouseIncomeOutcomeResult, error) {
	tmp := make(map[string]interface{})
	tmp["filterStoreKeeperKeyCode"] = staffCode
	tmp["filterbookdate"] = date

	request := veritechModels.CheckRequest{
		Request: &veritechModels.RequestBody{
			Command:    "sales_QTY_MTM_004",
			Username:   q.Username,
			Password:   q.Password,
			Parameters: tmp,
		},
	}

	res, err := q.httpRequest(request)
	if err != nil {
		return nil, err
	}

	var resCheck veritechModels.StaffWarehouseIncomeOutcome
	json.Unmarshal(res, &resCheck)

	if resCheck.Response.Status == "success" {
		if resCheck.Response.Result != nil {
			return resCheck.Response.Result, nil
		}

		return nil, errors.New("veritech дээр байхгүй байна")
	}

	return nil, errors.New("veritech алдаа")
}

// * NOTE * - Veritech руу борлуулалтын мэдээлэл илгээх
func (q *veritech) SendInvoice(param veritechModels.InvoiceCreateParameters) (veritechModels.SellResponce, string, error) {
	httpRequest := veritechModels.InvoiceCreateRequest{
		Base: veritechModels.Base{
			Username: q.Username,
			Password: q.Password,
			Command:  "smSalesInvoiceHeaderStoreVMSteso",
		},
		Parameters: param,
	}

	res, err := q.httpRequest(httpRequest)
	if err != nil {
		return veritechModels.SellResponce{}, "", err
	}

	var response veritechModels.SellResponce
	json.Unmarshal(res, &response)

	return response, string(res), nil
}
