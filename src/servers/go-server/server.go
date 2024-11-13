package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "net/http"
    "strconv"
)

type User struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    HoursWorked int    `json:"hoursWorked"`
}

var users = []User{}
var nextID = 1

func main() {
    router := gin.Default()

    // CORS middleware
    router.Use(cors.New(cors.Config{
        AllowAllOrigins: true,
        AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:    []string{"Origin", "Content-Type"},
    }))

    // Routes
    router.GET("/users", func(c *gin.Context) {
        c.JSON(http.StatusOK, users)
    })

    // GET user by ID
    router.GET("/users/:id", func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        for _, user := range users {
            if user.ID == id {
                c.JSON(http.StatusOK, user)
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
    })

    // POST a new user
    router.POST("/users", func(c *gin.Context) {
        var newUser User
        if err := c.BindJSON(&newUser); err != nil || newUser.Name == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid name"})
            return
        }
        newUser.ID = nextID
        nextID++
        newUser.HoursWorked = 0
        users = append(users, newUser)
        c.JSON(http.StatusCreated, newUser)
    })

    // PATCH update hoursWorked for a user
    router.PATCH("/users/:id", func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        var update struct {
            HoursToAdd int `json:"hoursToAdd"`
        }
        if err := c.BindJSON(&update); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }
        for i := range users {
            if users[i].ID == id {
                users[i].HoursWorked += update.HoursToAdd
                c.JSON(http.StatusOK, users[i])
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
    })

    // PUT update user by ID
    router.PUT("/users/:id", func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        var updateData struct {
            Name        string `json:"name"`
            HoursWorked int    `json:"hoursWorked"`
        }
        if err := c.BindJSON(&updateData); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }
        for i := range users {
            if users[i].ID == id {
                if updateData.Name != "" {
                    users[i].Name = updateData.Name
                }
                users[i].HoursWorked = updateData.HoursWorked
                c.JSON(http.StatusOK, users[i])
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
    })

    // DELETE all users
    router.DELETE("/users", func(c *gin.Context) {
        users = []User{}
        nextID = 1
        c.JSON(http.StatusOK, users)
    })

    // DELETE user by ID
    router.DELETE("/users/:id", func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        for i, user := range users {
            if user.ID == id {
                users = append(users[:i], users[i+1:]...)
                c.JSON(http.StatusOK, user)
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
    })

    // Start the server
    router.Run(":5004")
}