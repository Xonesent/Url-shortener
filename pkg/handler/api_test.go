package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	url_shortener "url-shortener"
	"url-shortener/pkg/service"
	mock_service "url-shortener/pkg/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestHandler_send_url(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUrl_api, link *url_shortener.Link)

	tests := []struct {
		name               string
		input_body         string
		input              url_shortener.Link
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:       "ok",
			input_body: `{"base_url": "https://github.com/Xonesent", "short_url": ""}`,
			input: url_shortener.Link{
				Base_URL:  "https://github.com/Xonesent",
				Short_URL: "",
			},
			mockBehavior: func(r *mock_service.MockUrl_api, link *url_shortener.Link) {
				r.EXPECT().Create_Short_URL(gomock.Any(), link).Return("short_url", nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:       "validation_error",
			input_body: `{"base_url": "//github.com/Xonesent", "short_url": ""}`,
			input: url_shortener.Link{
				Base_URL:  "//github.com/Xonesent",
				Short_URL: "",
			},
			mockBehavior:       func(r *mock_service.MockUrl_api, link *url_shortener.Link) {},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repos := mock_service.NewMockUrl_api(c)
			test.mockBehavior(repos, &test.input)

			services := &service.Service{Url_api: repos}
			handler := Handler{services}

			r := gin.New()
			r.POST("/api/send_url", handler.send_url)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/send_url", bytes.NewBufferString(test.input_body))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
		})
	}
}

func TestHandler_get_url(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUrl_api)

	tests := []struct {
		name               string
		input              string
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:  "ok",
			input: "9fWC_NBfmi",
			mockBehavior: func(r *mock_service.MockUrl_api) {
				r.EXPECT().Get_Base_URL(gomock.Any(), &url_shortener.Link{Short_URL: "9fWC_NBfmi"}).Return("https://github.com/Xonesent", nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "validation_error",
			input:              "----------",
			mockBehavior:       func(r *mock_service.MockUrl_api) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:  "service_error",
			input: "9fWC_NBfmi",
			mockBehavior: func(r *mock_service.MockUrl_api) {
				r.EXPECT().Get_Base_URL(gomock.Any(), &url_shortener.Link{Short_URL: "9fWC_NBfmi"}).Return("", fmt.Errorf("some error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:  "not_found",
			input: "9fWC_NBfmi",
			mockBehavior: func(r *mock_service.MockUrl_api) {
				r.EXPECT().Get_Base_URL(gomock.Any(), &url_shortener.Link{Short_URL: "9fWC_NBfmi"}).Return("", nil)
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repos := mock_service.NewMockUrl_api(c)

			services := &service.Service{Url_api: repos}
			handler := Handler{services}

			r := gin.New()
			r.GET("/get_url", handler.get_url)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/get_url/"+test.input, nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
		})
	}
}
