package sbi

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ChiikawaCharacterList struct {
	Characters []string
}

var chiiikawaList = &ChiikawaCharacterList{
	Characters: []string{},
}

func (s *Server) getChiikawaRoute() []Route {
	return []Route{
		{
			Name:    "Yee~ya~ha!",
			Method:  http.MethodGet,
			Pattern: "/",
			APIFunc: s.YeeYaHa,
			// Use
			// curl -X GET http://127.0.0.163:8000/chiikawa/ -w "\n"
		},
		{
			Name:    "Chiikawa Character List",
			Method:  http.MethodGet,
			Pattern: "/list",
			APIFunc: s.ListChiikawaCharacter,
			// Use
			// curl -X GET http://127.0.0.163:8000/chiikawa/list -w "\n"
		},
		{
			Name:    "Create Chiikawa Character",
			Method:  http.MethodPost,
			Pattern: "/character",
			APIFunc: s.CreateChiikawaCharacter,
			// Use
			// curl -X POST http://127.0.0.163:8000/chiikawa/character -d '{"name": "Usagi"}' -w "\n"
		},
	}
}

func (s *Server) YeeYaHa(c *gin.Context) {
	c.String(http.StatusOK, "Yee~ya~ha!")
}

func (s *Server) ListChiikawaCharacter(c *gin.Context) {
	c.String(http.StatusOK, "Chiikawa Character List : "+strings.Join(chiiikawaList.Characters, ", "))
}

func (s *Server) CreateChiikawaCharacter(c *gin.Context) {
	var character struct {
		Name string `json:"name"`
	}

	err := c.ShouldBindJSON(&character)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chiiikawaList.Characters = append(chiiikawaList.Characters, character.Name)
	c.String(http.StatusOK, "Hello "+character.Name+"!")
}
