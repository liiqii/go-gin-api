package order

import (
    "net/http"

    "github.com/xinliangnote/go-gin-api/internal/code"
    "github.com/xinliangnote/go-gin-api/internal/pkg/core"
    "github.com/xinliangnote/go-gin-api/internal/pkg/validation"
)

type cancelRequest struct {
    Id       int32 `form:"id"`        // 主键ID
    Status   int32 `form:"status"`    // 订单状态 1:已创建  2:已取消
    OrderFee int32 `form:"order_fee"` // 订单金额(分)
}

type cancelResponse struct {
    Id int32 `json:"id"` // 主键ID
}

// Cancel 取消订单
// @Summary 取消订单
// @Description 取消订单
// @Tags API.order
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body cancelRequest true "请求信息"
// @Success 200 {object} cancelResponse
// @Failure 400 {object} code.Failure
// @Router /api/order/cancel [post]
func (h *handler) Cancel() core.HandlerFunc {
    return func(ctx core.Context) {
        req := new(cancelRequest)
        res := new(cancelResponse)
        if err := ctx.ShouldBindForm(req); err != nil {
            ctx.AbortWithError(core.Error(
                http.StatusBadRequest,
                code.ParamBindError,
                validation.Error(err)).WithError(err),
            )
            return
        }

        err := h.orderService.Cancel(ctx, req.Id, req.Status, req.OrderFee)
        if err != nil {
            ctx.AbortWithError(core.Error(
                http.StatusBadRequest,
                code.AdminUpdateError,
                code.Text(code.AdminUpdateError)).WithError(err),
            )
            return
        }

        res.Id = req.Id
        ctx.Payload(res)
    }
}
