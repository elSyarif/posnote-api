CREATE TABLE employee_plants(
    id varchar(36) primary key not null,
    employee_id varchar(36) not null,
    plant_id varchar(36) not null,
    position varchar(50) not null,
    join_date datetime,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
   	updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN key (employee_id) REFERENCES employees(id) ON DELETE CASCADE,
    FOREIGN key (plant_id) REFERENCES plants (id) ON DELETE CASCADE
);

CREATE UNIQUE INDEX employee_id_employee_plants ON employee_plants (employee_id);
CREATE UNIQUE INDEX plant_id_employee_plants ON employee_plants (plant_id);
