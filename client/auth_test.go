package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCient_GetAuthMeta(t *testing.T) {
	meta := getAuthMeta("michael", "geheim")
	assert.Equal(t, "basic bWljaGFlbDpnZWhlaW0=", meta["authorization"])

	meta = getAuthMeta("Yårkl→nd", "`564ΣΛ");
	assert.Equal(t, "basic WcOlcmts4oaSbmQ6YDU2NM6jzps=", meta["authorization"]);
}
