package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/L-Bellei/api-gin-rest/controllers"
	"github.com/L-Bellei/api-gin-rest/database"
	"github.com/L-Bellei/api-gin-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	database.ConectarBancoDeDados()

	router := gin.Default()
	return router
}

func CriaAlunoMock() {
	aluno := models.Aluno{
		Nome: "Fernanda Ribeiro",
		RG:   "58275192",
		CPF:  "98979660130",
	}

	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno

	database.DB.Delete(&aluno, ID)
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

func TestBuscaUmAlunoPeloID(t *testing.T) {
	r := SetupDasRotasDeTeste()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r.GET("/alunos/:id", controllers.ExibeAlunoPorId)

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/alunos/"+strconv.Itoa(ID), nil)

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Não retornou status 200")

	var aluno models.Aluno
	json.Unmarshal(resp.Body.Bytes(), &aluno)

	assert.Equal(t, aluno.ID, uint(ID), "Não retornou o aluno esperado")
	assert.Equal(t, aluno.Nome, "Fernanda Ribeiro", "Não retornou o nome esperado")
	assert.Equal(t, aluno.RG, "58275192", "Não retornou o RG esperado")
	assert.Equal(t, aluno.CPF, "98979660130", "Não retornou o CPF esperado")
}

func TestBuscaAlunoPorCPF(t *testing.T) {
	r := SetupDasRotasDeTeste()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r.GET("/alunos/cpf/:cpf", controllers.BuscaCPF)

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/alunos/cpf/98979660130", nil)

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Não retornou status 200")

	var aluno models.Aluno
	json.Unmarshal(resp.Body.Bytes(), &aluno)

	assert.Equal(t, aluno.ID, uint(ID), "Não retornou o aluno esperado")
	assert.Equal(t, aluno.Nome, "Fernanda Ribeiro", "Não retornou o nome esperado")
	assert.Equal(t, aluno.RG, "58275192", "Não retornou o RG esperado")
	assert.Equal(t, aluno.CPF, "98979660130", "Não retornou o CPF esperado")
}

func TestDeletaAluno(t *testing.T) {
	r := SetupDasRotasDeTeste()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r.DELETE("/alunos/:id", controllers.DeletaAluno)

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/alunos/"+strconv.Itoa(ID), nil)

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Não retornou status 200")
}

func TestEditaAluno(t *testing.T) {
	r := SetupDasRotasDeTeste()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	var aluno = models.Aluno{
		Nome: "Aluno Editado",
		CPF:  "98979660130",
		RG:   "123456789",
	}

	valorJson, _ := json.Marshal(aluno)

	r.PATCH("/alunos/:id", controllers.EditaAluno)
	req, _ := http.NewRequest("PATCH", "/alunos/"+strconv.Itoa(ID), bytes.NewBuffer(valorJson))
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	var alunoEditado models.Aluno
	json.Unmarshal(resp.Body.Bytes(), &alunoEditado)

	assert.Equal(t, http.StatusOK, resp.Code, "Não retornou status 200")
	assert.Equal(t, aluno.Nome, alunoEditado.Nome, "Não retornou o nome esperado")
	assert.Equal(t, aluno.CPF, alunoEditado.CPF, "Não retornou o CPF esperado")
	assert.Equal(t, aluno.RG, alunoEditado.RG, "Não retornou o RG esperado")
}
