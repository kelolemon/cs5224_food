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