package tests

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/goccy/go-json"
	"net/http"
	"net/http/httptest"
	"persia_atlas/server/models"
	"testing"
)

func refreshBrandTable() {
	server.DB.Migrator().DropTable(models.Brand{})
	server.DB.AutoMigrate(models.Brand{})
}

func TestBrandCreate(t *testing.T) {
	refreshBrandTable()
	gin.SetMode(gin.TestMode)
	r := GetNewRouter()
	server.Router = r
	url := "/api/products/brands"
	r.POST(url, brandController.CreateBrand())

	brandTitle := "test brand"
	request := map[string]any{
		"title": brandTitle,
	}
	jsonValue, _ := json.Marshal(request)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonValue))

	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	response := make(map[string]any)
	err := json.Unmarshal([]byte(recorder.Body.String()), &response)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}
	assert.Equal(t, response["title"], brandTitle)
}
