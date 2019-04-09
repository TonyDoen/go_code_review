package common

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
)

type TRenderJson struct {
	ReturnCode    int                    `json:"returnCode"`
	ReturnMsg     string                 `json:"returnMsg"`
	ReturnUserMsg string                 `json:"returnUserMsg"`
	Data          map[string]interface{} `json:"data"`
}

func RenderJsonSuccess(ctx *gin.Context, data map[string]interface{}) {
	renderJson := TRenderJson{0, "success", "成功", data}
	ctx.JSON(http.StatusOK, renderJson)
	ctx.Set("render", renderJson)
	return
}

func RenderJsonFail(ctx *gin.Context, err error) {
	var renderJson TRenderJson

	switch errors.Cause(err).(type) {

	case Error:
		renderJson.ReturnMsg = errors.Cause(err).(Error).ErrorMsg
		renderJson.ReturnUserMsg = errors.Cause(err).(Error).ErrorUserMsg
		renderJson.ReturnCode = errors.Cause(err).(Error).ErrorCode

	case validator.ValidationErrors:
		renderJson.ReturnMsg = "ParamError"
		renderJson.ReturnUserMsg = "参数错误"
		renderJson.ReturnCode = 1000

		var validateData []interface{}
		for _, err := range errors.Cause(err).(validator.ValidationErrors) {
			validateData = append(validateData, map[string]interface{}{"message": "params check error", "tag": err.Tag, "field": err.Field})
		}

		renderJson.Data = map[string]interface{}{"validateData": validateData}

	default:
		renderJson.ReturnMsg = errors.Cause(err).Error()
		renderJson.ReturnUserMsg = ""
		renderJson.ReturnCode = -1
	}
	ctx.JSON(http.StatusOK, renderJson)
	ctx.Set("render", renderJson)

	// print error stack
	StackLogger(ctx, err)
	return
}
