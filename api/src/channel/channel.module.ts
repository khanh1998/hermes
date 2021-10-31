import { Module } from '@nestjs/common';
import PrismaModule from 'src/myprisma/prisma.module';
import { ChannelController } from './channel.controller';
import { ChannelService } from './channel.service';

@Module({
  controllers: [ChannelController],
  providers: [ChannelService],
  exports: [ChannelService],
  imports: [PrismaModule],
})
export class ChannelModule {}
