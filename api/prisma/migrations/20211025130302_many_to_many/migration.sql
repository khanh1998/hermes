-- DropForeignKey
ALTER TABLE "User" DROP CONSTRAINT "User_clanId_fkey";

-- CreateTable
CREATE TABLE "_ClanToUser" (
    "A" INTEGER NOT NULL,
    "B" INTEGER NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "_ClanToUser_AB_unique" ON "_ClanToUser"("A", "B");

-- CreateIndex
CREATE INDEX "_ClanToUser_B_index" ON "_ClanToUser"("B");

-- AddForeignKey
ALTER TABLE "_ClanToUser" ADD FOREIGN KEY ("A") REFERENCES "Clan"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_ClanToUser" ADD FOREIGN KEY ("B") REFERENCES "User"("id") ON DELETE CASCADE ON UPDATE CASCADE;
