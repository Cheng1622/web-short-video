package pkg

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"golang.org/x/exp/slog"
	"gorm.io/gorm/clause"
)

type Filter struct {
}

var s2f = map[string]func(c any, v ...any) clause.Expression{
	// 	gt：大于（greater than）
	"gt": func(c any, v ...any) clause.Expression { return clause.Gt{Column: c, Value: v[0]} },
	// gte：大于等于（greater than or equal）
	"gte": func(c any, v ...any) clause.Expression { return clause.Gte{Column: c, Value: v[0]} },
	"in":  func(c any, v ...any) clause.Expression { return clause.IN{Column: c, Values: v} },
	// lt:：小于（less than）
	"lt": func(c any, v ...any) clause.Expression { return clause.Lt{} },
	// lte：小于等于（less than or equal）
	"lte": func(c any, v ...any) clause.Expression { return clause.Lte{Column: c, Value: v[0]} },
	// neq：不等于
	"neq": func(c any, v ...any) clause.Expression { return clause.Neq{Column: c, Value: v[0]} },
	// eq：等于（equal）
	"eq":     func(c any, v ...any) clause.Expression { return clause.Eq{Column: c, Value: v[0]} },
	"like":   func(c any, v ...any) clause.Expression { return clause.Like{Column: c, Value: v[0]} },
	"isnull": func(c any, v ...any) clause.Expression { return isnull{Column: c} },
	"offest": func(c any, v ...any) clause.Expression { return offest{Column: c, Value: v} },
}

func New(c *gin.Context, db *gorm.DB) any {
	args := c.Request.URL.Query()
	// args := map[string][]string{"eq|c": {"1"}}
	var res []any
	for k, v := range args {
		ks := strings.SplitN(k, `|`, 2)
		if len(ks) != 2 {
			slog.Info("filter-ks:", ks)
			continue
		}
		if s2f[ks[0]] != nil {
			db.Where(s2f[ks[0]](ks[1], v[0]))
		}
		slog.Info("", k, v)
	}
	err := db.Find(&res).Error
	if err != nil {
		slog.Error("", err)
	}
	return res
}

// Gte greater than or equal to for where
type isnull clause.Eq

func (gte isnull) Build(builder clause.Builder) {
	builder.WriteQuoted(gte.Column)
	builder.WriteString(" is null ")
}

// Gte greater than or equal to for where
type offest clause.Eq

// db.Offset((pageNumber - 1) * pageSize).Limit(pageSize).
func (gte offest) Build(builder clause.Builder) {
	builder.WriteQuoted(gte.Column)
	builder.WriteString(" offest ")
	builder.AddVar(builder, gte.Value, gte.Value)
}
