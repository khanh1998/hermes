import { User } from '.prisma/client';
import {
  Body,
  Controller,
  Get,
  Param,
  Post,
  Req,
  UseGuards,
} from '@nestjs/common';
import { Request } from 'express';
import { JwtAuthGuard } from 'src/auth/jwt.guard';
import { ClanService } from 'src/clan/clan.service';
import { UserService } from './user.service';

@Controller({
  path: '/user',
})
export class UserController {
  constructor(
    private userService: UserService,
    private clanService: ClanService,
  ) {}
  @Post()
  async create(
    @Body() user: { username: string; fullname: string; password: string },
  ): Promise<User> {
    const { fullname, password, username } = user;
    return await this.userService.create({ fullname, password, username });
  }

  @Get()
  @UseGuards(JwtAuthGuard)
  async findCurrentUser(@Req() req: Request): Promise<User> {
    const username: string = (req as any).user.username;
    return await this.userService.findOne({ username });
  }

  @Get('/:username')
  async findOne(@Param('username') username: string): Promise<User> {
    return this.userService.findOne({ username });
  }

  @Get('/:username/clan')
  async findAllClanOfUser(@Param('username') username: string) {
    return this.clanService.findAll({
      where: { members: { some: { username } } },
    });
  }
}
