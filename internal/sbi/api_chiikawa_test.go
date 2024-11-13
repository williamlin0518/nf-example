package sbi_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andy89923/nf-example/internal/sbi"
	"github.com/andy89923/nf-example/pkg/factory"
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func Test_getChiikawaRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockCtrl := gomock.NewController(t)
	nfApp := sbi.NewMocknfApp(mockCtrl)
	nfApp.EXPECT().Config().Return(&factory.Config{
		Configuration: &factory.Configuration{
			Sbi: &factory.Sbi{
				Port: 8000,
			},
		},
	}).AnyTimes()
	server := sbi.NewServer(nfApp, "")

	t.Run("Yee~ya~ha!", func(t *testing.T) {
		const EXPECTED_STATUS = http.StatusOK
		const EXPECTED_BODY = "Yee~ya~ha!"

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)

		var err error
		ginCtx.Request, err = http.NewRequest("GET", "/chiikawa", nil)
		if err != nil {
			t.Errorf("Failed to create request: %s", err)
			return
		}

		server.YeeYaHa(ginCtx)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if httpRecorder.Body.String() != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})

	t.Run("Create Character", func(t *testing.T) {
		const EXPECTED_STATUS = http.StatusOK
		const CHARACTER_NAME = "Usagi"
		const EXPECTED_BODY = "Hello " + CHARACTER_NAME + "!"

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)

		var err error
		jsonBody := `{"name":"` + CHARACTER_NAME + `"}`
		ginCtx.Request, err = http.NewRequest("POST", "/chiikawa/character", strings.NewReader(jsonBody))
		if err != nil {
			t.Errorf("Failed to create request: %s", err)
			return
		}

		server.CreateChiikawaCharacter(ginCtx)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if httpRecorder.Body.String() != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})

	t.Run("List Character", func(t *testing.T) {
		const EXPECTED_STATUS = http.StatusOK
		const EXPECTED_BODY = "Chiikawa Character List : Usagi"

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)

		var err error
		ginCtx.Request, err = http.NewRequest("GET", "/chiikawa/list", nil)
		if err != nil {
			t.Errorf("Failed to create request: %s", err)
			return
		}

		server.ListChiikawaCharacter(ginCtx)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if httpRecorder.Body.String() != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})
}
