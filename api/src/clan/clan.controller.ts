import { Clan, User } from '.prisma/client';
import { Body, Controller, Get, Param, Post } from '@nestjs/common';
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
  createClan(@Body() clan: { domain: string; name: string }): Promise<Clan> {
    const { domain, name } = clan;
    const chiefId = 1;
    return this.clanService.create({ domain, name, chiefId });
  }
}
