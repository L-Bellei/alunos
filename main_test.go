package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/L-Bellei/api-gin-rest/controllers"
	"github.com/L-Bellei/api-gin-rest/database"
	"github.com/L-Bellei/api-gin-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.TestMode)
	database.ConectarBancoDeDados()

	router := gin.Default()
	return router
}

func TestBuscaListagemDeTodosAlunos(t *testing.T) {
	r := SetupDasRotasDeTeste()

	r.GET("/alunos", controllers.ExibeTodosAlunos)

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/alunos", nil)

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Não retornou status 200")

	var alunos []models.Aluno
	json.Unmarshal(resp.Body.Bytes(), &alunos)

	assert.GreaterOrEqual(t, len(alunos), 1, "Não retornou nenhum aluno")
}
