package main

type RouteJumpReq struct {
	FileId  string `query:"file_id"`
	GroupId string `query:"group_id"`
}

// 老师，前端是不是可能传这两个参数？然后两个参数只会传一个，比如只有Fileid或者只有groupid还是说两个参数都必须有？
// 然后我返回前端一个这样的链接是吧？https://365.kdocs.cn/wiki/l/0lcvDUdngRI7wD/i?from=space
// 现在有两个方案：方案一：直接通过返回302状态码和一个新URL，让浏览器自动重定向到智能文档库界面
// 				方案二：通过提供接口，根据前端传入的fileid和groupid返回对应的智能文档库链接，前端再去请求这个链接
