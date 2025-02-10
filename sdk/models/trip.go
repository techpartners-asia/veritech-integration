package veritechModels

type (
	MakeTripBodyModel struct {
		CodeModel
		BookDate    string         "json:'bookdate'"    // Огноо
		Description string         "json:'description'" // Тайлбар
		BookDtls    []BookDtlModel "json:'bookdtl'"     // Бараа бүтээгдэхүүнүүд
	}

	MakeTripResponse struct {
		CodeModel
		BookDate    string                         "json:'bookdate'"
		Description string                         "json:'description'"
		BookDtl     map[string]MakeTripResponseDtl "json:'bookdtl'"
	}

	MakeTripResponseDtl struct {
		BookDtlModel
		CheckKey             CheckKeyModel "json:'checkkey'"
		EndQty               int           "json:'endqty'"
		BlockQty             string        "json:'blockqty'"
		BoxQty               string        "json:'boxqty'"
		UnitQty              string        "json:'unitqty'"
		UnitPrice            string        "json:'unitprice'"
		LineTotalPrice       string        "json:'linetotalprice'"
		UnitCost             string        "json:'unitcost'"
		LineTotalCost        string        "json:'linetotalcost'"
		DeficitQty           string        "json:'deficitqty'"
		CompanyDepartmentID  string        "json:'companydepartmentid'"
		ConnectionManagerID  string        "json:'_connectionmanagerid'"
		UDepartmentID        string        "json:'_departmentid'"
		UCompanyDepartmentID string        "json:'_companydepartmentid'"
		ItemMovementID       string        "json:'itemmovementid'"
		RefLocationID        string        "json:'reflocationid'"
		LocationID           string        "json:'locationid'"
		CheckKeyID           string        "json:'checkkeyid'"
		ToStoreKeeperKeyID   string        "json:'tostorekeeperkeyid'"
		FromStoreKeeperKeyID string        "json:'fromstorekeeperkeyid'"
		FromAccountID        string        "json:'fromaccountid'"
		ToAccountID          string        "json:'toaccountid'"
		ID                   string        "json:'id'"
		CreatedDate          string        "json:'createddate'"
		CreatedUserID        string        "json:'createduserid'"
		CreatedCommandID     string        "json:'createdcommandid'"
		CreatedUsername      string        "json:'createdusername'"
	}
)
