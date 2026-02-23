-- +goose up
CREATE TABLE IF NOT EXISTS "Sentences" (
    "id" INTEGER NOT NULL UNIQUE,
    "text" TEXT NOT NULL UNIQUE,
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "Words" (
    "id" INTEGER NOT NULL UNIQUE,
    "word" TEXT NOT NULL,
    "meaning" TEXT,
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "Notes" (
    "id" INTEGER NOT NULL UNIQUE,
    "sentence_id" INTEGER,
    "target_id" INTEGER,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("sentence_id") REFERENCES "Sentences" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("target_id") REFERENCES "Words" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- +goose down
DROP TABLE "Notes";

DROP TABLE "Words"
DROP TABLE "Sentences"
