import { Body, Controller, Get, Post, Query } from '@nestjs/common';
import { ChannelService } from './channel.service';

@Controller({
  path: '/channel',
})
export class ChannelController {
  constructor(private channelService: ChannelService) {}
  @Post()
  createChannel(@Body() body: CreateChanneDTO) {
    const { clanId, name } = body;
    return this.channelService.create({
      clan: { connect: { id: clanId } },
      name,
    });
  }

  @Get()
  getChannel(@Query('clanId') clanId: number) {
    return this.channelService.findAll({ where: { clanId: Number(clanId) } });
  }
}
