package handler

import (
	"demo09/internal/model"
	"demo09/internal/service"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskSercice service.TaskService
}

type taskResponse struct {
	Id         string    `json:"id"`
	Title      string    `json:"title"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type taskCreateRequest struct {
	Title string `json:"title"`
}

type taskUpdateRequest struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type resultsResponse struct {
	Results any `json:"results"`
}

func NewTaskHandler(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskSercice: taskService,
	}
}

func (h *TaskHandler) SetupRoutes(router *gin.Engine) {

	gTasks := router.Group("/tareas")
	{
		gTasks.GET("", h.GetAll)
		gTasks.POST("", h.Create)
		gTasks.GET("/:id", h.GetById)
		gTasks.PUT("/:id", h.Update)
		gTasks.DELETE("/:id", h.Delete)
	}
}

func (h *TaskHandler) GetAll(c *gin.Context) {
	tasks, err := h.taskSercice.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: "internal server error"})
		log.Fatalf("internal server error: %v", err)
		return
	}

	response := make([]taskResponse, 0, len(tasks))
	for _, task := range tasks {
		response = append(response, taskResponse{
			Id:         task.Id,
			Title:      task.Title,
			CreatedAt:  task.CreatedAt,
			ModifiedAt: task.ModifiedAt,
		})
	}

	rp := resultsResponse{
		Results: response,
	}

	c.JSON(http.StatusOK, rp)
}

func (h *TaskHandler) Create(c *gin.Context) {
	var rTask taskCreateRequest
	if err := c.ShouldBindJSON(&rTask); err != nil {
		rp := resultsResponse{Results: errorResponse{Error: "invalid request body"}}
		c.JSON(http.StatusBadRequest, rp)
		return
	}

	task := model.Task{
		Title: rTask.Title,
	}

	nTask, err := h.taskSercice.Create(c.Request.Context(), task)
	if err != nil {
		rp := resultsResponse{Results: errorResponse{Error: "error al crear tarea"}}
		c.JSON(http.StatusBadRequest, rp)
		return
	}

	rp := resultsResponse{Results: nTask}
	c.JSON(http.StatusCreated, rp)
}

func (h *TaskHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	task, err := h.taskSercice.GetById(c.Request.Context(), id)
	if err != nil {
		if task.Id == "" {
			c.JSON(http.StatusNotFound, errorResponse{Error: "task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse{Error: "internal server error"})
		log.Fatalf("internal server error: %v", err)
		return
	}

	rp := resultsResponse{
		Results: task,
	}

	c.JSON(http.StatusOK, rp)
}

func (h *TaskHandler) Update(c *gin.Context) {
	var rTask taskUpdateRequest
	if err := c.ShouldBindJSON(&rTask); err != nil {
		rp := resultsResponse{Results: errorResponse{Error: "invalid request body"}}
		c.JSON(http.StatusBadRequest, rp)
		return
	}

	task := model.Task{
		Id:    c.Param("id"),
		Title: rTask.Title,
	}

	nTask, err := h.taskSercice.Update(c.Request.Context(), task)
	if err != nil {
		rp := resultsResponse{Results: errorResponse{Error: "invalid request body"}}
		c.JSON(http.StatusBadRequest, rp)
		return
	}

	rp := resultsResponse{Results: nTask}
	c.JSON(http.StatusCreated, rp)
}

func (h *TaskHandler) Delete(c *gin.Context) {
	nTask, err := h.taskSercice.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		rp := resultsResponse{Results: errorResponse{Error: "invalid request body"}}
		c.JSON(http.StatusBadRequest, rp)
		return
	}

	rp := resultsResponse{Results: nTask}
	c.JSON(http.StatusCreated, rp)
}
