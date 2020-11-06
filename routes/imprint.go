package routes

import (
	conf "github.com/muety/wakapi/config"
	"github.com/muety/wakapi/models"
	"github.com/muety/wakapi/models/view"
	"github.com/muety/wakapi/services"
	"net/http"
)

type ImprintHandler struct {
	config       *conf.Config
	keyValueSrvc *services.KeyValueService
}

func NewImprintHandler(keyValueService *services.KeyValueService) *ImprintHandler {
	return &ImprintHandler{
		config:       conf.Get(),
		keyValueSrvc: keyValueService,
	}
}

func (h *ImprintHandler) GetImprint(w http.ResponseWriter, r *http.Request) {
	if h.config.IsDev() {
		loadTemplates()
	}

	text := "failed to load content"
	if data, err := h.keyValueSrvc.GetString(models.ImprintKey); err == nil {
		text = data.Value
	}

	templates[conf.ImprintTemplate].Execute(w, h.buildViewModel(r).WithHtmlText(text))
}

func (h *ImprintHandler) buildViewModel(r *http.Request) *view.ImprintViewModel {
	return &view.ImprintViewModel{
		Success: r.URL.Query().Get("success"),
		Error:   r.URL.Query().Get("error"),
	}
}
