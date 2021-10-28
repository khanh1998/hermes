import { Injectable, CanActivate, ExecutionContext } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { Response } from 'express';
import { firstValueFrom, Observable } from 'rxjs';
import { TokenService } from 'src/token/token.service';
import { User } from 'src/user/user.entity';
import { UserService } from 'src/user/user.service';

@Injectable()
export class WsGuard implements CanActivate {
  constructor(
    private tokenService: TokenService,
    private jwtService: JwtService,
    private userService: UserService,
  ) {}
  async canActivate(host: ExecutionContext): Promise<boolean> {
    const ctx = host.switchToHttp();
    const request = ctx.getRequest<Request>();
    const response = ctx.getResponse<Response>();
    const bearerToken: string = request.headers['authorization'];
    const token = bearerToken.substr(7);
    let isTokenValid;
    try {
      isTokenValid = await this.jwtService.verifyAsync(token, {
        secret: process.env.JWT_SECRET,
      });
    } catch (error) {
      return false;
    }
    console.log(isTokenValid);
    if (isTokenValid) {
      // TODO: using transaction right here
      const res = await this.tokenService.findOne(token);
      await this.tokenService.remove(token);
      if (res) {
        const userObser = this.userService.findOneByUsername(
          isTokenValid.username,
        );
        const user: User = await firstValueFrom(userObser);
        console.log(user);
        request.headers['user'] = user;
        return true;
      }
    }
    return false;
  }
}
