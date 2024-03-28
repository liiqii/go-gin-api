package order

import (
    "net/http"

    "github.com/xinliangnote/go-gin-api/internal/code"
    "github.com/xinliangnote/go-gin-api/internal/pkg/core"
    "github.com/xinliangnote/go-gin-api/internal/pkg/validation"
    "github.com/xinliangnote/go-gin-api/internal/services/order"
)

type createRequest struct {
    OrderNo  string `form:"order_no" binding:"required"`  // 订单号
    OrderFee int32  `form:"order_fee" binding:"required"` // 订单金额(分)
}

type createResponse struct {
    Id int32 `json:"id"` // 主键ID
}

// Create 创建订单
// @Summary 创建订单
// @Description 创建订单
// @Tags API.order
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/order/create [post]
func (h *handler) Create() core.HandlerFunc {
    return func(c core.Context) {
        req := new(createRequest)
        res := new(createResponse)
        if err := c.ShouldBindForm(req); err != nil {
            c.AbortWithError(core.Error(
                http.StatusBadRequest,
                code.ParamBindError,
                validation.Error(err)).WithError(err),
            )
            return
        }

        createData := new(order.CreateOrderData)
        createData.OrderNo = req.OrderNo
        createData.OrderFee = req.OrderFee

        id, err := h.orderService.Create(c, createData)
        if err != nil {
            c.AbortWithError(core.Error(
                http.StatusBadRequest,
                code.OrderCreateError,
                code.Text(code.OrderCreateError)).WithError(err),
            )
            return
        }

        res.Id = id
        c.Payload(res)
    }
}
