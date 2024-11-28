package models

import (
	"time"
)

type UsersTable struct {
	Name        string    `json:"name" validate:"required" gorm:"column:name;type:varchar(20)"`
	Email       string    `json:"email" validate:"required" gorm:"unique,email;type:varchar(50)"`
	Password    string    `json:"password" validate:"required" gorm:"column:password;type:varchar(255)"`
	PhoneNumber string    `json:"phone_number" validate:"required" gorm:"column:phone_number;type:varchar(255)"`
	RoleType    string    `json:"role_type" gorm:"column:role_type;type:varchar(25)"`
	RoleId      string    `json:"role_id" gorm:"column:role_id;type:varchar(255)"`
	Token       string    `json:"token" gorm:"column:token;type:varchar(255)"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type JobCreation struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	CompanyName  string `json:"company_name" gorm:"column:company_name ;type:varchar(100)"`
	CompanyEmail string `json:"company_email" gorm:"column:company_email ;type:varchar(100)"`
	JobTitle     string `json:"job_title" gorm:"column:job_title ;type:varchar(100)"`
	JobType      string `json:"job_type" gorm:"column:job_type ;type:varchar(100)"`
	JobStatus    string `json:"job_status" gorm:"column:job_status ;type:varchar(10); constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	JobTime      string `json:"job_time" gorm:"column:job_time ;type:varchar(20)"`
	Description  string `json:"description" gorm:"column:description ;type:varchar(255)"`
	Skills       string `json:"skills" gorm:"column:skills ;type:varchar(255)"`
	City         string `json:"city" gorm:"column:city ;type:varchar(100)"`
	State        string `json:"state" gorm:"column:state ;type:varchar(100)"`
	Address      string `json:"address" gorm:"column:address ;type:varchar(255)"`
	Country      string `json:"country" gorm:"column:country ;type:varchar(10)"`
}

type UserJobDetails struct {
	UserId      int `json:"id" gorm:"primaryKey"`
	JobID       int
	JobCreateId JobCreation `json:"job_id" gorm:"foreignKey:JobID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Experience  int         `json:"experience" gorm:"column:experience;type:int"`
	Skills      string      `json:"skills" gorm:"column:skills;type:varchar(255)"`
	Language    string      `json:"language" gorm:"column:language;type:varchar(255)"`
	Country     string      `json:"country" gorm:"column:country;type:varchar(255)"`
	JobRole     string      `json:"job_role" gorm:"column:job_role;type:varchar(255)"`
}
