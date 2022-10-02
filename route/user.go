package route

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"log"
	"net/http"
	"redfox/db"
	"redfox/model"
	"time"
)

var tracer = otel.Tracer("Router /users")

func User(router *gin.Engine) {

	router.GET("/users", func(g *gin.Context) {
		tCtx, span := tracer.Start(context.Background(), "GET /users")
		defer span.End()

		user := db.GetUser(tCtx)

		sampleWait(tCtx)
		log.Printf("traceId: %s", span.SpanContext().TraceID())

		// TODO: How to securely manage sensitive fields
		g.JSON(http.StatusOK, user)
	})

	router.POST("/users", func(g *gin.Context) {
		tCtx, span := tracer.Start(context.Background(), "POST /users")
		defer span.End()

		var req model.User
		if err := g.BindJSON(&req); err != nil {
			// TODO: suppress error details
			g.JSON(http.StatusBadRequest, err)
			return
		}

		// TODO: Validate input
		//if vErr := validate.Struct(&user); vErr != nil {
		//	g.JSON(http.StatusBadRequest, vErr})
		//	return

		creatingUser := model.User{
			Username: req.Username,
		}

		createdUser := db.CreateUser(creatingUser, tCtx)
		g.JSON(http.StatusCreated, createdUser)
	})
}

func sampleWait(tCtx context.Context) {
	tCtx, span := tracer.Start(tCtx, "Just sleep")
	defer span.End()
	time.Sleep(1 * time.Millisecond)
	span.SetAttributes(attribute.String("username", "user001"))
	span.SetStatus(codes.Error, "Add exception here")
}
