-- generated by github.com/mackee/go-genddl. DO NOT EDIT!!!

DROP TABLE IF EXISTS "group";

CREATE TABLE "group" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "name" TEXT NOT NULL,
    "leader_user_id" INTEGER NOT NULL,
    "sub_leader_user_id" INTEGER NULL,
    "child_group_id" INTEGER NULL,
    "created_at" DATETIME NOT NULL,
    "updated_at" DATETIME NULL
) ;


DROP TABLE IF EXISTS "user";

CREATE TABLE "user" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "name" TEXT NOT NULL,
    "age" INTEGER NULL,
    "rate" REAL NOT NULL DEFAULT 0,
    "icon_image" BLOB NOT NULL,
    "created_at" DATETIME NOT NULL,
    "updated_at" DATETIME NULL,
    UNIQUE ("name")
) ;


DROP TABLE IF EXISTS "user_external";

CREATE TABLE "user_external" (
    "id" INTEGER NOT NULL PRIMARY KEY,
    "user_id" INTEGER NOT NULL,
    "icon_image" BLOB NULL,
    "created_at" DATETIME NOT NULL,
    "updated_at" DATETIME NOT NULL
) ;


DROP TABLE IF EXISTS "user_item";

CREATE TABLE "user_item" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "user_id" INTEGER NOT NULL,
    "item_id" TEXT NOT NULL,
    "is_used" INTEGER NOT NULL,
    "has_extension" INTEGER NULL,
    "used_at" DATETIME NULL
) ;


DROP TABLE IF EXISTS "user_sns";

CREATE TABLE "user_sns" (
    "id" INTEGER NOT NULL PRIMARY KEY,
    "sns_type" TEXT NOT NULL,
    "created_at" DATETIME NOT NULL,
    "updated_at" DATETIME NOT NULL
) ;

