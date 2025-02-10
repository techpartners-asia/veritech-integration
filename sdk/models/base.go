package veritechModels

type (
	BaseServiceModel struct {
		Username         string `json:"username"`
		Password         string `json:"password"`
		Command          string `json:"command"`
		LanguageCode     string `json:"languageCode"`
		GroupCode        string `json:"systemmetagroupcode"`
		IgnorePermission string `json:"ignorepermission"`
	}

	// give me base response
	BaseServiceResponse[T any] struct {
		Response struct {
			Status string `json:"status"`
			Result T      `json:"result"`
		}
	}

	BookDtlModel struct {
		RefItemID   string `json:"refitemid"`   // Барааны систем дугаар
		Transferqty string `json:"transferqty"` // Тоо ширхэг
	}

	CodeModel struct {
		FromStoreKeeperKeyCode string "json:'fromstorekeeperkeycode'" // Хт код
		ToStoreKeeperKeyCode   string "json:'tostorekeeperkeycode'"   // Машины код
	}

	CheckKeyModel struct {
		ItemID               string `json:"itemid"`
		ImportID             string `json:"importid"`
		CompanyDepartmentID  string `json:"companydepartmentid"`
		ConnectionManagerID  string `json:"_connectionmanagerid"`
		UDepartmentID        string `json:"_departmentid"`
		UCompanyDepartmentID string `json:"_companydepartmentid"`
		StoreKeeperKeyID     string `json:"storekeeperkeyid"`
		LocationID           string `json:"locationid"`
		WarehouseID          string `json:"warehouseid"`
		ModifiedDate         string `json:"modifieddate"`
		ModifiedUserID       string `json:"modifieduserid"`
		ModifiedCommandID    string `json:"modifiedcommandid"`
		ModifiedUsername     string `json:"modifiedusername"`
		ID                   string `json:"id"`
	}
)
