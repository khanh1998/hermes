import { Body, Controller, HttpStatus, Post, Res } from '@nestjs/common';
import { Response } from 'express';
import { AuthenticationDTO } from './authentication.req.dto';
import { AuthenticationService } from './authentication.service';

@Controller({
  path: '/authentication',
})
export class AuthenticationController {
  constructor(private readonly auth: AuthenticationService) {}

  @Post()
  async getHello(@Body() auth: AuthenticationDTO, @Res() res: Response) {
    const token = await this.auth.login(auth);
    if (token) {
      return res.send(token);
    }
    return res.status(HttpStatus.UNAUTHORIZED).send('unauthorized');
  }
}
