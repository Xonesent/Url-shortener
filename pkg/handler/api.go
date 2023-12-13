package handler

import (
	"net/http"

	url_shortener "url-shortener"

	"github.com/gin-gonic/gin"
)

func (h *Handler) send_url(c *gin.Context) {
	var input url_shortener.Link

	if err := c.BindJSON(&input); err != nil {
		New_Error_Response(c, http.StatusBadRequest, "invalid input body error")
		return
	}

	if err := url_shortener.Validate_Base_URL(&input); err != nil {
		New_Error_Response(c, http.StatusBadRequest, "validation error")
		return
	}

	short_url, err := h.services.Url_api.Create_Short_URL(c.Request.Context(), &input)
	if err != nil {
		New_Error_Response(c, http.StatusBadRequest, "impossible to create short url")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Short_URL": short_url,
	})
}

func (h *Handler) get_url(c *gin.Context) {
	var input url_shortener.Link
	input.Short_URL = c.Param("short_url")

	if err := url_shortener.Validate_Short_URL(&input); err != nil {
		New_Error_Response(c, http.StatusBadRequest, "validation error")
		return
	}

	base_url, err := h.services.Url_api.Get_Base_URL(c.Request.Context(), &input)
	if err != nil {
		New_Error_Response(c, http.StatusBadRequest, "impossible to get base url")
		return
	}

	if base_url == "" {
		New_Error_Response(c, http.StatusBadRequest, "empty created short url")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Base_URL": base_url,
	})
}
