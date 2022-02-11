package respository

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Kontribute/kontribute-web-backend/entity"
	"golang.org/x/crypto/bcrypt"
	gomail "gopkg.in/mail.v2"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type UserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	ProfileUser(userID string) entity.User
	InsertWebUser(user entity.User) entity.User
	VerifyOTP(email string, otp string) entity.User
	SendOTP(email string) error
}

type userConnection struct {
	connection *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

// type UserRepository interface {
// 	InsertUser(user entity.User) (entity.User, error)
// 	UpdateUser(user entity.User) (entity.User, error)
// 	FindByEmail(email string) (entity.User, error)
// 	FindByUserID(userID string) (entity.User, error)
// 	VerifyCredential(email string, password string) interface{}
// 	IsDuplicateEmail(email string) (tx *gorm.DB)
// 	ProfileUser(userID string) entity.User
// }

// func NewUserRepository(db *gorm.DB) UserRepository {
// 	return &userConnection{
// 		connection: db,
// 	}
// }

func (db *userConnection) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) InsertWebUser(user entity.User) entity.User {
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) entity.User {
	var user entity.User
	db.connection.Preload("Books").Preload("Books.User").Find(&user, userID)
	return user
}

func (db *userConnection) SendOTP(email string) error {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	rand.Seed(time.Now().UnixNano())

	rn := rand.Intn(100000)
	converted_rn := strconv.Itoa(rn)
	user.Verified_Otp = converted_rn

	db.connection.Save(&user)
	if err := GenerateAndMailOTP(email, converted_rn); err != nil {
		return err
	} else {
		return nil
	}
}

func (db *userConnection) VerifyOTP(email string, otp string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)

	//err := errors.New("invalid code, please check your email and try again")
	if user.Verified_Otp == otp {
		user.Verified = 1
		db.connection.Save(&user)
		return user
	} else {
		user = entity.User{}
		return user
	}
}

//it will take in an otp
func GenerateAndMailOTP(email string, rn string) error {
	//rand.Seed(time.Now().UnixNano())
	//rn := rand.Intn(10000)
	fmt.Println("random number:", rn)
	password := "%VGl%G6]pOK9"
	abc := gomail.NewMessage()
	message := fmt.Sprintf("Your Authentication Code Is %v", rn)

	abc.SetHeader("From", "info@mykontribute.com")
	abc.SetHeader("To", email)
	abc.SetHeader("Subject", "This Is Your Authentication info")
	abc.SetBody("text/plain", message)

	a := gomail.NewDialer("mail.mykontribute.com", 465, "info@mykontribute.com", password)
	// a.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := a.DialAndSend(abc); err != nil {
		log.Println(err)
		panic(err)
	}

	return nil
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

// type UserRepository interface {
// 	InsertUser(user entity.User) (entity.User, error)
// 	UpdateUser(user entity.User) (entity.User, error)
// 	FindByEmail(email string) (entity.User, error)
// 	FindByUserID(userID string) (entity.User, error)
// 	VerifyCredential(email string, password string) interface{}
// 	IsDuplicateEmail(email string) (tx *gorm.DB)
// 	ProfileUser(userID string) entity.User
// }

// type userConnection struct {
// 	connection *gorm.DB
// }

//newuserrespository creates a new instance of UserRepository

// func NewUserRepository(db *gorm.DB) UserRepository {
// 	return &userConnection{
// 		connection: db,
// 	}
// }

// func (db *userConnection) InsertUser(user entity.User) entity.User {
// 	user.Password = hashAndSalt([]byte(user.Password))
// 	db.connection.Save(&user)
// 	return user
// }

// func (c *userConnection) UpdateUser(user entity.User) (entity.User, error) {
// 	if user.Password != "" {
// 		user.Password = hashAndSalt([]byte(user.Password))
// 	} else {
// 		var tempUser entity.User
// 		c.connection.Find(&tempUser, user.ID)
// 		user.Password = tempUser.Password
// 	}

// 	c.connection.Save(&user)
// 	return user, nil
// }

// func (db *userConnection) VerifyCredential(email string, password string) interface{} {
// 	var user entity.User
// 	res := db.connection.Where("email = ?", email).Take(&user)
// 	if res.Error == nil {
// 		return user
// 	}
// 	return nil
// }

// func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
// 	var user entity.User
// 	return db.connection.Where("email = ?", email).Take(&user)
// }

// func (db *userConnection) ProfileUser(userID string) entity.User {
// 	var user entity.User
// 	db.connection.Preload("Books").Preload("Books.User").Find(&user, userID)
// 	return user
// }

// func (db *userConnection) FindByEmail(email string) entity.User {
// 	var user entity.User
// 	db.connection.Where("email = ?", email).Take(&user)
// 	return user
// }

// func hashAndSalt(pwd []byte) string {
// 	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
// 	if err != nil {
// 		log.Println(err)
// 		panic("failed to hash passworde")
// 	}
// 	return string(hash)
// }
