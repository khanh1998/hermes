import { Module } from '@nestjs/common';
import { JwtModule } from '@nestjs/jwt';
import { TypeOrmModule } from '@nestjs/typeorm';
import { TokenController } from './token.controller';
import { Token } from './token.entity';
import { TokenService } from './token.service';

@Module({
  imports: [
    TypeOrmModule.forFeature([Token]),
    JwtModule.register({
      secret: 'jwtsecretkey',
      signOptions: {
        expiresIn: '1 hour',
      },
    }),
  ],
  providers: [TokenService],
  controllers: [TokenController],
  exports: [TokenService, TypeOrmModule],
})
export class TokenModule {}
