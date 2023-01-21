package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/persistance"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/handler"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
)

func (r *Router) InitEvent(c *ent.Client) {
	repo := persistance.NewEvent(c)
	uc := usecase.NewEvent(repo)
	h := handler.NewEvent(uc)

	r.mux.Route("/events", func(r chi.Router) {
		r.Post("/", h.CreateNewEvent)
		r.Get("/{id}", h.GetEventById)
		r.Delete("/{id}", h.DeleteEventById)
		r.Patch("/{id}", h.UpdateEventById)
		r.Get("/{id}/admin", h.GetEventAdminById)
		r.Get("/{id}/comments", h.GetEventComments)
		r.Post("/{id}/participants", h.AddNewEventParticipant)
		r.Patch("/{id}/state", h.ChangeEventStatusOfId)
		r.Get("/{id}/users", h.GetEventUsers)
	})
}
