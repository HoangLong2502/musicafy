-- Xóa các foreign key
ALTER TABLE "transaction" DROP CONSTRAINT IF EXISTS "transaction_package_fkey";
ALTER TABLE "transaction" DROP CONSTRAINT IF EXISTS "transaction_payment_fkey";
ALTER TABLE "payment" DROP CONSTRAINT IF EXISTS "payment_user_fkey";
ALTER TABLE "package_price" DROP CONSTRAINT IF EXISTS "package_price_package_fkey";
ALTER TABLE "users" DROP CONSTRAINT IF EXISTS "users_package_fkey";

-- Xóa các bảng theo thứ tự (từ bảng con đến bảng cha)
DROP TABLE IF EXISTS "transaction";
DROP TABLE IF EXISTS "payment";
DROP TABLE IF EXISTS "package_price";
DROP TABLE IF EXISTS "packages";

-- Xóa các cột trong bảng users
ALTER TABLE "users" DROP COLUMN IF EXISTS "package_expire";
ALTER TABLE "users" DROP COLUMN IF EXISTS "package";