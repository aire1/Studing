--Основная таблица с пользователями
create table if not exists users (
	id serial primary key,
	username varchar(50) not null unique,
	password_hash text not null,
	deleted bool default false,
	online bool default false,
	created_at timestamp default now(),
	updated_at timestamp default now()
);

--Основная таблица с заметками
create table if not exists notes (
	id serial primary key,
	name varchar(50) not null,
	text text not null,
	owner_id int not null,
	deleted bool default false,
	foreign key (owner_id) references users(id) on delete cascade,
	created_at timestamp default now(),
	updated_at timestamp default now()
);

--Справочная таблица с возможными действиями
create table if not exists actions (
	id serial primary key,
	name varchar(50) unique not null
);

--Таблица с историей действий пользователя
create table if not exists history (
	id serial primary key,
	user_id int not null,
	action_id int not null,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	
	foreign key (user_id) references users(id) on delete cascade,
	foreign key (action_id) references actions(id) on delete cascade
);


--Триггеры для таблиц, используются только для автообновления поля updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

create or replace trigger set_timestamp
before update on users
for each row
execute function update_updated_at_column();

create or replace trigger set_timestamp
before update on history
for each row
execute function update_updated_at_column();

create or replace trigger set_timestamp
before update on notes
for each row
execute function update_updated_at_column();

--Заполнение справочниов
insert into actions (name) values ('Не в сети') on conflict do nothing;
insert into actions (name) values ('В сети') on conflict do nothing;
insert into actions (name) values ('Добавил заметку') on conflict do nothing;
insert into actions (name) values ('Удалил заметку') on conflict do nothing;
insert into actions (name) values ('Изменил заметку') on conflict do nothing;

