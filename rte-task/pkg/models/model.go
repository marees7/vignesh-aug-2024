package models

import "time"

type UserDetails struct {
	UserId      int    `json:"user_id" gorm:"primarykey;autoIncrement"`
	Name        string `json:"name"  gorm:"column:name;type:varchar(100)"`
	Email       string `json:"email"  gorm:"unique;type:varchar(100)"`
	Password    string `json:"password"  gorm:"column:password;type:varchar(255)"`
	PhoneNumber string `json:"phone_number"  gorm:"column:phone_number;type:varchar(15)"`
	RoleType    string `json:"role_type"  gorm:"column:role_type;type:varchar(25)"`
	// RoleId      string    `json:"role_id" gorm:"column:role_id;type:varchar(255)"`
	Token     string    `json:"token,omitempty" gorm:"column:token;type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type JobCreation struct {
	JobId        int       `json:"job_id" gorm:"primarykey;autoIncrement"`
	UserID       int       `json:"user_id"  validate:"required" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CompanyName  string    `json:"company_name"  validate:"required" gorm:"column:company_name;type:varchar(100)"`
	CompanyEmail string    `json:"company_email"  validate:"required" gorm:"column:company_email;type:varchar(100)"`
	JobTitle     string    `json:"job_title"  validate:"required" gorm:"column:job_title;type:varchar(100)"`
	JobStatus    string    `json:"job_status"  validate:"required" gorm:"column:job_status;type:varchar(100);constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	JobTime      string    `json:"job_time"  validate:"required" gorm:"column:job_time;type:varchar(50);constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Description  string    `json:"description"  validate:"required" gorm:"column:description;type:text"`
	Skills       string    `json:"skills"  validate:"required" gorm:"column:skills;type:varchar(255)"`
	Vacancy      int       `json:"vacancy" validate:"required"  gorm:"column:vacancy;type:int;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Country      string    `json:"country"  validate:"required" gorm:"column:country ;type:varchar(20)"`
	Address      Address   `json:"address"  validate:"required" gorm:"embedded"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Address struct {
	Street  string `json:"street"  validate:"required" gorm:"type:varchar(255)"`
	City    string `json:"city"  validate:"required" gorm:"type:varchar(100)"`
	State   string `json:"state"  validate:"required" gorm:"type:varchar(100)"`
	ZipCode string `json:"zip_code"  validate:"required" gorm:"type:varchar(20)"`
}

type UserJobDetails struct {
	ID         int          `json:"id" gorm:"primaryKey"`
	UserID     int          `json:"user_id"  validate:"required" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	JobID      int          `json:"job_id"  validate:"required" gorm:"foreignKey:JobID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Experience int          `json:"experience"  validate:"required" gorm:"column:experience;type:int"`
	Skills     string       `json:"skills"  validate:"required" gorm:"column:skills;type:varchar(255)"`
	Language   string       `json:"language"  validate:"required" gorm:"column:language;type:varchar(255)"`
	Country    string       `json:"country"  validate:"required" gorm:"column:country;type:varchar(255)"`
	JobRole    string       `json:"job_role"  validate:"required" gorm:"column:job_role;type:varchar(255)"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	User       *UserDetails `json:"User" gorm:"foreignKey:UserID"`
	Job        *JobCreation `json:"Job" gorm:"foreignKey:JobID"`
}

type CommonResponse struct {
	Message string      `json:"message,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
