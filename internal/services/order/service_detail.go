package order

import (
    "github.com/xinliangnote/go-gin-api/internal/pkg/core"
    "github.com/xinliangnote/go-gin-api/internal/repository/mysql"
    "github.com/xinliangnote/go-gin-api/internal/repository/mysql/order"
)

type SearchOneData struct {
    Id int32 // 订单ID
}

func (s *service) Detail(ctx core.Context, searchOneData *SearchOneData) (info *order.Order, err error) {

    qb := order.NewQueryBuilder()
    qb.WhereIsDeleted(mysql.EqualPredicate, -1)

    if searchOneData.Id != 0 {
        qb.WhereId(mysql.EqualPredicate, searchOneData.Id)
    }

    info, err = qb.QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
    if err != nil {
        return nil, err
    }

    return
}
