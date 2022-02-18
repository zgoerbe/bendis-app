package middleware

import (
	"github.com/zgoerbe/bendis"
	"myapp/data"
)

type Middleware struct {
	App    *bendis.Bendis
	Models data.Models
}
