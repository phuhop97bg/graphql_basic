package handler

import (
	"backend-food/internal/pkg/domain/domain_model/dto"
	"backend-food/pkg/infrastucture/db"
	"backend-food/pkg/infrastucture/schema"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

type HTTPClientHandler struct {
	Schema *graphql.Schema
}

func NewHTTPClientHandler(db db.Database) *HTTPClientHandler {

	schema := schema.NewClientSchema(db)
	return &HTTPClientHandler{Schema: schema}
}

func (h *HTTPClientHandler) Handle(c *gin.Context) {
	req := dto.BaseRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data := graphql.Do(graphql.Params{
		Context:        c,
		Schema:         *h.Schema,
		RequestString:  req.Query,
		VariableValues: req.Variables,
	})
	code := http.StatusOK
	if len(data.Errors) > 0 {
		code = http.StatusBadRequest
	}
	resp := dto.BaseResponse{
		Status: code,
		Error:  data.Errors,
		Result: data.Data,
	}

	c.JSON(code, resp)
}
