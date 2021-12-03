package handlers

import (
	"github.com/zgoerbe/bendis"
	"myapp/data"
	"net/http"
	"time"
)

type Handlers struct {
	App    *bendis.Bendis
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	//err := h.App.Render.Page(w, r, "home", nil, nil)
	defer h.App.LoadTime(time.Now())
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
