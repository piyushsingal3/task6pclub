package handlers

import (
	"attendance-app/helper"
	"attendance-app/models"
	"attendance-app/store"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

// hashes the password
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)

	}

	return string(bytes)
}
func SignUp(c *gin.Context, m *store.MongoStore) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.Users

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	count, err := m.UsersCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	defer cancel()
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
		return
	}
	count1, err := m.UsersCollection.CountDocuments(ctx, bson.M{"rollno": user.RollNo})
	defer cancel()
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
		return
	}

	password := HashPassword(*user.Password)
	user.Password = &password

	if count > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "this email already exists"})
		return

	}

	if count1 > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "this rollno already exists"})
		return

	}
	if user.Image != nil {
		_, err := base64.StdEncoding.DecodeString(*user.Image)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image data"})
			return
		}
	}
	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()

	token, refreshToken, _ := helper.GenerateAllTokens(user.Email, user.RollNo)
	user.Token = &token
	user.Refresh_token = &refreshToken

	resultInsertionNumber, insertErr := m.UsersCollection.InsertOne(ctx, user)
	if insertErr != nil {
		msg := fmt.Sprintf("User item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, resultInsertionNumber)

}

// func SignUpAdmin(c *gin.Context, m *store.MongoStore) {
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	var admin models.Admin

// 	if err := c.BindJSON(&admin); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	validationErr := validate.Struct(admin)
// 	if validationErr != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
// 		return
// 	}

// 	count, err := m.AdminCollection.CountDocuments(ctx, bson.M{"email": admin.Email})
// 	defer cancel()
// 	if err != nil {
// 		log.Panic(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
// 		return
// 	}

// 	password := HashPassword(*admin.Password)
// 	admin.Password = &password

// 	if count > 0 {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "this email already exists"})
// 		return

// 	}

// 	admin.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 	admin.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 	admin.ID = primitive.NewObjectID()

// 	token, refreshToken, _ := helper.GenerateAllTokens(admin.Email, admin.RollNo)
// 	admin.Token = &token
// 	admin.Refresh_token = &refreshToken

// 	resultInsertionNumber, insertErr := m.AdminCollection.InsertOne(ctx, admin)
// 	if insertErr != nil {
// 		msg := fmt.Sprintf("Admin item was not created")
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 		return
// 	}
// 	defer cancel()

// 	c.JSON(http.StatusOK, resultInsertionNumber)

// }
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("login or passowrd is incorrect")
		check = false
	}

	return check, msg
}
func Login(c *gin.Context, m *store.MongoStore) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.Users
	var foundUser models.Users

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := m.UsersCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "login or passowrd is incorrect"})
		return
	}

	passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
	defer cancel()
	if passwordIsValid != true {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	token, refreshToken, _ := helper.GenerateAllTokens(foundUser.Email, foundUser.RollNo)

	helper.UpdateAllTokens(token, refreshToken, foundUser.RollNo, m)
	err = m.UsersCollection.FindOne(ctx, bson.M{"email": foundUser.Email}).Decode(&foundUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, foundUser)

}
func LoginAdmin(c *gin.Context, m *store.MongoStore) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var admin models.Admin
	var foundAdmin models.Admin

	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := m.AdminCollection.FindOne(ctx, bson.M{"email": admin.Email}).Decode(&foundAdmin)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password is incorrect"})
		return
	}

	passwordIsValid, msg := VerifyPassword(*admin.Password, *foundAdmin.Password)
	defer cancel()
	if passwordIsValid != true {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	token, refreshToken, _ := helper.GenerateAdminTokens(foundAdmin.Email)

	helper.UpdateAdminTokens(token, refreshToken, foundAdmin.Email, m)
	err = m.AdminCollection.FindOne(ctx, bson.M{"email": foundAdmin.Email}).Decode(&foundAdmin)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, foundAdmin)

}
func Auth(c *gin.Context, m *store.MongoStore) {

	c.JSON(http.StatusOK, "Done")

}
func InsertAttendance(c *gin.Context, m *store.MongoStore) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var attendance models.Attendance
	if err := c.BindJSON(&attendance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(attendance)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	log.Printf("Checking for existing attendance: email=%s, date=%s", attendance.UserEmail, attendance.Date)

	count, err := m.AttendanceCollection.CountDocuments(ctx, bson.M{
		"$and": bson.A{
			bson.M{"useremail": attendance.UserEmail},
			bson.M{"date": attendance.Date},
		},
	})
	log.Print(count)
	defer cancel()
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
		return
	}

	if count > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Attendance already marked"})
		return

	}

	resultInsertionNumber, insertErr := m.AttendanceCollection.InsertOne(ctx, attendance)
	if insertErr != nil {
		msg := fmt.Sprintf("Attendance not marked")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, resultInsertionNumber)

}
func GetUsersAttendance(c *gin.Context, m *store.MongoStore) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var userattendance models.Attendance
	if err := c.BindJSON(&userattendance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var foundUserAttendances []models.Attendance
	cursor, err := m.AttendanceCollection.Find(ctx, bson.M{"useremail": userattendance.UserEmail})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error finding documents: " + err.Error()})
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(context.Background()) {
		var foundUserAttendance models.Attendance
		if err := cursor.Decode(&foundUserAttendance); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error finding documents: " + err.Error()})
			return
		}
		foundUserAttendances = append(foundUserAttendances, foundUserAttendance)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error finding documents: " + err.Error()})
		return
	}
	if err := cursor.All(ctx, &foundUserAttendances); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error decoding results: " + err.Error()})
		return
	}

	if len(foundUserAttendances) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no documents found"})
		return
	}

	// Return the blog entries if found
	c.JSON(http.StatusOK, foundUserAttendances)
}
func GetUsers(c *gin.Context, m *store.MongoStore) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var foundUsers []models.Users
	cursor, err := m.UsersCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error finding documents: " + err.Error()})
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(context.Background()) {
		var foundUser models.Users
		if err := cursor.Decode(&foundUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error finding documents: " + err.Error()})
			return
		}
		foundUsers = append(foundUsers, foundUser)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error finding documents: " + err.Error()})
		return
	}
	if err := cursor.All(ctx, &foundUsers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error decoding results: " + err.Error()})
		return
	}

	if len(foundUsers) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no documents found"})
		return
	}

	// Return the blog entries if found
	c.JSON(http.StatusOK, foundUsers)
}
