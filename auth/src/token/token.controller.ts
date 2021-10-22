import {
  Body,
  Controller,
  Get,
  Post,
  Request,
  UseGuards,
} from '@nestjs/common';
import { JwtAuthGuard } from 'src/authentication/jwt.guard';
import { TokenService } from './token.service';

@Controller({
  path: '/token',
})
export class TokenController {
  constructor(private tokenService: TokenService) {}

  @UseGuards(JwtAuthGuard)
  @Post('/ws')
  async loginWebsocket(@Request() req) {
    const wsToken = await this.tokenService.create(req.user);
    if (wsToken) {
      return {
        ws_token: wsToken.token,
      };
    }
    return;
  }
}
