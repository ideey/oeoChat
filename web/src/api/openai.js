import service from '@/utils/request'

// @Tags Openai
// @Summary 创建Openai
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Openai true "创建Openai"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /myOpenai/createOpenai [post]
export const createOpenai = (data) => {
  return service({
    url: '/myOpenai/createOpenai',
    method: 'post',
    data
  })
}

// @Tags Openai
// @Summary 删除Openai
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Openai true "删除Openai"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /myOpenai/deleteOpenai [delete]
export const deleteOpenai = (data) => {
  return service({
    url: '/myOpenai/deleteOpenai',
    method: 'delete',
    data
  })
}

// @Tags Openai
// @Summary 删除Openai
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Openai"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /myOpenai/deleteOpenai [delete]
export const deleteOpenaiByIds = (data) => {
  return service({
    url: '/myOpenai/deleteOpenaiByIds',
    method: 'delete',
    data
  })
}

// @Tags Openai
// @Summary 更新Openai
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Openai true "更新Openai"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /myOpenai/updateOpenai [put]
export const updateOpenai = (data) => {
  return service({
    url: '/myOpenai/updateOpenai',
    method: 'put',
    data
  })
}

// @Tags Openai
// @Summary 用id查询Openai
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Openai true "用id查询Openai"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /myOpenai/findOpenai [get]
export const findOpenai = (params) => {
  return service({
    url: '/myOpenai/findOpenai',
    method: 'get',
    params
  })
}

// @Tags Openai
// @Summary 分页获取Openai列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Openai列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /myOpenai/getOpenaiList [get]
export const getOpenaiList = (params) => {
  return service({
    url: '/myOpenai/getOpenaiList',
    method: 'get',
    params
  })
}
