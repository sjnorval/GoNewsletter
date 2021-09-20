package controllers

import (
	"github.com/gin-gonic/gin"
	models "github.com/sjnorval/newsletter/api/models"
	formaterror "github.com/sjnorval/newsletter/api/utils"
	"net/http"
	"strconv"
)

func (server *Server) CreateUser(c *gin.Context) {
	var newUser models.User
	var err error

	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, formaterror.FormatError(err.Error()))
		return
	}

	newUser.Prepare()
	err = newUser.Validate("")

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, formaterror.FormatError(err.Error()))
		return
	}

	userCreated, err := newUser.SaveUser(server.DB)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, formaterror.FormatError(err.Error()))
		return
	}

	//Would be cool to have callable URLs in the header but as I am very green with these libs I will leave for last.
	//c.Header("Location",fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))

	c.IndentedJSON(http.StatusCreated, userCreated)
}

func (server *Server) GetUsers(c *gin.Context) {
	user := models.User{}

	users, err := user.FindAllUsers(server.DB)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, formaterror.FormatError(err.Error()))
		return
	}

	c.IndentedJSON(http.StatusInternalServerError, users)
}


func (server *Server) GetUser(c *gin.Context) {
	id := c.Param("id")


	uid, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, formaterror.FormatError(err.Error()))
		return
	}

	user := models.User{}
	userGotten, err := user.FindUserByID(server.DB, uint32(uid))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, formaterror.FormatError(err.Error()))
		return
	}

	c.IndentedJSON(http.StatusInternalServerError, userGotten)
}

func (server *Server) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var targetUser models.User
	var err error

	if err := c.BindJSON(&targetUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, formaterror.FormatError(err.Error()))
		return
	}

	uid, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, formaterror.FormatError(err.Error()))
		return
	}

	targetUser.Prepare()
	err = targetUser.Validate("update")

	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, formaterror.FormatError(err.Error()))
		return
	}

	updatedUser, err := targetUser.UpdateAUser(server.DB, uint32(uid))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, formaterror.FormatError(err.Error()))
		return
	}

	c.IndentedJSON(http.StatusOK, updatedUser)
}

//func (server *Server) DeleteUser(c *gin.Context) {
//
//	vars := mux.Vars(r)
//
//	user := models.User{}
//
//	uid, err := strconv.ParseUint(vars["id"], 10, 32)
//
//	if err != nil {
//		responses.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//
//	_, err = user.DeleteAUser(server.DB, uint32(uid))
//
//	if err != nil {
//		responses.ERROR(w, http.StatusInternalServerError, err)
//		return
//	}
//
//	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
//	responses.JSON(w, http.StatusNoContent, "")
//}
