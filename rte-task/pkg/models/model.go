package models

import "time"

type UserDetails struct {
	UserId      int       `json:"user_id" gorm:"primarykey;autoIncrement"`
	Name        string    `json:"name"  gorm:"column:name;type:varchar(100)"`
	Email       string    `json:"email"  gorm:"unique;type:varchar(100)"`
	Password    string    `json:"password"  gorm:"column:password;type:varchar(255)"`
	PhoneNumber string    `json:"phone_number"  gorm:"column:phone_number;type:varchar(15)"`
	RoleType    string    `json:"role_type"  gorm:"column:role_type;type:varchar(25)"`
	Token       string    `json:"token,omitempty" gorm:"column:token;type:varchar(255)"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type JobCreation struct {
	JobId        int       `json:"job_id"  gorm:"primarykey;autoIncrement"`
	AdminID      int       `json:"admin_id"  gorm:"column:admin_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CompanyName  string    `json:"company_name" gorm:"column:company_name;type:varchar(100)"`
	CompanyEmail string    `json:"company_email"   gorm:"column:company_email;type:varchar(100)"`
	JobTitle     string    `json:"job_title"   gorm:"column:job_title;type:varchar(100)"`
	JobStatus    string    `json:"job_status"   gorm:"column:job_status;type:varchar(100);constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	JobTime      string    `json:"job_time"   gorm:"column:job_time;type:varchar(50);constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Description  string    `json:"description"   gorm:"column:description;type:text"`
	Experience   string    `json:"experience"   gorm:"column:experience;type:varchar(50)"`
	Skills       string    `json:"skills"   gorm:"column:skills;type:varchar(255)"`
	Vacancy      int       `json:"vacancy"   validate:"required" gorm:"column:vacancy;type:int;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Country      string    `json:"country"   gorm:"column:country ;type:varchar(20)"`
	Address      Address   `json:"address"   gorm:"embedded"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Address struct {
	Street  string `json:"street"   gorm:"type:varchar(255)"`
	City    string `json:"city"     gorm:"type:varchar(100)"`
	State   string `json:"state"    gorm:"type:varchar(100)"`
	ZipCode string `json:"zip_code" gorm:"type:varchar(20)"`
}

type UserJobDetails struct {
	UserId     int          `json:"user_id"  gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	JobID      int          `json:"job_id"   gorm:"foreignKey:JobID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Experience int          `json:"experience" gorm:"column:experience;type:int"`
	Skills     string       `json:"skills"   gorm:"column:skills;type:varchar(255)"`
	Language   string       `json:"language" gorm:"column:language;type:varchar(255)"`
	Country    string       `json:"country"  gorm:"column:country;type:varchar(255)"`
	JobRole    string       `json:"job_role" gorm:"column:job_role;type:varchar(255)"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	User       *UserDetails `json:"User,omitempty" gorm:"foreignKey:UserId"`
	Job        *JobCreation `json:"Job,omitempty" gorm:"foreignKey:JobID"`
}

type Response struct {
	Message string      `json:"message,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ApplicantDetail struct {
	UserId      int       `json:"user_id,omitempty"`
	JobID       int       `json:"job_id,omitempty"`
	Experience  int       `json:"experience,omitempty"`
	Skills      string    `json:"skills,omitempty"`
	Language    string    `json:"language,omitempty"`
	Country     string    `json:"country,omitempty"`
	JobRole     string    `json:"job_role,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Name        string    `json:"name,omitempty"`
	Email       string    `json:"email,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
}

type ErrorResponse struct {
	Error      error
	StatusCode int
}
