package cat

import (
	"context"
	"net/http"

	_ "github.com/hsxflowers/cat-api/cat"
	"github.com/hsxflowers/cat-api/cat/domain"

	"github.com/hsxflowers/cat-api/exceptions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type CatHandler struct {
	ctx        context.Context
	catService domain.Service
}

func NewCatHandler(ctx context.Context, catService domain.Service) CatHandler {
	return CatHandler{
		ctx,
		catService,
	}
}

// Get
//
//	@Summary		Chamar gato
//	@Description	Endpoint que permite a chamada de um gato de acordo com a tag informada.
//	@Accept			json
//	@Produce		json
//	@Success		200			{object}	[]domain.Catresponse	"OK"
//	@Failure		404			"Cat with tag {:tag} not found"
//	@Failure		500			"internal Server Error"
//	@Router			/cat/{:tag} [get]
func (h *CatHandler) Get(c echo.Context) error {
	tag := c.Param("tag")
	if tag == "" || tag == ":tag" {
		log.Error("handler_get: tag is required.", exceptions.ErrTagIsNotValid)
		return exceptions.New(exceptions.ErrTagIsRequired, nil)
	}

	response, err := h.catService.Get(h.ctx, tag)
	if err != nil {
		log.Error("handler_get: error on get a cat.", err)
		return err
	}

	return c.JSON(http.StatusOK, response)
}

// Create
//
//	@Summary		Criação de gatos.
//	@Description	Endpoint que permite a criação de gatos.
//	@Accept			json
//	@Produce		json
//	@Param			body	body		domain.CatRequest	true	"body"
//	@Success		201		{object}	domain.CatResponse
//	@Failure		422		"Unprocessable Json: Payload enviado com erro de syntax do json"
//	@Failure		400		"Erros de validação ou request indevido"
//	@Failure		500			"internal Server Error"
//	@Router			/channels [post]
func (h *CatHandler) Create(c echo.Context) error {
	req := new(domain.CatRequest)

	if err := c.Bind(req); err != nil {
		log.Error("handler_create: error marshal cat", err)
		return exceptions.New(exceptions.ErrBadData, err)
	}

	if err := req.Validate(); err != nil {
		log.Error("handler_create: error on create cat", err)
		return exceptions.New(err, nil)
	}

	err := h.catService.Create(h.ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, nil)
}
