package httpclient

import (
	"testing"

	"github.com/benthosdev/benthos/v4/internal/httpclient/oldconfig"
	"github.com/benthosdev/benthos/v4/internal/manager/mock"
	"github.com/benthosdev/benthos/v4/internal/message"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSingleMessageHeaders(t *testing.T) {
	oldConf := oldconfig.NewOldConfig()
	oldConf.Headers["Content-Type"] = "foo"
	oldConf.Metadata.IncludePrefixes = []string{"more_"}

	reqCreator, err := RequestCreatorFromOldConfig(oldConf, mock.NewManager())
	require.NoError(t, err)

	part := message.NewPart([]byte("hello world"))
	part.MetaSet("more_bar", "barvalue")
	part.MetaSet("ignore_baz", "bazvalue")

	b := message.Batch{part}

	req, err := reqCreator.Create(b)
	require.NoError(t, err)

	assert.Equal(t, []string{"foo"}, req.Header.Values("Content-Type"))
	assert.Equal(t, []string{"barvalue"}, req.Header.Values("more_bar"))
	assert.Equal(t, []string(nil), req.Header.Values("ignore_baz"))
}
