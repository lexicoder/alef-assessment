create table if not exists category(
    id serial primary key,
    description varchar(100) not null
);

create table if not exists products(
    id bigserial primary key,
    name varchar(255) not null,
    price real not null,
    quantity integer default 0,
    amount real default 0.0,
    category bigint not null,
    constraint products_category_fk foreign key(category)
    references category(id)
);

create table if not exists users(
    id bigserial primary key,
    firstname varchar(15) not null,
    lastname varchar(20) not null,
    email varchar(40) not null unique,
    password varchar(100) not null,
    status char(1) default '0'
);

insert into category(description) values ('Costumes');
insert into category(description) values ('Media');
insert into category(description) values ('Accessories');
insert into category(description) values ('Bags');

insert into products(name, price, quantity, amount, category) values
('Minnie Mouse Rose Gold Costume for Adults', 16.80, 100, (16.80 * 100), 1);
insert into products(name, price, quantity, amount, category) values
('Captain America Tutu Dress Costume', 23.10, 80, (23.10 * 80), 1);

insert into products(name, price, quantity, amount, category) values
('Harry Potter DVD', 20, 10, (20 * 10), 2);
insert into products(name, price, quantity, amount, category) values
('Game of Thrones DVD', 40, 20, (40 * 20), 2);

insert into products(name, price, quantity, amount, category) values
('Minnie Mouse Icon Ring', 75.00, 5, (75.00 * 5), 3);
insert into products(name, price, quantity, amount, category) values
('Evil Queen Collar Necklace', 600000, 2, (600000 * 2), 3);

insert into products(name, price, quantity, amount, category) values
('The Little Mermaid Crossbody Bag', 228, 12, (228 * 12), 4);
insert into products(name, price, quantity, amount, category) values
('The Haunted Mansion 50th Anniversary Tote', 248, 8, (248 * 8), 4);
