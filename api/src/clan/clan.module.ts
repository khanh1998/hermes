import { Module } from '@nestjs/common';
import PrismaModule from 'src/prisma/prisma.module';
import { ClanController } from './clan.controller';
import { ClanService } from './clan.service';

@Module({
  controllers: [ClanController],
  providers: [ClanService],
  imports: [PrismaModule],
})
export class ClanModule {}
