create table users(
  id int auto_increment,
  email varchar (320),
  password char (64),
  isBanned int,
  primary key (id)
);

create table userInfo(
  id int,
  nickname varchar (20),
  avatar varchar (200),
  introduction varchar (200),
  primary key (id)
);