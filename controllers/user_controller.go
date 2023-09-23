package controllers

import (
    "net/http"
    "strconv"
    "github.com/labstack/echo/v4"
    "github.com/jinzhu/gorm"
    "../models"
)

// get all users
func GetUsersController(c echo.Context) error {
    var users []models.User
    if err := models.DB.Find(&users).Error; err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success get all users",
        "users":   users,
    })
}

// get user by id
func GetUserController(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    user := models.User{}
    if err := models.DB.First(&user, id).Error; err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "User Not Found")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success user found",
        "user":    user,
    })
}

// create new user
func CreateUserController(c echo.Context) error {
    user := models.User{}
    c.Bind(&user)
    if err := models.DB.Save(&user).Error; err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success create new user",
        "user":    user,
    })
}

// delete user by id
func DeleteUserController(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    user := models.User{}
    if err := models.DB.First(&user, id).Error; err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "User Not Found")
    }
    if err := models.DB.Delete(&user).Error; err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "delete success",
    })
}

// update user by id
func UpdateUserController(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    input := models.User{}
    c.Bind(&input)

    user := models.User{}
    if err := models.DB.First(&user, id).Error; err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "User Not Found")
    }

    user.Name = input.Name
    user.Email = input.Email
    user.Password = input.Password

    if err := models.DB.Save(&user).Error; err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success updated user",
        "user":    user,
    })
}
