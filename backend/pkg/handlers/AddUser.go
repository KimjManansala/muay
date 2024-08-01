package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) AddUser(c *gin.Context) {
	// Read to request body
	c.JSON(http.StatusOK, gin.H{
		"message": "Ad user",
	})

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// var user models.User
	// json.Unmarshal(body, &user)

	// // Append to the User table
	// if result := h.DB.Create(&user); result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	// // Send a 201 created response
	// w.Header().Add("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode("Created")
}
