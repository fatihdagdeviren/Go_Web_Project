package DBProvider
import (
	"Go_Web_Project/Entities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"log"
)

type Manager interface{
	AddProduct(myProduct *Entities.Product) error
	AddUser(myUser *Entities.User) error
	AddUserProfile(myUserProfile *Entities.UserProfile) error
	Migrate()
	GetAllUserProfiles() []Entities.UserProfile
	GetAllUsers()  (myErr error ,myUserList []Entities.User)
	GetUserByID(id string) Entities.User
	GetUserProfileByUserID(userID string) Entities.UserProfile
}
type manager struct{
	db *gorm.DB
}

var Mgr Manager
var MyDB *gorm.DB

func init(){

	db, err := gorm.Open("mssql", "sqlserver://sa:sifre@localhost?database=HelloGorm&connection+timeout=30")
	if err != nil{
		log.Fatal("Failed to init db:", err)
	}
	MyDB  = db
	Mgr = &manager{db: db}
}

func (mgr *manager) Migrate(){
	mgr.db.AutoMigrate(&Entities.Product{})
	mgr.db.AutoMigrate(&Entities.User{})
	mgr.db.AutoMigrate(&Entities.UserProfile{})
}

func (mgr *manager) AddProduct(myProduct *Entities.Product)(myErr error){
	mgr.db.Create(myProduct)
	if errs := mgr.db.GetErrors(); len(errs)>0 {
		myErr = errs[0]
	}
	return
}

func (mgr *manager) AddUser(myUser *Entities.User) (myErr error){
	mgr.db.Create(myUser)
	if errs := mgr.db.GetErrors(); len(errs)>0 {
		myErr = errs[0]
	}
	return
}

func (mgr *manager) AddUserProfile(myUserProfile *Entities.UserProfile)(myErr error){
	mgr.db.Create(myUserProfile)
	if errs := mgr.db.GetErrors(); len(errs)>0{
		myErr = errs[0]
	}
	return
}

func (mgr *manager) GetAllUserProfiles() []Entities.UserProfile{
	var userProfiles []Entities.UserProfile
	mgr.db.Find(&userProfiles)
	mgr.db.Preload("User").Find(&userProfiles)
	return userProfiles
}

func (mgr *manager) GetAllUsers()  (myErr error ,myUserList []Entities.User) {
	err := mgr.db.Find(&myUserList).Error
	if err != nil{
		return  err , nil
	}
	return
}

func (mgr *manager) GetUserByID(id string) Entities.User {
	var user Entities.User
	mgr.db.Find(&user, id)
	return user
}

func (mgr *manager) GetUserProfileByUserID(userID string) Entities.UserProfile{
	var userProfile Entities.UserProfile
	mgr.db.Where("user_id = ?", userID).Find(&userProfile)
	return userProfile

}
