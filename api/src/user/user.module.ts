import { Module } from '@nestjs/common';
import { UserService } from './user.service';
import { UserController } from './user.controller';
import PrismaModule from 'src/myprisma/prisma.module';
import { JwtStrategy } from 'src/auth/jwt.strategy';
import { ClanModule } from 'src/clan/clan.module';

@Module({
  imports: [PrismaModule, ClanModule],
  providers: [UserService, JwtStrategy],
  controllers: [UserController],
  exports: [UserService],
})
export class UserModule {}
