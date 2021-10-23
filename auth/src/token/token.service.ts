import { Injectable } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { InjectRepository } from '@nestjs/typeorm';
import { User } from 'src/user/user.entity';
import { Repository } from 'typeorm';
import { Token, TokenType } from './token.entity';

@Injectable()
export class TokenService {
  constructor(
    @InjectRepository(Token)
    private tokenRepository: Repository<Token>,
    private jwtService: JwtService,
  ) {}

  async create(user: User): Promise<Token> {
    const payload = {
      username: user.username,
      id: user.id,
      type: TokenType.WEBSOCKET.toString(),
    };
    const tokenStr = this.jwtService.sign(payload, { expiresIn: '5m' });
    const token = {
      id: null,
      token: tokenStr,
      type: TokenType.WEBSOCKET,
    };
    return await this.tokenRepository.save(token);
  }

  async findOne(token: string): Promise<Token> {
    return await this.tokenRepository.findOne({ where: { token } });
  }

  async remove(token: string): Promise<void> {
    await this.tokenRepository.delete({ token });
  }
}
