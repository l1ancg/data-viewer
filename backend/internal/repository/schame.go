package repository

var sqls = map[string]string{
	"1.0.0": `
CREATE TABLE IF NOT EXISTS "version" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "version" TEXT
);

CREATE TABLE IF NOT EXISTS "view" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "resource_id" TEXT,
	"resource_type" TEXT,
	"display_type" TEXT,
    "name" TEXT,
    "desc" TEXT
);

CREATE TABLE IF NOT EXISTS "resource" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "name" TEXT,
    "type" INTEGER,
    "data" TEXT
);

CREATE TABLE IF NOT EXISTS "column" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "view_id" TEXT,
    "dict_id" TEXT,
    "name" TEXT,
    "dataType" TEXT,
    "orderBy" TEXT,
    "display" INTEGER,
    "condition" INTEGER
);


CREATE TABLE IF NOT EXISTS "dict" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "name" TEXT
);


CREATE TABLE IF NOT EXISTS "dict_detail" (
    "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    "dict_id" INTEGER,
    "key" TEXT,
    "value" TEXT
);
`,
}
