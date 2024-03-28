package order

import (
    "github.com/xinliangnote/go-gin-api/internal/pkg/core"
    "github.com/xinliangnote/go-gin-api/internal/repository/mysql/order"
)

type CreateOrderData struct {
    OrderNo  string // 订单号
    OrderFee int32  // 订单金额(分)
}

func (s *service) Create(ctx core.Context, orderData *CreateOrderData) (id int32, err error) {
    model := order.NewModel()
    model.OrderNo = orderData.OrderNo
    model.OrderFee = orderData.OrderFee
    model.CreatedUser = ctx.SessionUserInfo().UserName
    model.Status = 1
    model.IsDeleted = -1

    id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
    if err != nil {
        return 0, err
    }
    return
}
