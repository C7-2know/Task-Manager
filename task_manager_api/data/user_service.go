package data

import (
	"context"
	"errors"
	"task_manager/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// signup
// login
// promote
// getUsers
// getUser
var JwtSecret = []byte("secret")

func Get_user(email string) (models.User, error) {
	user_collection := Client.Database("task_manager").Collection("users")
	filter := bson.D{{"email", email}}
	result := user_collection.FindOne(context.TODO(), filter)
	var user models.User
	err := result.Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil

}

func Get_users() ([]models.User, error) {
	user_collection := Client.Database("task_manager").Collection("users")
	cursor, err := user_collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return []models.User{}, err
	}
	var users []models.User
	for cursor.Next(context.TODO()) {
		var user models.User
		err = cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func Sign_up(user models.User) error {
	// get the email and password from req

	// hash the password
	// create user
	// respond
	user_collection := Client.Database("task_manager").Collection("users")
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPass)
	if users, _ := Get_users(); len(users) == 0 {
		user.Role = "admin"
	}

	_, err = user_collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil

}

func Login(email, password string) (string, error) {
	user, _ := Get_user(email)
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("invalid password")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"role":  user.Role,
			"exp":   time.Now().Add(time.Hour).Unix(),
		})
	jwtToken, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil

}

func Promote(email string) error {
	user_collection := Client.Database("task_manager").Collection("users")
	filter := bson.D{{"email", email}}

	updated := bson.D{{"$set", bson.D{{"role", "admin"}}}}

	_, err := user_collection.UpdateOne(context.TODO(), filter, updated)
	if err != nil {
		return err
	}
	return nil
}
