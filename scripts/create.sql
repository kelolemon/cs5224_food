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