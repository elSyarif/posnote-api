CREATE TABLE roles(
    id varchar(36) primary key not null,
    name varchar(36) not null ,
    description text
);

CREATE UNIQUE INDEX name_roles ON roles (name);
