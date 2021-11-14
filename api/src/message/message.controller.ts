import { Controller, Get, Query } from '@nestjs/common';
import { MessageService } from './message.service';

@Controller({
  path: '/message',
})
export class MessageController {
  constructor(private messageService: MessageService) {}
  @Get('/search')
  message(@Query() query: { q: string; clan: number; channel: number }) {
    const { q, clan, channel } = query;
    console.log('search for ', q);
    return this.messageService.filter(clan, channel, q);
  }

  @Get('/latest')
  getLatestMessages(
    @Query() query: { clan: number; limit: number; channel: number },
  ) {
    const { limit, clan, channel } = query;
    return this.messageService.getLastest(clan, channel, limit);
  }

  @Get('/before')
  getMessageBefore(
    @Query()
    query: {
      time: number;
      limit: number;
      clan: number;
      channel: number;
    },
  ) {
    const { limit, time, clan, channel } = query;
    return this.messageService.getBefore(clan, channel, time, limit);
  }
}
