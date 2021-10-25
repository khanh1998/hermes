import { Strategy } from 'passport-local';
import { PassportStrategy } from '@nestjs/passport';
import { Injectable, UnauthorizedException } from '@nestjs/common';
import { AuthenticationService } from './authentication.service';
import { lastValueFrom, Observable } from 'rxjs';
import { User } from 'src/user/user.entity';

@Injectable()
export class LocalStrategy extends PassportStrategy(Strategy) {
  constructor(private authService: AuthenticationService) {
    super();
  }

  async validate(username: string, password: string): Promise<any> {
    const userObservable: Observable<User> =
      await this.authService.validateUser(username, password);
    const user = await lastValueFrom(userObservable);
    if (!user) {
      throw new UnauthorizedException();
    }
    return user;
  }
}
