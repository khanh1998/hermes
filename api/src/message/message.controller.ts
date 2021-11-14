import { Controller, Get, Param, Query } from '@nestjs/common';
import { MessageService } from './message.service';

@Controller({
  path: '/message',
})
export class MessageController {
  constructor(private messageService: MessageService) {}
  @Get()
  message(@Query() param: { q: string; clan: number }) {
    const { q, clan } = param;
    console.log('search for ', q);
    return this.messageService.filter(q, clan);
  }
}
