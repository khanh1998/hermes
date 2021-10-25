import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { UserService } from './user.service';
import { HttpModule } from '@nestjs/axios';

@Module({
  imports: [
    HttpModule.register({
      baseURL: process.env.API_HOST,
    }),
  ],
  providers: [UserService],
  controllers: [],
  exports: [UserService],
})
export class UsersModule {}
