CREATE TABLE authentication (
    token varchar(225) primary key not null
);

CREATE UNIQUE INDEX token_authentication ON authentication (token);
