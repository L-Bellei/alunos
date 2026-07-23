package routes

import (
	"github.com/L-Bellei/api-gin-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:id", controllers.ExibeAlunoPorId)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaCPF)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.POST("/alunos-em-lote", controllers.CriaNovosAlunosEmLote)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)

	r.Run(":8080")
}
