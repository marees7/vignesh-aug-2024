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
| POST |	/auth/signup	| Register a new user |
| POST |	/auth/login	| Log in and obtain JWT |


## ADMIN API

| Method | 	Endpoint | 	Description |
| ---- | -------- | -------- |
| POST |	/admin/insert/:admin_id | Create a new posts |
| PUT  |	/admin/userjobsbyid/:job_id/:admin_id	| Update thier own posts | 
| GET  |	/admin/userid/:user_id/:admin_id |Get Jobs By Admin for who applied in their jobposts|
| GET   |   /admin/userdetails/:job_role/:admin_id  | Get by role and Userid|
| GET   |   /admin/postdetails/:admin_id    |Get Jobs for their Own AdminIds|
| GET   |  /admin/userjobsbyid/:job_id/:admin_id   | Get by JobId and UserId|

## USER API

| Method | 	Endpoint | 	Description |
| ---- | -------- | -------- |
| GET  |	/users/allposts	| user or admin Get all job post details |
| GET  |	/users/jobs/:job_title/:country	| Get the jobs by their country |
| GET  |	/users/company/:company_name	| Get all posts in the company |
| POST |	/users/post/:user_id    | Apply JobPost for the Job |
| GET  |	/usersowndetails/user/:user_id| user get by their userowndetails |



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
    token character varying(255) COLLATE pg_catalog."default",
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT user_details_pkey PRIMARY KEY (user_id),
    CONSTRAINT uni_user_details_email UNIQUE (email)
)
CREATE TABLE IF NOT EXISTS public.job_creations
(
    job_id bigint NOT NULL DEFAULT nextval('job_creations_job_id_seq'::regclass),
    domain_id bigint,
    company_name character varying(100) COLLATE pg_catalog."default",
    company_email character varying(100) COLLATE pg_catalog."default",
    job_title character varying(100) COLLATE pg_catalog."default",
    job_status character varying(100) COLLATE pg_catalog."default",
    job_time character varying(50) COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    experience bigint,
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
