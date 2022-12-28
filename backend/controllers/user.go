package controllers

import (
	"fmt"
	"go-typescript/database"
	"go-typescript/models"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserController struct{}

// ANCHOR - Get All User Data
func (UserController) Index(c *fiber.Ctx) error {

	//ANCHOR - Query Parameters for pagination and search by first name and last name of user data in database
	type Parameters struct {
		Page      int    `json:"page" ` //ANCHOR - Page number
		FirstName string `json:"first_name" query:"first_name"`
		LastName  string `json:"last_name" query:"last_name"`
		Limit     int    //ANCHOR - Number of data per page
	}
	queryParams := Parameters{}
	user := []models.User{}
	//ANCHOR - Get query parameters from request
	if err := c.QueryParser(&queryParams); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err.Error(),
		})
	}
	//ANCHOR - Default value for page number and number of data per page
	if queryParams.Page == 0 {
		queryParams.Page = 1
	}
	if queryParams.Limit == 0 {
		queryParams.Limit = 50
	}
	//ANCHOR - Get all user data from database with pagination and search by first name and last name of user data in database if query parameters are not empty or null or undefined or 0 or false or NaN or "" or " "
	if queryParams.FirstName == "" && queryParams.LastName == "" {
		database.Conn.DB.Preload("File").Limit(queryParams.Limit).Offset((queryParams.Page - 1) * queryParams.Limit).Find(&user)
	} else {
		database.Conn.DB.Preload("File").Limit(queryParams.Limit).Offset((queryParams.Page-1)*queryParams.Limit).Where("UPPER(first_name) LIKE ? AND UPPER(last_name) LIKE ?", "%"+strings.ToUpper(queryParams.FirstName)+"%", "%"+strings.ToUpper(queryParams.LastName)+"%").Find(&user)
	}

	return c.JSON(fiber.Map{
		"count":   len(user),
		"message": "Get All is Successfully",
		"data":    user,
		"page":    queryParams.Page,
	})

}

// ANCHOR - Get Create User Data
func (UserController) Store(c *fiber.Ctx) error {

	//ANCHOR - Get file from request
	file, fileControl := c.FormFile("file")
	//ANCHOR - rollback if there is an error in the process of creating user data
	tx := database.Conn.DB.Begin()
	tx.SavePoint("sp1")
	//ANCHOR - Create user data without file
	if fileControl != nil {
		user := models.User{}
		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		//ANCHOR - Save user data to database if there is no error
		errnotfileData := tx.Preload("File").Create(&user).Error
		if errnotfileData != nil {
			tx.RollbackTo("errorONCreate")
			return c.JSON(errnotfileData)

		}
		//ANCHOR - Commit if there is no error in the process of creating user data
		tx.Commit()
		return c.Status(200).JSON(fiber.Map{
			"message":         "Created is Successfully",
			"employesWarning": user,
		})

	}

	//ANCHOR - Create user data with file
	filePath := strings.Split(file.Filename, ".")
	filePath = strings.Split(filePath[len(filePath)-1], ",")
	//ANCHOR - Check file type
	if filePath[0] != "jpg" && filePath[0] != "png" && filePath[0] != "jpeg" && filePath[0] != "gif" {
		return c.JSON("File Type Not Supported")
	}
	//ANCHOR - Unique file name for each file uploaded
	fileName := time.Now().UnixNano() / int64(time.Millisecond)
	str := strconv.FormatInt(fileName, 16)
	file.Filename = str
	file.Filename = file.Filename + "." + filePath[0]

	//ANCHOR - Check if file already exists
	if fileControl == nil {
		//ANCHOR - Save file to storage folder
		errorSave := c.SaveFile(file, fmt.Sprintf("./storage/users/%s", file.Filename))
		if errorSave != nil {
			tx.RollbackTo("Error Save File On Storage")
			return c.Status(400).JSON("Error Save File")
		}
	}
	//ANCHOR - User data and file data to be saved to database if there is no error in the process of creating user data
	user := models.User{File: models.File{File: file.Filename}}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	//ANCHOR - Save user and file data to database if there is no error
	err1 := tx.Preload("File").Create(&user).Error
	if err1 != nil {
		tx.RollbackTo("sp1")
		return c.JSON("Error Save Data")

	}
	//ANCHOR - Commit if there is no error in the process of creating user data
	tx.Commit()
	return c.Status(201).JSON(fiber.Map{
		"message": "Created is Successfully",
		"user":    user,
	})
}

// ANCHOR - Get User Data By ID
func (UserController) Get(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id")) //ANCHOR - Get id from request
	var user models.User
	database.Conn.DB.Preload("File").First(&user, id) //ANCHOR - Get user data by id from database

	//ANCHOR - Check if user data is not found
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Get User",
		"user":    user,
	})

}

func (UserController) Update(c *fiber.Ctx) error {

	//ANCHOR - Get id from request
	id, _ := strconv.Atoi(c.Params("id"))
	var user models.User
	tx := database.Conn.DB.Begin()
	tx.SavePoint("sp1")
	tx.Preload("Files").Find(&user, id) //ANCHOR - Get user data by id from database and file data
	if user.ID == 0 {
		return c.JSON(fiber.Map{ //ANCHOR - Check if id is not found
			"message": "ID Not Found",
		})
	}
	//ANCHOR - Get file from request
	file, fileControl := c.FormFile("file")
	//ANCHOR - rollback if there is an error in the process of updating user data

	//ANCHOR - Update user data without file
	if fileControl != nil {
		user := models.User{}
		errData := database.Conn.DB.First(&user, id).Error
		if errData != nil {
			tx.RollbackTo("notGetUser")
			return c.JSON(errData)
		}
		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		//ANCHOR - Save user data to database if there is no error
		errnotfileData := tx.Model(&user).Where("id = ?", id).Updates(&user).Error
		if errnotfileData != nil {
			tx.RollbackTo("errorONCreate")
			return c.JSON(errnotfileData)

		}
		//ANCHOR - Commit if there is no error in the process of updating user data
		tx.Commit()
		return c.Status(200).JSON(fiber.Map{
			"message": "Updated is Successfully",
			"user":    &user,
		})

	}

	//ANCHOR - Update user data with file
	filePath := strings.Split(file.Filename, ".")
	filePath = strings.Split(filePath[len(filePath)-1], ",")
	//ANCHOR - Check file type
	if filePath[0] != "jpg" && filePath[0] != "png" && filePath[0] != "jpeg" && filePath[0] != "gif" {
		return c.JSON("File Type Not Supported")
	}
	//ANCHOR - Unique file name for each file uploaded
	fileName := time.Now().UnixNano() / int64(time.Millisecond)
	str := strconv.FormatInt(fileName, 16)
	file.Filename = str
	file.Filename = file.Filename + "." + filePath[0]

	//ANCHOR - Check if file already exists
	if fileControl == nil {
		//ANCHOR - Save file to storage folder
		errorSave := c.SaveFile(file, fmt.Sprintf("./storage/users/%s", file.Filename))
		if errorSave != nil {
			tx.RollbackTo("Error Save File On Storage")
			return c.Status(400).JSON("Error Save File")
		}
	}
	//ANCHOR - User data and file data to be saved to database if there is no error in the process of creating user data
	newFile := models.File{File: file.Filename}
	newUser := models.User{}
	oldUser := models.User{}

	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	//ANCHOR - Old user data from database
	err := tx.Where("id = ?", id).Preload("File").First(&oldUser, id).Error
	if err != nil {
		tx.RollbackTo("sp1")
		return c.JSON("Error Save Data")
	}
	//ANCHOR - Save user data to database if there is no error
	errUp := tx.Model(&oldUser).Where("id = ?", id).Updates(&newUser).Error
	if errUp != nil {
		tx.RollbackTo("sp1")
		return c.JSON("Error Save Data 1")
	}
	//ANCHOR - Check if file already exists
	if oldUser.File.File != "" {
		errDel := os.Remove(fmt.Sprintf("./storage/users/%s", oldUser.File.File))
		if errDel != nil {
			tx.RollbackTo("sp1")
			return c.JSON("File Not Found")
		}
		//ANCHOR - Save file to storage folder if there is no error in the process of deleting the file
		err = tx.Model(&models.File{}).Where("table_id=? and table_type=?", oldUser.File.TableID, "users").Updates(newFile).Error
		if err != nil {
			tx.RollbackTo("sp1")
			return c.JSON("Error Save Data 3")
		}
	} else {
		//ANCHOR - Save file to database if there is no error
		oldUser.File.File = file.Filename
		err = tx.Model(&oldUser).Where("id=?", id).Save(oldUser).Error
		if err != nil {
			tx.RollbackTo("sp1")
			return c.JSON("Error Save Data 4")
		}

	}
	tx.Commit()

	//ANCHOR - Commit if there is no error in the process of updating user data
	return c.Status(200).JSON(fiber.Map{
		"message": "Updated is Successfully",
		"user":    oldUser,
	})
}

func (UserController) Delete(c *fiber.Ctx) error {
	//ANCHOR - Get id from request
	id, _ := strconv.Atoi(c.Params("id"))
	var user models.User
	//ANCHOR - Rollback if there is an error in the process of deleting user data
	txDelete := database.Conn.DB.Begin()
	txDelete.SavePoint("sp1")
	txDelete.Preload("File").First(&user, id)
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}
	//ANCHOR - Delete user data and file data if there is no error in the process of deleting user data
	errData := txDelete.Delete(&user.File).Error
	if errData != nil {
		txDelete.RollbackTo("errorOne")
		return c.Status(400).JSON(fiber.Map{
			"message": "Error Delete File",
		})
	}
	//ANCHOR - Delete file data if file data is not empty
	if user.File.File != "" {
		//ANCHOR - Delete file data if there is no error in the process of deleting user data
		errFileData := os.Remove(fmt.Sprintf("./storage/users/%s", user.File.File))
		if errFileData != nil {
			txDelete.RollbackTo("errorTwo")
			return c.Status(400).JSON(fiber.Map{
				"message": "Error Delete File",
			})
		}
	}
	//ANCHOR - Delete user data if there is no error in the process of deleting user data
	txDelete.Commit()
	return c.Status(200).JSON(fiber.Map{
		"message": "Delete User",
		"user":    user,
	})
}
