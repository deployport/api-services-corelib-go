package signingv1

import (
	"log/slog"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	sdk "go.deployport.com/specular-runtime/client"
)

func TestSignature(t *testing.T) {
	lvl := new(slog.LevelVar)
	lvl.Set(slog.LevelDebug)

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: lvl,
	}))
	// defaultCredentials is the default credentials used in testing
	var defaultCredentials = Credentials{
		KeyID:  "testkeyid1",
		Secret: "secretkey1",
	}
	body := []byte{1, 2, 3}
	req, err := http.NewRequest("GET", "http://example.com/mwssvcapi/crm/contacts/list", nil)
	require.NoError(t, err)
	date, err := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	require.NoError(t, err)

	pk := sdk.NewPackage(sdk.ModulePathFromTrustedValues("test", "crm"))
	contactsRes, err := pk.NewResource("contacts")
	require.NoError(t, err)

	listOp, err := contactsRes.NewOperation("list")
	require.NoError(t, err)

	signature, err := SignHeaders(logger, "mws", "hellosvc", defaultCredentials, "us-west-x", req, body, listOp, date)
	require.NoError(t, err)
	require.NotNil(t, signature)
	headers := req.Header
	require.Equal(t, "hellosvc", headers.Get("X-Mws-Service"))
	require.Equal(t, "list", headers.Get("X-Mws-Operation"))
	require.Equal(t, "contacts", headers.Get("X-Mws-Resource"))
	require.Equal(t, "us-west-x", headers.Get("X-Mws-Region"))
	require.Equal(t, "20210101T000000Z", headers.Get("X-Mws-Date"))
	require.Equal(t, "f68207b02aa178bb44061750d83bfc4208bedaa7a03ffe99ba6d507783ac4b56", signature.Digest)
}
