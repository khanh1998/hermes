import { Injectable } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { map, Observable } from 'rxjs';
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

  validateUser(username: string, password: string): Observable<User> {
    const obser = this.userService.findOneByUsername(username);
    return obser.pipe(
      map((user: User) => {
        if (user.password === password) {
          return user;
        }
        return null;
      }),
    );
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
