import { User } from '.prisma/client';
import {
  Body,
  Controller,
  Get,
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
  // @UseGuards(JwtAuthGuard)
  async create(@Body() user: User): Promise<User> {
    return await this.userService.create(user);
  }

  @Get()
  // @UseGuards(JwtAuthGuard)
  async findOne(@Request() req): Promise<User> {
    return req.user;
  }
}
