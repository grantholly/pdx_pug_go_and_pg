CREATE TABLE public.tickets (
       id           serial CONSTRAINT pk PRIMARY KEY,
       subdomain_id int NOT NULL,
       subject      varchar NOT NULL DEFAULT '',
       state        varchar NOT NULL DEFAULT 'open' 
);
