/**
* @Author: sui_liut@163.com
* @Date: 2020/5/13 15:03
 */

package controller

import (
	"strings"
)

func getmap(m map[string][]string, key string) (map[string]string, bool) {
	dicts := make(map[string]string)
	exist := false
	for k, v := range m {
		if i := strings.IndexByte(k, '['); i >= 1 && k[0:i] == key {
			if j := strings.IndexByte(k[i+1:], ']'); j >= 1 {
				exist = true
				dicts[k[i+1:][:j]] = v[0]
			}
		}
	}
	return dicts, exist
}

func nameSplit(names []string) []ContainerListResp {
	// /k8s_dashboard-metrics-scraper_dashboard-metrics-scraper-dc6947fbf-g9nsn_kubernetes-dashboard_2d52cfd6-2ff6-4342-80a5-d1bc86473f54_0
	// /k8s_POD_dashboard-metrics-scraper-dc6947fbf-g9nsn_kubernetes-dashboard_2d52cfd6-2ff6-4342-80a5-d1bc86473f54_0
	var resp []ContainerListResp
	for _, name := range names {

		if !strings.HasPrefix(name, "/k8s") {
			continue
		}

		name = name[5:]
		name = name[:strings.LastIndex(name, "_")]
		var t ContainerListResp
		t.PodUID = name[strings.LastIndex(name, "_")+1:]

		name = name[:strings.LastIndex(name, "_")]
		name = name[strings.Index(name, "_")+1:]

		t.PodName = name[:strings.LastIndex(name, "_")]
		t.NameSpace = name[strings.Index(name, "_")+1:]
		resp = append(resp, t)
	}

	return resp

}
