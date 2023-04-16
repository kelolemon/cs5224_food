# cs5224_food

## build
```bash
go mod tidy 
go get -u
go build main.go
./main
```

## api 
Using restful API
```go
	api.POST("/user", method.CreateNewUser)
// create user
	api.POST("/login", method.CreateLoginSession)
// login
	api.GET("/user", method.GetUser)
// get user info
	api.GET("/diners", method.GetAllDiners)
// get all diners
	api.POST("/diner", method.CreateDiner)
// create diner
	api.POST("/comment", method.CreateComment)
// create comment / review
	api.GET("/comment", method.GetComments)
// get comment
	api.POST("/dialog", method.CreateDialog)
// create dialog
	api.GET("/dialog", method.GetDialogs)
// get dialog
	api.GET("/ads", method.GetAds)
	api.POST("/ads", method.CreateAds)
// create and get ads
```

## SQL
```sql
CREATE TABLE food.users (
    id varchar(100),
    name varchar(100),
    password varchar(100),
    email varchar(100),
    mobile varchar(100),
    age int,
    gender int,
    primary key(id)
);

create table food.diners (
    id int not null primary key auto_increment,
    name text,
    category text,
    address text,
    url text,
    ratings text,
    reviewers text,
    reviews text,
    review_times text,
    image_url text
);

create table food.comments (
    id int not null  primary key  auto_increment,
    user_id text,
    diner_id int,
    msg text
);

create table food.dialogs (
    id int not null primary key auto_increment,
    father_id int,
    user_id text,
    diner_id int,
    msg text
);

create table food.ads (
    id int not null primary key auto_increment,
    diner_id int,
    vip int,
    end_date timestamp
);
```