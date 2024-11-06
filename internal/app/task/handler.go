package task

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yuuki798/TimerMe3/core/database"
	"log"
	"time"
)

var db = database.GetDb("MainMysql")

func GetTasks(c *gin.Context) {
	var tasks []Task
	// 返回切片
	db.Find(&tasks)
	c.JSON(200, tasks)
	//fmt.Print(tasks)
}

func CreateTask(c *gin.Context) {
	var task Task
	// json->struct
	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		log.Printf("Error: %s", err.Error())
		return
	}
	log.Printf("Task: %v", task)
	task.StartTime = time.Now()
	// struct->db
	db.Create(&task)
	// 201 Created
	c.JSON(201, task)
}

func UpdateTask(c *gin.Context) {
	var task Task
	err := db.Where("id=?", c.Param("id")).First(&task)
	if err.Error != nil {
		c.JSON(404, gin.H{"error": "record not found"})
		return
	}
	db.Save(&task)
	c.JSON(200, task)
}

func DeleteTask(c *gin.Context) {
	var task Task
	err := db.Where("id=?", c.Param("id")).First(&task)
	if err.Error != nil {
		c.JSON(404, gin.H{"error": "record not found"})
		return
	}
	db.Delete(&task)
	c.JSON(204, nil)
}

func StartTask(c *gin.Context) {
	var task Task
	taskID := c.Param("id")
	fmt.Printf("Task ID: %s\n", taskID)
	err := db.Where("id = ?", taskID).First(&task).Error
	if err != nil {
		fmt.Printf("Error finding task: %s\n", err.Error())
		c.JSON(404, gin.H{
			"error": "record not found",
		})
		return
	}
	if task.Status == "started" {
		c.JSON(400, gin.H{
			"error": "Task already started",
		})
		return
	}
	task.Status = "started"
	task.StartTime = time.Now()
	db.Save(&task)
	c.JSON(200, task)
}

func PauseTask(c *gin.Context) {
	var task Task
	err := db.Where("id=?", c.Param("id")).First(&task).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "record not found",
		})
		return
	}
	if task.Status != "started" {
		c.JSON(400, gin.H{
			"error": "Task is not started",
		})
		return
	}
	elapsedTime := int(time.Since(task.StartTime).Seconds())
	task.Duration += elapsedTime
	task.Status = "paused"
	db.Save(&task)
	c.JSON(200, task)
}

func CompleteTask(c *gin.Context) {
	var task Task
	err := db.Where("id=?", c.Param("id")).First(&task).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "record not found",
		})
		return
	}
	task.Status = "completed"
	db.Save(&task)
	c.JSON(200, task)
}

func ResetTask(c *gin.Context) { // 新增reset_task函数
	var task Task
	err := db.Where("id=?", c.Param("id")).First(&task).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "record not found",
		})
		return
	}
	task.Duration = 0
	task.Status = "pending"
	db.Save(&task)
	c.JSON(200, task)
}
