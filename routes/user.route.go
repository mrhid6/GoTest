package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrhid6/GoTest/models"
	"github.com/rahul-sinha1908/go-mongoose/mongoose"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func AddUserRoutes(rg *gin.RouterGroup){
	users := rg.Group("/users")

	users.GET("/", GetAllUsers)
	users.GET("/:userid", GetUserByID)
}

func GetAllUsers(c *gin.Context){
	accountId := c.Param("accountid")

	account := models.Accounts{}
	err := mongoose.FindByID(accountId, &account);

	if(err != nil){
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	err = mongoose.PopulateObjectArray(&account,"Users",&account.UserObjects);

	if(err != nil){
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	for idx, _ := range account.UserObjects {
		(&account.UserObjects[idx]).PopulateUserRole();
	}

	c.IndentedJSON(http.StatusOK, account.UserObjects);
}

func GetUserByID(c *gin.Context){
	accountId := c.Param("accountid")
	userId := c.Param("userid")

	useroid,err := primitive.ObjectIDFromHex(userId);

	if(err != nil){
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	account := models.Accounts{}
	err = mongoose.FindByID(accountId, &account);

	if(err != nil){
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	err = mongoose.PopulateObjectArray(&account,"Users",&account.UserObjects);

	if(err != nil){
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	var theUser models.Users;
	for idx, u := range account.UserObjects {
		(&account.UserObjects[idx]).PopulateUserRole();

		if(u.ID == useroid){
			theUser = u;
		}
	}

	if(!theUser.ID.IsZero()){
		theUser.PopulateUserRole();
		c.IndentedJSON(http.StatusOK, theUser);
		return;
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error":"User Not Found!", "success": false});
}