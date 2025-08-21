ALTER TABLE "todos" DROP CONSTRAINT todos_user_id_fkey;

ALTER TABLE "todos"
ADD CONSTRAINT todos_user_id_fkey
FOREIGN KEY ("user_id")
REFERENCES "users" ("id")
ON DELETE CASCADE;