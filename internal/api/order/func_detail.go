package order

import (
    "net/http"

    "github.com/xinliangnote/go-gin-api/internal/code"
    "github.com/xinliangnote/go-gin-api/internal/pkg/core"
    "github.com/xinliangnote/go-gin-api/internal/pkg/validation"
    "github.com/xinliangnote/go-gin-api/internal/services/order"
)

type detailRequest struct {
    Id int32 `uri:"id"` // HashID
}

type detailResponse struct {
    Id       int32  `json:"id"`        // id
    OrderNo  string `json:"order_no"`  // 订单号
    OrderFee int32  `json:"order_fee"` // 订单金额(分)
    Status   int32  `json:"status"`    // 订单状态 1:已创建  2:已取消
}

// Detail 订单详情
// @Summary 订单详情
// @Description 订单详情
// @Tags API.order
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body detailRequest true "请求信息"
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/order/{id} [get]
func (h *handler) Detail() core.HandlerFunc {
    return func(ctx core.Context) {
        req := new(detailRequest)
        res := new(detailResponse)

        if err := ctx.ShouldBindURI(req); err != nil {
            ctx.AbortWithError(core.Error(
                http.StatusBadRequest,
                code.ParamBindError,
                validation.Error(err)).WithError(err),
            )
            return
        }
        searchOneData := new(order.SearchOneData)
        searchOneData.Id = req.Id

        info, err := h.orderService.Detail(ctx, searchOneData)
        if err != nil {
            ctx.AbortWithError(core.Error(
                http.StatusBadRequest,
                code.AdminDetailError,
                code.Text(code.AdminDetailError)).WithError(err),
            )
            return
        }

        res.Id = info.Id
        res.OrderNo = info.OrderNo
        res.OrderFee = info.OrderFee
        res.Status = info.Status
        ctx.Payload(res)
    }
}
