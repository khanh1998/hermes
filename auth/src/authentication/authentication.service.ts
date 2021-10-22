import { Injectable } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { TokenType } from 'src/token/token.entity';
import { TokenService } from 'src/token/token.service';
import { User } from 'src/user/user.entity';
import { UserService } from 'src/user/user.service';

@Injectable()
export class AuthenticationService {
  constructor(
    private userService: UserService,
    private jwtService: JwtService,
    private tokenService: TokenService,
  ) {}

  async validateUser(username: string, password: string): Promise<User> {
    const user = await this.userService.findOneByUsername(username);
    if (user.password === password) {
      return user;
    }
    return null;
  }

  async generateToken(user: User) {
    const payload = {
      username: user.username,
      id: user.id,
      type: TokenType.MAIN_APP.toString(),
    };
    return {
      main_token: this.jwtService.sign(payload),
    };
  }
}
