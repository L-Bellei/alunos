package controllers

import (
	"net/http"

	"github.com/L-Bellei/api-gin-rest/database"
	"github.com/L-Bellei/api-gin-rest/models"
	"github.com/gin-gonic/gin"
)

func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno

	database.DB.Find(&alunos)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno

	database.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func ExibeAlunoPorId(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno

	if err := database.DB.First(&aluno, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func BuscaCPF(c *gin.Context) {
	cpf := c.Params.ByName("cpf")
	var aluno models.Aluno

	if err := database.DB.Where("cpf = ?", cpf).First(&aluno).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidaDadosDeAlunos(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func CriaNovosAlunosEmLote(c *gin.Context) {
	var alunos []models.Aluno

	if err := c.ShouldBindJSON(&alunos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	for _, aluno := range alunos {
		if err := models.ValidaDadosDeAlunos(&aluno); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro": err.Error(),
			})
			return
		}
	}

	if result := database.DB.CreateInBatches(&alunos, 10); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, alunos)
}

func EditaAluno(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno

	if err := database.DB.First(&aluno, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "aluno não encontrado",
		})
		return
	}

	var alunoEdicao models.Aluno
	if err := c.ShouldBindJSON(&alunoEdicao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidaDadosDeAlunos(&alunoEdicao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Model(&aluno).UpdateColumns(alunoEdicao).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno

	result := database.DB.Delete(&aluno, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Aluno deletado com sucesso",
	})
}
