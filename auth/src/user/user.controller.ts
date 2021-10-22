import {
  Body,
  Controller,
  Get,
  Post,
  Request,
  UseGuards,
} from '@nestjs/common';
import { JwtAuthGuard } from 'src/authentication/jwt.guard';
import { User } from './user.entity';
import { UserService } from './user.service';

@Controller({
  path: '/user',
})
export class UserController {
  constructor(private userService: UserService) {}
  @Post()
  async create(@Body() user: User): Promise<User> {
    return await this.userService.create(user);
  }

  // @Get()
  async findAll(): Promise<Array<User>> {
    return await this.userService.findAll();
  }

  @UseGuards(JwtAuthGuard)
  @Get()
  async findOne(@Request() req): Promise<User> {
    return req.user;
  }
}
