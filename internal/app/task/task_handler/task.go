package task_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuki798/TimerMe3/core/ginx/dbx"
	"github.com/yuuki798/TimerMe3/core/libx"
	"github.com/yuuki798/TimerMe3/internal/app/task/task_dto"
	"github.com/yuuki798/TimerMe3/internal/app/task/task_entity"
	"net/http"
	"strconv"
	"time"
)

func GetTasks(c *gin.Context) {
	var tasks []task_entity.Task
	uid := libx.Uid(c)
	// 返回切片
	err := dbx.DB.Where("uid=?", uid).Find(&tasks).Error
	if err != nil {
		libx.Err(c, 500, "获取任务表失败", err)
		return
	}
	libx.Ok(c, "获取任务表成功", gin.H{
		"tasks": tasks,
	})
}

func CreateTask(c *gin.Context) {
	var req task_dto.AddTaskReq
	var task task_entity.Task

	uid := libx.Uid(c)
	// json->struct
	err := c.ShouldBindJSON(&req)
	if err != nil {
		libx.Err(c, 400, "json req 绑定错误", err)
		return
	}
	task.Name = req.Name
	task.Duration = req.Duration
	task.TotalTime = req.TotalTime
	task.StartTime = time.Now()
	task.Status = "paused"
	task.Uid = uid

	// struct->dbx.DB
	err = dbx.DB.Create(&task).Error
	if err != nil {
		libx.Err(c, 500, "创建任务失败", err)
	}
	libx.Ok(c, "创建"+task.Name+"任务成功", gin.H{
		"task": task,
	})
}

func UpdateTask(c *gin.Context) {
	var task task_entity.Task
	var req task_dto.UpdateTaskReq

	uid := libx.Uid(c)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		libx.Err(c, http.StatusBadRequest, "json req 绑定错误", err)
		return
	}
	if req.Recover {
		err = dbx.DB.Unscoped().Where("id=? and uid=?", c.Param("id"), uid).First(&task).Error
		if err != nil {
			libx.Err(c, http.StatusNotFound, "record not found", err)
			return
		}
	} else {
		err = dbx.DB.Where("id=? and uid=?", c.Param("id"), uid).First(&task).Error
		if err != nil {
			libx.Err(c, http.StatusNotFound, "record not found", err)
			return
		}
	}

	if req.Name != "" {
		task.Name = req.Name
	}
	if req.Duration != 0 {
		task.Duration = req.Duration
	}
	if req.TotalTime != 0 {
		task.TotalTime = req.TotalTime
	}
	if req.Recover {
		dbx.DB.Unscoped().Where("uid=?", uid).Model(&task).Update("deleted_at", nil)
	}
	dbx.DB.Save(&task)
	libx.Ok(c, "更新任务成功", gin.H{
		"task": task,
	})
}

func DeleteTask(c *gin.Context) {
	var task task_entity.Task

	uid := libx.Uid(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		libx.Err(c, http.StatusBadRequest, "id必须为数字", err)
		return
	}
	if id == 0 {
		libx.Err(c, http.StatusBadRequest, "id不能为空", nil)
		return
	}
	err = dbx.DB.Where("id=? and uid=?", id, uid).Find(&task).Error
	if err != nil {
		libx.Err(c, http.StatusNotFound, "未找到记录", err)
		return
	}

	dbx.DB.Delete(&task)

	libx.Ok(c, "删除任务成功", nil)
}

func StartTask(c *gin.Context) {
	var task task_entity.Task

	uid := libx.Uid(c)
	taskID := c.Param("id")
	err := dbx.DB.Where("id = ? and uid=?", taskID, uid).First(&task).Error
	if err != nil {
		libx.Err(c, 404, "任务未找到", err)
		return
	}
	if task.Status == "running" {
		libx.Err(c, 400, "任务已经开始！", nil)
		return
	}
	task.Status = "running"
	task.StartTime = time.Now()
	dbx.DB.Save(&task)
	libx.Ok(c, "任务开始", gin.H{
		"task": task,
	})
}

func PauseTask(c *gin.Context) {
	var task task_entity.Task

	uid := libx.Uid(c)
	err := dbx.DB.Where("id=? and uid=?", c.Param("id"), uid).First(&task).Error
	if err != nil {
		libx.Err(c, 404, "未找到任务", err)
		return
	}
	if task.Status != "running" {
		return
	}
	elapsedTime := int(time.Since(task.StartTime).Seconds())
	task.Duration += elapsedTime
	task.Status = "paused"
	dbx.DB.Save(&task)
	libx.Ok(c, "任务暂停", gin.H{
		"task": task,
	})
}

func CompleteTask(c *gin.Context) {
	var task task_entity.Task

	uid := libx.Uid(c)
	err := dbx.DB.Where("id=? and uid=?", c.Param("id"), uid).First(&task).Error
	if err != nil {
		libx.Err(c, 404, "record not found", err)
		return
	}
	task.Status = "completed"
	dbx.DB.Save(&task)
	libx.Ok(c, "Task completed", gin.H{
		"task": task,
	})
}

func ResetTask(c *gin.Context) { // 新增reset_task函数
	var task task_entity.Task

	uid := libx.Uid(c)
	err := dbx.DB.Where("id=? and uid=?", c.Param("id"), uid).First(&task).Error
	if err != nil {
		libx.Ok(c, "record not found", nil)
		return
	}
	task.Duration = 0
	task.Status = "paused"
	dbx.DB.Save(&task)
	libx.Ok(c, "Task reset", gin.H{
		"task": task,
	})
}
