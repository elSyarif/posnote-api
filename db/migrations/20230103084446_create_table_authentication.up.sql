CREATE TABLE authentication (
    token text not null
);

CREATE UNIQUE INDEX token_authentication ON authentication (token);
