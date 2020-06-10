/**
* @Author: sui_liut@163.com
* @Date: 2020/5/9 15:51
 */

package controller

type CommonResp struct {
	//Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type ContainerListReq struct {
	All     bool              `form:"all"`
	Limit   int               `form:"limit"`
	Size    bool              `form:"size"`
	Filters map[string]string `form:"filters"`
}

type ContainerListResp struct {
	NameSpace string `json:"name_space"`
	PodName   string `json:"pod_name"`
	PodUID    string `json:"pod_uid"`
}

type ContainerRemoveReq struct {
	//Id    string `uri:"id" binding:"required,id"`
	V     bool `form:"v"`
	Force bool `form:"force"`
	Link  bool `form:"link"`
}

type ContainerLogsReq struct {
	Follow     bool   `form:"follow"`
	Stdout     bool   `form:"stdout"`
	Stderr     bool   `form:"stderr"`
	Since      int64  `form:"since"`
	Until      int64  `form:"until"`
	Timestamps bool   `form:"timestamps"`
	Tail       string `form:"tail"`
}
