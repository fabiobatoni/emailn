package endpoints

import (
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaings, err := h.CampaignService.Repository.Get()
	return campaings, 200, err
}
