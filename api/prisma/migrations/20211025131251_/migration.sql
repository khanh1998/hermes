/*
  Warnings:

  - You are about to drop the column `clanId` on the `User` table. All the data in the column will be lost.

*/
-- DropForeignKey
ALTER TABLE "User" DROP CONSTRAINT "User_foundedClanId_fkey";

-- DropIndex
DROP INDEX "User_foundedClanId_key";

-- AlterTable
ALTER TABLE "User" DROP COLUMN "clanId";

-- AddForeignKey
ALTER TABLE "Clan" ADD CONSTRAINT "Clan_chiefId_fkey" FOREIGN KEY ("chiefId") REFERENCES "User"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
