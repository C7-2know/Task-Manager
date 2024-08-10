package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context){
	user,err:=data.Get_users()
	if err!=nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,user)
}

func GetUser(c *gin.Context){
	email:=c.Param("email")
	if email==""{
		c.JSON(http.StatusBadRequest,gin.H{"error":"id is required"})
		return
	}
	user,err:=data.Get_user(email)
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"message":err.Error()})
		return
	}
	c.JSON(http.StatusOK,user)
}

func SignUp(c *gin.Context){
	var new_user models.User
	if err:=c.ShouldBindJSON(&new_user); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	_,err_:=data.Get_user(new_user.Email)
	if err_==nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"user already exists"})
		return
	}
	err:=data.Sign_up(new_user)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusCreated,gin.H{"message":"user created"})
}

func LogIn(c *gin.Context){
	var user models.User
	if err:=c.ShouldBindJSON(&user);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	token,err:=data.Login(user.Email,user.Password)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "Bearer "+token,3600,"","",false,true)
	c.JSON(http.StatusOK,gin.H{"token":token})
}

func PromoteUser(c *gin.Context){
	email:=c.Param("email")
	if email==""{
		c.JSON(http.StatusBadRequest,gin.H{"error":"email is required"})
		return
	}	
	err:=data.Promote(email)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"user promoted"})
}


// Task controllers


func GetTasks(c *gin.Context) {
	response := data.Get_tasks()
	c.JSON(http.StatusOK, response)

}

func GetTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	response, err := data.Get_task(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data.Create_task(newTask)
	c.JSON(http.StatusCreated, gin.H{"message": "task created"})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	var updated models.Task
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data.Update_task(id, updated)
	c.JSON(http.StatusOK, gin.H{"message": "task updated"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	_, err := data.Get_task(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	data.Delete_task(id)
	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
