import { User } from '.prisma/client';
import {
  Body,
  Controller,
  Get,
  Param,
  Post,
  Request,
  UseGuards,
} from '@nestjs/common';
import { JwtAuthGuard } from 'src/auth/jwt.guard';
import { UserService } from './user.service';

@Controller({
  path: '/user',
})
export class UserController {
  constructor(private userService: UserService) {}
  @Post()
  async create(
    @Body() user: { username: string; fullname: string; password: string },
  ): Promise<User> {
    const { fullname, password, username } = user;
    return await this.userService.create({ fullname, password, username });
  }

  @Get('/:username')
  async findOne(@Param('username') username: string): Promise<User> {
    return this.userService.findOne({ username });
  }
}
