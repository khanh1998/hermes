import { Module } from '@nestjs/common';
import { UsersModule } from 'src/user/user.module';
import { AuthenticationController } from './authentication.controller';
import { AuthenticationService } from './authentication.service';

@Module({
  controllers: [AuthenticationController],
  providers: [AuthenticationService],
  imports: [UsersModule],
})
export class AuthenticationModule {}
