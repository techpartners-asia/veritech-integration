package reportFinance

import (
	"fmt"
	"testing"
	"time"

	"github.com/gookit/goutil/testutil/assert"
	veritechSdk "github.com/techpartners-asia/veritech-integration/sdk-2"
)

func TestFinanceList(t *testing.T) {
	sdk := veritechSdk.NewTest()

	now := time.Now()

	finance, err := sdk.ReportFinance.List("MTM-INV-001", &now)
	fmt.Println(err)
	assert.Error(t, err)
	assert.Nil(t, finance)
}
