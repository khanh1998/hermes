import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { AuthenticationModule } from './authentication/authentication.module';
import { Token } from './token/token.entity';
import { TokenModule } from './token/token.module';
import { User } from './user/user.entity';
import { UsersModule } from './user/user.module';

@Module({
  imports: [
    TypeOrmModule.forRoot({
      type: 'postgres',
      host: 'localhost',
      port: 5433,
      username: 'khanh',
      password: 'handsome',
      database: 'auth',
      entities: [User, Token],
      synchronize: true,
    }),
    AuthenticationModule,
    UsersModule,
    TokenModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
