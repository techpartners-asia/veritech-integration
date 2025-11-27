# Veritech Integration SDK

This SDK provides a Go interface for integrating with the Veritech ERP system. It simplifies interaction with various Veritech services including Product Management, Sales Reporting, Warehouse Staff, and Warehouse Storage.

## Installation

```bash
go get github.com/techpartners-asia/veritech-integration
```

## Initialization

To use the SDK, you need to initialize it with your Veritech credentials and host URL.

```go
package main

import (
    "fmt"
    "github.com/techpartners-asia/veritech-integration/sdk-2/models"
    veritechSdk "github.com/techpartners-asia/veritech-integration/sdk-2"
)

func main() {
    config := veritechSdkModels.BaseServiceInput{
        Host:         "http://your-veritech-host/erp-services/RestWS/runJson",
        Username:     "your-username",
        Password:     "your-password",
        LanguageCode: "mn", // Optional
    }

    sdk := veritechSdk.New(config)

    // Now you can access sdk.Product, sdk.Staff, etc.
}
```

## Modules

The SDK is organized into several modules, each handling a specific domain of the Veritech system.

### 1. Product Module (`sdk.Product`)

Handles product-related operations.

#### Check if Product Exists

```go
exists := sdk.Product.IsExist("ITEM_CODE_001")
if exists {
    fmt.Println("Product exists")
}
```

#### Get Product Details

```go
product, err := sdk.Product.GetByCode("ITEM_CODE_001")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Product Name: %s\n", product.ItemName)
```

### 2. Warehouse Staff Module (`sdk.Staff`)

Manage warehouse staff verification and details.

#### Check if Staff Exists

```go
exists := sdk.Staff.IsExist("STAFF_CODE_001")
```

#### Get Staff Details

```go
staff, err := sdk.Staff.GetByCode("STAFF_CODE_001")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Staff Name: %s\n", staff.StoreKeeperName)
```

### 3. Warehouse Storage Module (`sdk.Storage`)

Retrieve warehouse storage information.

#### Get Storage Details

```go
// command: Specific Veritech command for storage retrieval
// code: Storage code
storage, err := sdk.Storage.GetByCode("STORAGE_COMMAND", "STORAGE_CODE")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Warehouse: %s\n", storage.WarehouseName)
```

### 4. Report Sale Module

Handles sending sales reports/invoices to Veritech. Note that this module usually requires initialization with the main client if used separately, but typical usage patterns might vary.

_Note: The Report Sale module structure involves sending detailed sales data including items, VAT, and totals._

#### Send Sales Report

```go
import (
    reportSaleDTO "github.com/techpartners-asia/veritech-integration/sdk-2/modules/report/dto/sale"
    usecaseReportSale "github.com/techpartners-asia/veritech-integration/sdk-2/modules/report/usecase/sale"
    "time"
)

// ... inside your function

// Initialize Report Sale Usecase if not directly on the main SDK struct (depending on your specific implementation)
// or if added to the main SDK struct:
// reportSdk := sdk.ReportSale

req := reportSaleDTO.SendRequestDTO{
    Code:        "STORE_001",
    Date:        time.Now(),
    SectionCode: "SECTION_001",
    Items: []reportSaleDTO.SendRequestItemDTO{
        {
            ProductID:  "PROD_001",
            Qty:        2,
            TotalPrice: 5000.0,
        },
    },
}

// Assuming you have access to the report usecase
// resp, err := reportSaleUsecase.Send(req)
```

## Directory Structure

```
.
├── sdk-2/
│   ├── base.go                 # Main SDK Entry point
│   ├── models/                 # Shared Models
│   └── modules/                # Functional Modules
│       ├── product/            # Product Management
│       ├── report/             # Reporting (Sales, Finance)
│       └── warehouse/          # Warehouse (Staff, Storage)
└── utils/                      # Utility functions
```

## License

[Add License Information Here]
