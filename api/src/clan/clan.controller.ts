import { Clan, Prisma, User } from '.prisma/client';
import {
  Body,
  Controller,
  Get,
  Param,
  Post,
  Query,
  Req,
  UseGuards,
} from '@nestjs/common';
import { JwtAuthGuard } from 'src/auth/jwt.guard';
import { ClanService } from './clan.service';

@Controller({
  path: '/clan',
})
export class ClanController {
  constructor(private readonly clanService: ClanService) {}

  @Get('/:id')
  getClanById(@Param('id') id: number): Promise<Clan> {
    return this.clanService.findOne({ id: Number(id) });
  }

  @Post()
  @UseGuards(JwtAuthGuard)
  createClan(
    @Req() req,
    @Body() clan: { domain: string; name: string },
  ): Promise<Clan> {
    const { domain, name } = clan;
    const requesterId = req.user.id;
    return this.clanService.createClanAndDefaultChannel({
      domain,
      name,
      chief: { connect: { id: requesterId } },
    });
  }

  @Get()
  getClans(): Promise<Clan[]> {
    return this.clanService.findAll({});
  }

  @Post('/:clanId/user/:username')
  addUser(
    @Param('clanId') clanId: number,
    @Param('username') username: string,
  ) {
    return this.clanService.update({
      where: { id: Number(clanId) },
      data: {
        members: { connect: { username } },
      },
    });
  }
}
