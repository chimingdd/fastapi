package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	sdkm "github.com/iimeta/fastapi-sdk/model"
)

// Create image接口请求参数
type GenerationsReq struct {
	g.Meta `path:"/generations" tags:"image" method:"post" summary:"Create image接口"`
	sdkm.ImageRequest
}

// Create image接口响应参数
type GenerationsRes struct {
	g.Meta `mime:"application/json" example:"json"`
}

// Create image edit接口请求参数
type EditsReq struct {
	g.Meta `path:"/edits" tags:"image" method:"post" summary:"Create image edit接口"`
	sdkm.ImageRequest
}

// Create image edit接口响应参数
type EditsRes struct {
	g.Meta `mime:"application/json" example:"json"`
}
