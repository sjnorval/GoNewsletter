package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/sjnorval/newsletter/api/models"
	"github.com/sjnorval/newsletter/api/responses"
	formaterror "github.com/sjnorval/newsletter/api/utils"
)

func (server *Server) CreateTopic(c *gin.Context) {
	var newTopic models.Topic
	var err error

	if err := c.BindJSON(&newTopic); err != nil {
		c.IndentedJSON(http.StatusBadRequest, formaterror.FormatError(err.Error()))
		return
	}

	newTopic.Prepare()
	err = newTopic.Validate("")

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, formaterror.FormatError(err.Error()))
		return
	}

	topicCreated, err := newTopic.SaveTopic(server.DB)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, formaterror.FormatError(err.Error()))
		return
	}

	c.IndentedJSON(http.StatusCreated, topicCreated)
}

func (server *Server) GetTopic(c *gin.Context) {
	topic := models.Topic{}

	topics, err := topic.FindAllTopics(server.DB)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, formaterror.FormatError(err.Error()))
		return
	}

	c.IndentedJSON(http.StatusOK, topics)
}

func (server *Server) GetTopics(c *gin.Context) {
	topic := models.Topic{}
	var err error

	topics, err := topic.FindAllTopics(server.DB)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, formaterror.FormatError(err.Error()))
		return
	}

	c.IndentedJSON(http.StatusOK, topics)
}

func (server *Server) RegisterUserToTopic(c *gin.Context) {

	topicid := c.Param("topicid")
	userid := c.Param("userid")

	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("update")

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedUser, err := user.UpdateAUser(server.DB, uint32(uid))

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, updatedUser)
}


