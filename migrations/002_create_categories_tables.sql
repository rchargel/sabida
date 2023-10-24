-- +goose Up

CREATE TABLE categories(
	id 			SERIAL 			NOT NULL,
	name 		varchar(100) 	NOT NULL,
	belongs_to 	UUID			NOT NULL,
	created_at 	timestamp 		NOT NULL DEFAULT now(),
	updated_at 	timestamp 		NOT NULL DEFAULT now(),
	deleted_at 	timestamp,
	created_by 	UUID 			NOT NULL,
	updated_by 	UUID 			NOT NULL,
	
	PRIMARY KEY(id),
	CONSTRAINT fk_created_by 
		FOREIGN KEY(created_by)
		REFERENCES users(id),
	CONSTRAINT fk_updated_by
		FOREIGN KEY(updated_by)
		REFERENCES users(id),
	CONSTRAINT fk_belongs_to
		FOREIGN KEY(belongs_to)
		REFERENCES organizations(id),
    CONSTRAINT uq_name_for_org
    	UNIQUE(name, belongs_to)
);

CREATE INDEX idx_categories_deleted_at ON categories(deleted_at);

INSERT INTO categories (name, belongs_to, created_by, updated_by) VALUES
('front-end', (select id from organizations where name = 'SYSTEM'), (select id from users where username = 'SYSTEM'), (select id from users where username = 'SYSTEM')),
('back-end', (select id from organizations where name = 'SYSTEM'), (select id from users where username = 'SYSTEM'), (select id from users where username = 'SYSTEM')),
('framework', (select id from organizations where name = 'SYSTEM'), (select id from users where username = 'SYSTEM'), (select id from users where username = 'SYSTEM')),
('language', (select id from organizations where name = 'SYSTEM'), (select id from users where username = 'SYSTEM'), (select id from users where username = 'SYSTEM')),
('database', (select id from organizations where name = 'SYSTEM'), (select id from users where username = 'SYSTEM'), (select id from users where username = 'SYSTEM')),
('no-sql', (select id from organizations where name = 'SYSTEM'), (select id from users where username = 'SYSTEM'), (select id from users where username = 'SYSTEM')),
('caching', (select id from organizations where name = 'SYSTEM'), (select id from users where username = 'SYSTEM'), (select id from users where username = 'SYSTEM')),
('rest', (select id from organizations where name = 'SYSTEM'), (select id from users where username = 'SYSTEM'), (select id from users where username = 'SYSTEM'));

-- +goose Down

DROP TABLE categories;