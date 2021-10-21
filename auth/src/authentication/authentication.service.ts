import { Injectable } from '@nestjs/common';
import { User } from 'src/user/user.entity';
import { UserService } from 'src/user/user.service';
import { AuthenticationDTO } from './authentication.req.dto';

@Injectable()
export class AuthenticationService {
  constructor(private userService: UserService) {}
  async login(auth: AuthenticationDTO): Promise<string> {
    const user = await this.userService.findOneByUsername(auth.username);
    if (user.password === auth.password) {
      return 'token';
    }
    return null;
  }
}
