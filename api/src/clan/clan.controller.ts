import { Clan, User } from '.prisma/client';
import {
  Body,
  Controller,
  Get,
  Param,
  Post,
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
  @UseGuards(JwtAuthGuard)
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
    return this.clanService.create({ domain, name, chief: req.user });
  }
}
