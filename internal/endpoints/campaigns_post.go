package endpoints

import (
	internalerrors "emailn/internal/InternalErrors"
	"emailn/internal/contract"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) {
	var request contract.NewCampaign
	render.DecodeJSON(r.Body, &request)
	id, err := h.CampaignService.Create(request)

	if err != nil {
		if errors.Is(err, internalerrors.ErrInternal) {
			render.Status(r, 500)
		} else {
			render.Status(r, 400)
		}
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, 201)
	render.JSON(w, r, map[string]string{"id": id})
}
