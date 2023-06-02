package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrhid6/GoTest/models"
	"github.com/rahul-sinha1908/go-mongoose/mongoose"
)


func AddAccountRoutes(rg *gin.RouterGroup){
	accounts := rg.Group("/accounts")

	accounts.GET("/:accountid", getAccountByID);

	accountGroup := accounts.Group("/:accountid");
	accountGroup.GET("/agents", getAccountAgents)

	AddUserRoutes(accountGroup);
}

func getAccountByID(c *gin.Context) {
    accountId := c.Param("accountid")

	account := models.Accounts{}
	err := mongoose.FindByID(accountId, &account);

	if(err != nil){
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    c.IndentedJSON(http.StatusOK, account)
}

func getAccountAgents(c *gin.Context){
	accountId := c.Param("accountid")

	account := models.Accounts{}
	err := mongoose.FindByID(accountId, &account);

	if(err != nil){
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

	err = account.PopulateAgents();

	if(err != nil){
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		return
	}

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    c.IndentedJSON(http.StatusOK, account.AgentObjects)
}