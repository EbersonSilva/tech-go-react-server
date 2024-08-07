package api

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"sync"

	"github.com/rocketseat-education/semana-tech-go-react-server/internal/store/pgstore"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5"
) 

type apiHandler struct{
	q *pgstore.Queries
	r *chi.Mux
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	 a := apiHandler{
		q: q,
	 }

	 r := chi.NewRouter()
	 r.Use(middleware.RequestId, middleware.Recoverer, middleware.Logger)

	 r.Route("/api", func(r chi.Router){
		r.Route("/rooms", func(r chi.Router){
			r.Post("/", h http.HandlerFunc)

			r.Route("/{room_id}/messages", func(r chi.Router){ //Parametro de caminho 
				r.Post("/", a.handleCreateMessage)
				r.Get("/", a.handleGetRoomMessages)

				r.Route("/{message_id}", func (r chi.Router){
					r.Get("/", a.handleGetRoomMessage)
					r.Patch("/react", a.handleReactToMessage)
					r.Delete("/react", a.handleRemoveReactFromMessage)
					r.Patch("/answer", a.handleMarkMessageAsAnswered)
				})
			})
		})
	 })

	 a.r = r
	 return a
}
// Criação de metodos
func (h apiapiHAndler) handleCreateRoom(w http.ResponseWriter, r *http.Request){}
func (h apiapiHAndler) handleGetRooms(w http.ResponseWriter, r *http.Request){}
func (h apiapiHAndler) handleCreateRoomMessage(w http.ResponseWriter, r *http.Request){}
func (h apiapiHAndler) handleGetRoomMessages(w http.ResponseWriter, r *http.Request){}
func (h apiapiHAndler) handleGetRoomMessage(w http.ResponseWriter, r *http.Request){}
func (h apiapiHAndler) handleReactToMessage(w http.ResponseWriter, r *http.Request){}
func (h apiapiHAndler) handleRemoveReactFromMessage(w http.ResponseWriter, r *http.Request){}
func (h apiapiHAndler) handleMarkMessageAsAnswered(w http.ResponseWriter, r *http.Request){}