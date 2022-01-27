package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/raulinoneto/url-shortener/logger"
	"net/http"
)

const (
	UrlIdParamKey = "url"
)


type Repository interface {
	Get(ctx context.Context, id string) (string, error)
	Shorten(ctx context.Context, u string) (string, error)
	List(ctx context.Context) ([]string, error)
}

type (
	UrlRequest struct {
		Url string `json:"url" validate:"required"`
	}
	UrlResponse struct {
		MainUrl  string `json:"main_url,omitempty"`
		ShortUrl string `json:"short_url,omitempty"`
	}
)

type ShortenerHandler struct {
	repo Repository
}

func (sh *ShortenerHandler) Shorten(c echo.Context) error {
	ctx := c.Request().Context()
	ctx, log := logger.FromContext(ctx)

	log.Info("Start to short an URL")

	req := UrlRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	sortUrl, err := sh.repo.Shorten(ctx, req.Url)
	if err != nil {
		log.Errorf("Error when shorts an URL: %w", err)
		return err
	}

	return c.JSON(http.StatusCreated, &UrlResponse{
		MainUrl:  req.Url,
		ShortUrl: sortUrl,
	})
}


func (sh *ShortenerHandler) Redirect(c echo.Context) error {
	ctx := c.Request().Context()
	ctx, log := logger.FromContext(ctx)

	log.Info("Start to redirect an URL")

	id := c.Param(UrlIdParamKey)
	sortUrl, err := sh.repo.Get(ctx, id)
	if err != nil {
		log.Errorf("Error when get a short URL: %w", err)
		return err
	}

	return c.Redirect(http.StatusTemporaryRedirect, sortUrl)
}


func (sh *ShortenerHandler) List(c echo.Context) error {
	ctx := c.Request().Context()
	ctx, log := logger.FromContext(ctx)

	log.Info("Start to List short URLs")

	sortUrls, err := sh.repo.List(ctx)
	if err != nil {
		log.Errorf("Error when list short URLs: %w", err)
		return err
	}

	return c.JSON(http.StatusOK, sortUrls)
}


