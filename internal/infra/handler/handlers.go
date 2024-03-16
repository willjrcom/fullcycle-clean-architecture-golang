package handlerimpl

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/willjrcom/fullcycle-clean-architecture-golang/bootstrap/server"
	"github.com/willjrcom/fullcycle-clean-architecture-golang/internal/domain"
	"github.com/willjrcom/fullcycle-clean-architecture-golang/internal/usecase"
	jsonpkg "github.com/willjrcom/fullcycle-clean-architecture-golang/pkg/json"
)

type handlerOrderImpl struct {
	s *usecase.Service
}

func NewHandlerOrder(service *usecase.Service) *server.Handler {
	c := chi.NewRouter()

	h := &handlerOrderImpl{
		s: service,
	}

	route := "/orders"

	c.With().Group(func(c chi.Router) {
		c.Post("/new", h.handlerNewOrder)
		c.Get("/all", h.handlerGetAllOrders)
	})

	return server.NewHandler(route, c)
}

func (h *handlerOrderImpl) handlerNewOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	order := domain.OrderCommonAttributes{}
	jsonpkg.ParseBody(r, &order)

	if id, err := h.s.NewOrder(ctx, order); err != nil {
		jsonpkg.ResponseJson(w, r, http.StatusInternalServerError, jsonpkg.Error{Message: err.Error()})
	} else {
		jsonpkg.ResponseJson(w, r, http.StatusCreated, jsonpkg.HTTPResponse{Data: id})
	}
}

func (h *handlerOrderImpl) handlerGetAllOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if orders, err := h.s.ListOrders(ctx); err != nil {
		jsonpkg.ResponseJson(w, r, http.StatusInternalServerError, jsonpkg.Error{Message: err.Error()})
	} else {
		jsonpkg.ResponseJson(w, r, http.StatusCreated, jsonpkg.HTTPResponse{Data: orders})
	}
}
