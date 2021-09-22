# threedee
This is my personal repository to learn microservice in Go without frameworks (go packages like httprouter and cors are still included)

## Helpful Non Vanilla Packages
### General
- cors          (CORS settings)
- httprouter    (Better routing alternative to vanilla net/http's one file per Handler)
- gotenv        (Useful for global variables using .env file)
- logrus        (Good for logging middleware)

### DB
- lib/pq        (Postgresql Driver. This is required by the database/sql for connection)


## Create DB and Populate Initial Data in Postgresql
### A. Create Table
```
CREATE TABLE tbl_m_3d_print_request (
   id bigserial primary key not null,
   item_name varchar(100) not null,
   est_weight float8 not null,
   est_filament_length float8 not null,
   est_duration int not null,
   file_url text not null,
   requestor varchar(100) not null,
   status varchar(20) not null default 'received',
   created_on timestamptz not null default now(),
   created_by varchar(100) not null default 'system',
   modified_on timestamptz null,
   modified_by varchar(100) null,
   is_active bool not null default true
);
```

### B. Populate Data
```
NOTE:
est_weight in gram
est_filament_length in cm
duration in second

INSERT INTO tbl_m_3d_print_request(item_name,est_weight,est_filament_length,est_duration,file_url,requestor) VALUES ('phone holder v2',75,10000,18000,'http://drive.google.com/file/1','Karim');
INSERT INTO tbl_m_3d_print_request(item_name,est_weight,est_filament_length,est_duration,file_url,requestor) VALUES ('cup holder v1',150,20000,36000,'http://drive.google.com/file/2','Kosasih');
INSERT INTO tbl_m_3d_print_request(item_name,est_weight,est_filament_length,est_duration,file_url,requestor) VALUES ('Gantungan baju',50,6666.67,12000,'http://drive.google.com/file/3','Burhan');
```

