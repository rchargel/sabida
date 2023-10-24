-- +goose Up

CREATE TABLE organizations (
	id			UUID 			NOT NULL DEFAULT gen_random_uuid(),
	name 		VARCHAR(200) 	NOT NULL UNIQUE,
	created_at	timestamp 		NOT NULL DEFAULT now(),
	updated_at	timestamp 		NOT NULL DEFAULT now(),
	deleted_at	timestamp,
	
	PRIMARY KEY(id)
);
CREATE INDEX idx_organizations_deleted_at ON organizations(deleted_at);

insert into organizations (name) values ('SYSTEM');

CREATE TABLE users (
	id 			UUID 			NOT NULL DEFAULT gen_random_uuid(),
	username 	varchar(200) 	NOT NULL UNIQUE,
	email 		varchar(200) 	NOT NULL UNIQUE,
	password 	varchar(200),
	created_at 	timestamp 		NOT NULL DEFAULT now(),
	updated_at 	timestamp		NOT NULL DEFAULT now(),
	deleted_at 	timestamp,
	
	PRIMARY KEY(id)
);

CREATE INDEX idx_users_deleted_at ON users(deleted_at);
insert into users (username, email) values ('SYSTEM', 'SYSTEM@SYSTEM.COM');

CREATE TABLE user_organizations (
	user_id 			UUID 		NOT NULL,
	organization_id 	UUID 		NOT NULL,
	created_at 			timestamp 	NOT NULL DEFAULT now(),
	updated_at 			timestamp 	NOT NULL DEFAULT now(),
	deleted_at 			timestamp,
	created_by 			UUID 		NOT NULL,
	updated_by 			UUID 		NOT NULL,
	deleted_by 			UUID,
	
	CONSTRAINT fk_user_organizations_user_id
		FOREIGN KEY(user_id)
		REFERENCES users(id),
	CONSTRAINT fk_user_organizations_organization_id
		FOREIGN KEY(organization_id)
		REFERENCES organizations(id),
	CONSTRAINT fk_user_organizations_created_by
		FOREIGN KEY(created_by)
		REFERENCES users(id),
	CONSTRAINT fk_user_organizations_updated_by
		FOREIGN KEY(updated_by)
		REFERENCES users(id),
	CONSTRAINT fk_user_organizations_deleted_by
		FOREIGN KEY(deleted_by)
		REFERENCES users(id),
	CONSTRAINT pk_user_organizations
		UNIQUE(user_id, organization_id)
);

CREATE INDEX idx_user_organizations_deleted_at ON user_organizations(deleted_at);

insert into user_organizations (user_id, organization_id, created_by, updated_by) VALUES (
	(select id from users where username = 'SYSTEM'),
	(select id from organizations where name = 'SYSTEM'),
	(select id from users where username = 'SYSTEM'),
	(select id from users where username = 'SYSTEM')
);

-- +goose Down

DROP TABLE user_organizations;
DROP TABLE users;
DROP TABLE organizations;