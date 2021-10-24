import { Injectable, CanActivate, ExecutionContext } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { Observable } from 'rxjs';
import { TokenService } from 'src/token/token.service';

@Injectable()
export class WsGuard implements CanActivate {
  constructor(
    private tokenService: TokenService,
    private jwtService: JwtService,
  ) {}
  async canActivate(host: ExecutionContext): Promise<boolean> {
    const ctx = host.switchToHttp();
    const request = ctx.getRequest<Request>();
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
        return true;
      }
    }
    return false;
  }
}
