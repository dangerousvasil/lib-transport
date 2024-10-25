package tags

import "lib-transport/ptransport"

func AccessibleRoles() map[string][]string {
	rls := map[string][]string{}
	rls[ptransport.Tags_Put_FullMethodName] = []string{}
	rls[ptransport.Tags_List_FullMethodName] = []string{}
	rls[ptransport.Tags_Search_FullMethodName] = []string{}
	rls[ptransport.Tags_Get_FullMethodName] = []string{}
	return rls
}
