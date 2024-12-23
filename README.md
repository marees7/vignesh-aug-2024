# vignesh-aug-2024

# Job Site
<!-- PROJECT LOGO -->
<br />
<div align="center">
    <img src="https://i.imgur.com/uQXaHzv.png" alt="Logo" width="150" height="60">

  <h3 align="center">Simple Job Site</h3>
</div>

REST API for JobSite using golang, gin framework & GORM ORM.


## Overview
- There are 2 different roles, User and admin
- A user can apply posts, view their job applied details ,get all job posts (User can see all job updates,view their own details)
- A admin creates the JobPost,update the job posts,view the applicant details(Admin can update their own posts)

### IDE and Tools I Use

<img height="50" width="50" src="https://img.icons8.com/color/48/000000/visual-studio-code-2019.png"/>    <img height="50" width="50" src="https://i.imgur.com/biFSWGU.png"/> <img height="50" width="50" src="https://i.imgur.com/eOGfY8b.png"/> <img height="50" width="50" src="https://i.imgur.com/8T8BE02.png"/> <img height="50" src="https://i.imgur.com/XU4XCrs.png"/> <img height="50" src="https://i.imgur.com/bHV6tyi.png"/> 


## Features
- User authentication and authorization using JSON Web Tokens (JWT)
- Admins can only see their company details and view their company applied users only
- Users can get the all company post and they can able to apply for this jobs
- Error handling and response formatting
- Input validation and data sanitization
- Database integration using PostgreSQL

## Run Locally

Clone the project

```bash
  git clone https://github.com/marees7/vignesh-aug-2024.git
```

Go to the project directory
go to the cmd folder and main.go file.
change the credentials of postgres db in the internals.

```bash
  go run main.go
```


## API Endpoints

The following endpoints are available in the API:

## AUTH API

| Method | 	Endpoint | 	Description |
| ---- | -------- | -------- |
| POST |	v1/auth/signup	| Register a new user |
| POST |	v1/auth/login	| Log in and obtain JWT |


## ADMIN API

| Method | 	Endpoint | 	Description |
| ---- | -------- | -------- |
| POST |	/v1/admin  | Admin creates new JobPost |
| GET  |	/v1/admin/jobs |Get by jobrole and jobID|
| GET   |   /v1/admin/:user_id | Get by Userid|
| GET   |   /v1/admin/   |admin get by his id to know about how many post created|
| PUT  |	/v1/admin/:job_id	| Update thier own posts | 
| DELETE  |	/v1/admin	| Automaticaly delete jobPost | 

## USER API

| Method | 	Endpoint | 	Description |
| ---- | -------- | -------- |

| POST |	/v1/user    | User apply the job in that posts |
| GET  |	/v1/user     |User get by their userowndetails |



## Database Schema

The application uses a PostgreSQL database with the following schema:

```sql
CREATE TABLE IF NOT EXISTS public.user_details
(
    user_id bigint NOT NULL DEFAULT nextval('user_details_user_id_seq'::regclass),
    name character varying(100) COLLATE pg_catalog."default",
    email character varying(100) COLLATE pg_catalog."default",
    password character varying(255) COLLATE pg_catalog."default",
    phone_number character varying(15) COLLATE pg_catalog."default",
    role_type character varying(25) COLLATE pg_catalog."default",
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT user_details_pkey PRIMARY KEY (user_id),
    CONSTRAINT uni_user_details_email UNIQUE (email)
)

CREATE TABLE IF NOT EXISTS public.job_creations
(
    job_id bigint NOT NULL DEFAULT nextval('job_creations_job_id_seq'::regclass),
    admin_id bigint,
    company_name character varying(100) COLLATE pg_catalog."default",
    company_email character varying(100) COLLATE pg_catalog."default",
    job_role character varying(100) COLLATE pg_catalog."default",
    job_status character varying(100) COLLATE pg_catalog."default",
    job_time character varying(50) COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    experience character varying(50) COLLATE pg_catalog."default",
    skills character varying(255) COLLATE pg_catalog."default",
    vacancy bigint,
    "country " character varying(20) COLLATE pg_catalog."default",
    street character varying(255) COLLATE pg_catalog."default",
    city character varying(100) COLLATE pg_catalog."default",
    state character varying(100) COLLATE pg_catalog."default",
    zip_code character varying(20) COLLATE pg_catalog."default",
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT job_creations_pkey PRIMARY KEY (job_id)
)

CREATE TABLE IF NOT EXISTS public.user_job_details
(
    user_id bigint,
    job_id bigint,
    experience bigint,
    skills character varying(255) COLLATE pg_catalog."default",
    language character varying(255) COLLATE pg_catalog."default",
    country character varying(255) COLLATE pg_catalog."default",
    job_role character varying(255) COLLATE pg_catalog."default",
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT fk_user_job_details_job FOREIGN KEY (job_id)
        REFERENCES public.job_creations (job_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_user_job_details_user FOREIGN KEY (user_id)
        REFERENCES public.user_details (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

```

## Register a new user

## First Create User for it

## Request
```
{
    "name": "user 2",
    "email": "d@gmail.com",
    "password": "Abcd123&",
    "phone_number": "2879478999",
    "role_type": "USER"
}

```

## Response
```
{
    "message": "Sucessfully Created the User Detail",
    "data": {
        "user_id": 1,
        "name": "user 2",
        "email": "d@gmail.com",
        "password": "$2a$14$nEMdU73k9F7sCbGRw/d0Uuj6QjfkmExjmKsB0hfbXSJmHHDrU0P02",
        "phone_number": "2879478999",
        "role_type": "USER",
        "created_at": "2024-12-23T11:26:17.839386+05:30",
        "updated_at": "2024-12-23T11:26:17.839386+05:30"
    }
}

```

## Login a User

## Request
```
  {  
     "email": "a@gmail.com",
     "password": "Abcd123&"
  }
```

## Response
```
{
    "message": "Login Sucessfully",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImRAZ21haWwuY29tIiwiTmFtZSI6InVzZXIgMiIsIlJvbGVUeXBlIjoiVVNFUiIsIlVzZXJJRCI6MSwiZXhwIjoxNzM0OTQwODQxfQ.e68l_W7X1v4TMJixDw38_93ylcaEo28PXW4MH8YKgxE",
    "user_id": 2,
    "role_type": "USER"
}
```
## ADMIN API'S
## Creation of Job Post

## Request
```
{
    "admin_id": 1,
    "company_name": "wipro",
    "company_email": "wipro@gmail.com",
    "job_role": "Developer",
    "job_status": "IN PROGRESS",
    "job_time": "FULL TIME",
    "experience":"0-5 years",
    "description": "youn dynamic persons need for this job role",
    "Skills": "Java,Golang,C,C++",
    "Country": "USA",
    "vacancy": 16,
    "address": {
        "street": "abc",
        "city": "chennnai",
        "State": "tamilnadu",
        "zip_code": "123453"
    }
}
```

## Response

```
{
    "message": "Sucessfully Created the JobPost Detail",
    "data": {
        "job_id": 1,
        "admin_id": 2,
        "company_name": "wipro",
        "company_email": "wipro@gmail.com",
        "job_role": "Developer",
        "job_status": "IN PROGRESS",
        "job_time": "FULL TIME",
        "description": "youn dynamic persons need for this job role",
        "experience": "0-5 years",
        "skills": "Java,Golang,C,C++",
        "vacancy": 16,
        "country": "USA",
        "address": {
            "street": "abc",
            "city": "chennnai",
            "state": "tamilnadu",
            "zip_code": "123453"
        },
        "created_at": "2024-12-23T11:34:14.2403358+05:30",
        "updated_at": "2024-12-23T11:34:14.2403358+05:30"
    }
}
```

## Get By Own Posts

## Response

```
{
    "message": "Sucessfully fetched the Jobcreated Details",
    "data": [
        {
            "job_id": 1,
            "admin_id": 2,
            "company_name": "wipro",
            "company_email": "wipro@gmail.com",
            "job_role": "Developer",
            "job_status": "IN PROGRESS",
            "job_time": "FULL TIME",
            "description": "youn dynamic persons need for this job role",
            "experience": "0-5 years",
            "skills": "Java,Golang,C,C++",
            "vacancy": 16,
            "country": "USA",
            "address": {
                "street": "abc",
                "city": "chennnai",
                "state": "tamilnadu",
                "zip_code": "123453"
            },
            "created_at": "2024-12-23T11:34:14.240335+05:30",
            "updated_at": "2024-12-23T11:34:14.240335+05:30"
        }
    ],
    "limit": 10
}
```

## Get by Appplied Job Posts By query param

## Response
```
{
    "message": "Successfully fetched Job Posts ",
    "data": [
        {
            "user_id": 1,
            "job_id": 1,
            "experience": 8,
            "skills": "Golang,Dotnet",
            "language": "English,Tamil",
            "country": "India",
            "job_role": "Developer",
            "updated_at": "0001-01-01T00:00:00Z",
            "name": "user 2",
            "email": "d@gmail.com",
            "phone_number": "2879478999"
        }
    ],
    "total": 1,
    "limit": 1
}
```

## Admin Get Users By Specific ID

## Response
```
{
    "message": "Successfully fetched User applied Jobs",
    "data": [
        {
            "user_id": 1,
            "job_id": 1,
            "experience": 8,
            "skills": "Golang,Dotnet",
            "language": "English,Tamil",
            "country": "India",
            "job_role": "Developer",
            "created_at": "2024-12-23T11:48:07.570183+05:30",
            "updated_at": "2024-12-23T11:48:07.570183+05:30",
            "Job": {
                "job_id": 1,
                "admin_id": 2,
                "company_name": "wipro",
                "company_email": "wipro@gmail.com",
                "job_role": "Developer",
                "job_status": "IN PROGRESS",
                "job_time": "FULL TIME",
                "description": "youn dynamic persons need for this job role",
                "experience": "0-5 years",
                "skills": "Java,Golang,C,C++",
                "vacancy": 16,
                "country": "USA",
                "address": {
                    "street": "abc",
                    "city": "chennnai",
                    "state": "tamilnadu",
                    "zip_code": "123453"
                },
                "created_at": "2024-12-23T11:34:14.240335+05:30",
                "updated_at": "2024-12-23T11:34:14.240335+05:30"
            }
        }
    ],
    "total": 1,
    "limit": 10
}
```
## Delete API'S

## Response

```{
  "Message":Sucessfully Deleted the Job post Details
}
```


## User API'S

## User Can Applied Job Post 
## Request
```
{
    "user_id": 1,
    "job_id":1,
    "experience": 8,
    "skills": "Golang,Dotnet",
    "language": "English,Tamil",
    "country": "India",
    "job_role": "Developer"
}
```

## Response

```
{
    "message": "Sucessfully Applied Job ",
    "data": {
        "user_id": 1,
        "job_id": 1,
        "experience": 8,
        "skills": "Golang,Dotnet",
        "language": "English,Tamil",
        "country": "India",
        "job_role": "Developer",
        "created_at": "2024-12-23T11:48:07.5701835+05:30",
        "updated_at": "2024-12-23T11:48:07.5701835+05:30"
    }
}
```

## User can get all applied Job Posts

## Response
```
{
    "message": "Sucessfully Get the details by thier ID",
    "data": [
        {
            "user_id": 1,
            "job_id": 1,
            "experience": 8,
            "skills": "Golang,Dotnet",
            "language": "English,Tamil",
            "country": "India",
            "job_role": "Developer",
            "created_at": "2024-12-23T11:48:07.570183+05:30",
            "updated_at": "2024-12-23T11:48:07.570183+05:30",
            "Job": {
                "job_id": 1,
                "admin_id": 2,
                "company_name": "wipro",
                "company_email": "wipro@gmail.com",
                "job_role": "Developer",
                "job_status": "IN PROGRESS",
                "job_time": "FULL TIME",
                "description": "youn dynamic persons need for this job role",
                "experience": "0-5 years",
                "skills": "Java,Golang,C,C++",
                "vacancy": 16,
                "country": "USA",
                "address": {
                    "street": "abc",
                    "city": "chennnai",
                    "state": "tamilnadu",
                    "zip_code": "123453"
                },
                "created_at": "2024-12-23T11:34:14.240335+05:30",
                "updated_at": "2024-12-23T11:34:14.240335+05:30"
            }
        }
    ],
    "total": 1,
    "limit": 10
}
```

## User can Get all Post Details

## Response
```
{
    "message": "Hurray! Successfully fetched the JobPosts",
    "data": [
        {
            "job_id": 1,
            "admin_id": 2,
            "company_name": "wipro",
            "company_email": "wipro@gmail.com",
            "job_role": "Developer",
            "job_status": "IN PROGRESS",
            "job_time": "FULL TIME",
            "description": "youn dynamic persons need for this job role",
            "experience": "0-5 years",
            "skills": "Java,Golang,C,C++",
            "vacancy": 16,
            "country": "USA",
            "address": {
                "street": "abc",
                "city": "chennnai",
                "state": "tamilnadu",
                "zip_code": "123453"
            },
            "created_at": "2024-12-23T11:34:14.240335+05:30",
            "updated_at": "2024-12-23T11:34:14.240335+05:30"
        }
    ],
    "total": 1,
    "limit": 10
}
```