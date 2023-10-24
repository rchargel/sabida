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
('front-end', (select id from organizations where name = 'system'), (select id from users where username = 'system'), (select id from users where username = 'system')),
('back-end', (select id from organizations where name = 'system'), (select id from users where username = 'system'), (select id from users where username = 'system')),
('framework', (select id from organizations where name = 'system'), (select id from users where username = 'system'), (select id from users where username = 'system')),
('language', (select id from organizations where name = 'system'), (select id from users where username = 'system'), (select id from users where username = 'system')),
('database', (select id from organizations where name = 'system'), (select id from users where username = 'system'), (select id from users where username = 'system')),
('no-sql', (select id from organizations where name = 'system'), (select id from users where username = 'system'), (select id from users where username = 'system')),
('caching', (select id from organizations where name = 'system'), (select id from users where username = 'system'), (select id from users where username = 'system')),
('rest', (select id from organizations where name = 'system'), (select id from users where username = 'system'), (select id from users where username = 'system'));

-- +goose Down

DROP TABLE categories;