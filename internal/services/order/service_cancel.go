package order

import (
    "fmt"

    "github.com/xinliangnote/go-gin-api/internal/pkg/core"
    "github.com/xinliangnote/go-gin-api/internal/repository/mysql"
    "github.com/xinliangnote/go-gin-api/internal/repository/mysql/order"
)

func (s *service) Cancel(ctx core.Context, id int32, status int32, OrderFee int32) (err error) {
    data := map[string]interface{}{
        "status":       status,
        "order_fee":    OrderFee,
        "updated_user": ctx.SessionUserInfo().UserName,
    }
    fmt.Println(ctx.SessionUserInfo())

    qb := order.NewQueryBuilder()
    qb.WhereId(mysql.EqualPredicate, id)
    err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
    if err != nil {
        return err
    }
    return
}
