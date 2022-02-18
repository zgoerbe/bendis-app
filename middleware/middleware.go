package middleware

import (
	"github.com/zgoerbe/bendis"
	"myapp_pt2/data"
)

type Middleware struct {
	App    *bendis.Bendis
	Models data.Models
}
