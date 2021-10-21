import { Body, Controller, Get, Post } from '@nestjs/common';
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

  @Get()
  async findAll(): Promise<Array<User>> {
    return await this.userService.findAll();
  }
}
