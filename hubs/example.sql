DROP TABLE IF EXISTS "hubs" CASCADE;

CREATE TABLE "hubs" (
  "id" SERIAL PRIMARY KEY,
  "name" TEXT,
  "location" GEOGRAPHY,
  "radius" INTEGER
);

DROP TABLE IF EXISTS "posts";

CREATE TABLE "posts" (
  "id" SERIAL PRIMARY KEY,
  "hub_id" INTEGER,
  "title" TEXT,
  "content" TEXT
);