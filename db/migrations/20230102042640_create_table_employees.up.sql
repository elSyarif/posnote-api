CREATE TABLE employees(
    id varchar(36) primary key not null,
    fullname varchar(75) not null,
    username varchar(35) not null,
    password text not null,
    role_id varchar(36) not null,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP, 
    updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX username_employee ON employees (username);
ALTER TABLE employees 
ADD CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES roles(id)