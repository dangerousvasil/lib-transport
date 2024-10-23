package tags

import (
	"lib-transport/pb"
)

func AccessibleRoles() map[string][]string {
	rls := map[string][]string{}
	rls[pb.Tags_Put_FullMethodName] = []string{}
	rls[pb.Tags_List_FullMethodName] = []string{}
	rls[pb.Tags_Search_FullMethodName] = []string{}
	rls[pb.Tags_Get_FullMethodName] = []string{}
	return rls
}
