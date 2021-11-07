import { Controller, Get, Post, Request, UseGuards } from '@nestjs/common';
import { JwtAuthGuard } from './jwt.guard';
import { AuthenticationService } from './authentication.service';
import { LocalAuthGuard } from './local-auth.guard';
import { WsGuard } from './ws.guard';
import { User } from 'src/user/user.entity';

@Controller({
  path: '/authentication',
})
export class AuthenticationController {
  constructor(private readonly authService: AuthenticationService) {}

  @UseGuards(LocalAuthGuard)
  @Post('/main')
  async login(@Request() req) {
    return this.authService.generateToken(req.user);
  }

  @UseGuards(WsGuard)
  @Post('/ws')
  async loginWebsocket(@Request() req: Request) {
    const user: User = req.headers['user'];
    console.log('login ws ok', user);
    return user;
  }
}
