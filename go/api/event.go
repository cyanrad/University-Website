package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/cyanrad/university/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createEventRequest struct {
	Name        string    `json:"name" binding:"required,alphanum"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
	Link        string    `json:"link" binding:"required"`
	Description string    `json:"description" binding:"required"`
}

func (server *Server) createEvent(ctx *gin.Context) {
	// >> reading request data into req var
	var req createEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// >> Creating sql query paramteres
	arg := db.CreateEventParams{
		Name:        req.Name,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Link:        req.Link,
		Description: req.Description,
	}

	// >> inserting new event into db
	event, err := server.store.CreateEvent(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() { // check and respond @ sql violation
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// >> send response with created event
	ctx.JSON(http.StatusOK, event)
}

type getEventRequest struct {
	id int64 `uri:"id" binding:"required,min=0"`
}

func (server *Server) getEvent(ctx *gin.Context) {
	var req getEventRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	event, err := server.store.GetEvent(ctx, int32(req.id))
	if err != nil {
		// if the error's cause is having no rows
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, event)
}
