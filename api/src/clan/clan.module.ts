import { Module } from '@nestjs/common';
import { JwtStrategy } from 'src/auth/jwt.strategy';
import PrismaModule from 'src/myprisma/prisma.module';
import { ClanController } from './clan.controller';
import { ClanService } from './clan.service';

@Module({
  controllers: [ClanController],
  providers: [ClanService, JwtStrategy],
  imports: [PrismaModule],
  exports: [ClanService],
})
export class ClanModule {}
