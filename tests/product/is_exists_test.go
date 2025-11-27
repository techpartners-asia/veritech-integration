package product

import (
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	veritechSdk "github.com/techpartners-asia/veritech-integration/sdk-2"
)

func TestIsExist(t *testing.T) {
	sdk := veritechSdk.NewTest()
	isExist := sdk.Product.IsExist("21080018")
	assert.True(t, isExist)

	isExist = sdk.Product.IsExist("1234567890")
	assert.False(t, isExist)
}
