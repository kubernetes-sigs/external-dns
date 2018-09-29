package endpoints

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMappingResolver_TryResolve(t *testing.T) {
	AddEndpointMapping("cn-hangzhou", "Ecs", "unreachable.aliyuncs.com")
	resolveParam := &ResolveParam{
		RegionId: "cn-hangzhou",
		Product:  "ecs",
	}
	endpoint, err := Resolve(resolveParam)
	assert.Nil(t, err)
	assert.Equal(t, endpoint, "unreachable.aliyuncs.com")
	fmt.Println("finished")
}
