// Code generated by goctl. DO NOT EDIT.
package types

type BaseDataInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data,omitempty"`
}

type BaseListInfo struct {
	Total uint64 `json:"total"`
	Data  string `json:"data,omitempty"`
}

type BaseMsgResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type SimpleMsg struct {
	Msg string `json:"msg"`
}

type PageInfo struct {
	Current  uint64 `json:"page" validate:"number"`
	PageSize uint64 `json:"pageSize" validate:"number,max=100000"`
}

type IDReq struct {
	Id uint64 `json:"id" validate:"number"`
}

type IDsReq struct {
	Ids []uint64 `json:"ids"`
}

type IDPathReq struct {
	Id uint64 `path:"id"`
}

type UUIDReq struct {
	Id string `json:"id" validate:"len=36"`
}

type UUIDsReq struct {
	Ids []string `json:"ids"`
}

type BaseInfo struct {
	Id        uint64 `json:"id"`
	CreatedAt int64  `json:"createdAt,optional"`
	UpdatedAt int64  `json:"updatedAt,optional"`
}

type BaseUUIDInfo struct {
	Id        string `json:"id"`
	CreatedAt int64  `json:"createdAt,optional"`
	UpdatedAt int64  `json:"updatedAt,optional"`
}
