package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
	"net/http"
)

// GetGenres...
func (h *Handler) GetAllMonthlyPlans(c *gin.Context) {
	r := models.NewMonthlyPlanRepository()
	monthlyPlans := r.GetAll()

	c.HTML(http.StatusOK, "monthly_plans.html", gin.H{
		"monthlyPlans": monthlyPlans,
	})
}

//// AddGenres...
//func (h *Handler) AddGenre(c *gin.Context) {
//	r := models.NewGenreRepository()
//	r.GenreName, _ = c.GetPostForm("genre_name")
//	r.Add(&r)
//
//	c.Redirect(http.StatusMovedPermanently, "/_genres")
//}
//
//// GetGenres...
//func (h *Handler) GetGenre(c *gin.Context) {
//	r := models.NewGenreRepository()
//	id, _ := strconv.Atoi(c.Param("id"))
//	genre := r.GetOne(id)
//	c.HTML(http.StatusOK, "genre_edit.html", gin.H{
//		"genre": genre,
//	})
//}
//
//// EditGenres...
//func (h *Handler) EditGenre(c *gin.Context) {
//	r := models.NewGenreRepository()
//	id, _ := strconv.Atoi(c.Param("id"))
//	genre := r.GetOne(id)
//	genre.GenreName, _ = c.GetPostForm("genreName")
//	r.Edit(genre)
//	c.Redirect(http.StatusMovedPermanently, "/_genres")
//}
//
//// DeleteGenres...
//func (h *Handler) DeleteGenre(c *gin.Context) {
//	r := models.NewGenreRepository()
//	id, _ := strconv.Atoi(c.Param("id"))
//	r.Delete(id)
//	c.Redirect(http.StatusMovedPermanently, "/_genres")
//}