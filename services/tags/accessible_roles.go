package tags

import (
	"lib-transport/transport"
)

func AccessibleRoles() map[string][]string {
	rls := map[string][]string{}
	rls[transport.Tags_Put_FullMethodName] = []string{}
	rls[transport.Tags_List_FullMethodName] = []string{}
	rls[transport.Tags_Search_FullMethodName] = []string{}
	rls[transport.Tags_Get_FullMethodName] = []string{}
	return rls
}
