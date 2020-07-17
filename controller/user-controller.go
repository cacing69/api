package controller

import (
	"fmt"
	"math"
	"strconv"

	"github.com/cacing69/api/conf"
	"github.com/cacing69/api/entity"
	. "github.com/cacing69/api/lib"
	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber"
	"github.com/rocketlaunchr/dbq/v2"
)

func UserIndex(c *fiber.Ctx) {
	_page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		c.Next(err)
		return
	}

	_limit, err := strconv.Atoi(c.Query("limit", "20"))
	if err != nil {
		c.Next(err)
		return
	}

	args := make([]interface{}, 0)

	_offset := (_page - 1) * _limit

	_where := "where 1=1"

	if c.Query("name") != "" {
		_where += " AND user_name like ? "
		args = append(args, fmt.Sprintf("%%%s%%", c.Query("name")))
	}

	_table := "m_tester"

	_query_count := fmt.Sprintf("SELECT count(*) count FROM %s %s", _table, _where)
	_query_fetch := fmt.Sprintf("SELECT * FROM %s %s LIMIT %d OFFSET %d", _table, _where, _limit, _offset)

	_run_count := dbq.MustQ(c.Context(), conf.DB, _query_count, entity.GenericSingleOption(), args...)
	_run_row := dbq.MustQ(c.Context(), conf.DB, _query_fetch, entity.TesterMultiOption(), args...)

	_count := _run_count.(*entity.T).Count

	d := float64(_count) / float64(_limit)
	_page_count := int(math.Ceil(d))

	c.JSON(Res{
		Data:    _run_row,
		Message: "index user",
		Status:  true,
		Code:    200,
		Meta: M{
			"count":      _count,
			"limit":      _limit,
			"page_count": _page_count,
		},
	})
}

func UserShow(c *fiber.Ctx) {

}

// func UserIndexOrm(c *fiber.Ctx) {
// 	db := orm.NewOrm()

// 	// Get a QuerySeter object. User is table name
// 	// qs := o.QueryTable("m_user")

// 	// Can also use object as table name

// 	var data []entity.Tester
// 	q := db.QueryTable("m_tester")

// 	count, _ := q.Count()

// 	q = q.Limit(20)

// 	q.All(&data)

// 	c.JSON(Res{
// 		Data: data,
// 		Meta: M{
// 			"count": count,
// 		},
// 	})

// }

func UserStore(c *fiber.Ctx) {
	rows := []interface{}{
		[]interface{}{"1", "2kkkk", "3"},
	}

	columns := []string{
		"user_name",
		"user_email",
		"user_password",
	}

	statement := dbq.INSERT("m_user", columns, len(rows), dbq.MySQL)

	res, err := dbq.E(c.Context(), conf.DB, statement, nil, rows)

	if err != nil {
		c.Next(err)
		return
	}

	spew.Dump(res.LastInsertId())
	i, _ := res.LastInsertId()
	c.JSON(M{
		"id": i,
	})
}

func UserUpdate(c *fiber.Ctx) {

}

func UserDelete(c *fiber.Ctx) {

}
