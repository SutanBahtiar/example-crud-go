CREATE TABLE "employee"(
    "id" bigserial PRIMARY KEY,
    "code" varchar NOT NULL,
    "name" varchar NOT NULL,
    "email" varchar NOT NULL,
    "phone_number" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "employee"("name");
CREATE INDEX ON "employee"("email");
CREATE INDEX ON "employee"("phone_number");