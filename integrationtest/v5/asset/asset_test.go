//go:build integrationtestv5asset

package integrationtestv5asset

import (
	"testing"

	"github.com/hirokisan/bybit/v2"
	"github.com/hirokisan/bybit/v2/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestGetInternalTransferRecords(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	limit := 1
	res, err := client.V5().Asset().GetInternalTransferRecords(bybit.V5GetInternalTransferRecordsParam{
		Limit: &limit,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-asset-get-internal-transfer-records.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}
