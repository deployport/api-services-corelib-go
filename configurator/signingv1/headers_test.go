package signingv1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeaderMatcher(t *testing.T) {
	t.Run("MatchHeader", func(t *testing.T) {
		t.Run("returns true if any of the matchers match", func(t *testing.T) {
			hm := headerMatchers{
				headerMatcherFunc(func(name string) bool {
					return name == "foo"
				}),
				headerMatcherFunc(func(name string) bool {
					return name == "bar"
				}),
			}
			require.True(t, hm.MatchHeader("foo"))
			require.True(t, hm.MatchHeader("bar"))
			require.False(t, hm.MatchHeader("baz"))
		})
	})
	require.Equal(t, "X-Dpp-Date", vendorHeaderCanonicalNameKey("dpp", "DAtE"))
}
