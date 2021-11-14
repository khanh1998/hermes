import { Module } from '@nestjs/common';
import { ElasticsearchModule } from '@nestjs/elasticsearch';
import { JwtStrategy } from 'src/auth/jwt.strategy';
import { MessageController } from './message.controller';
import { MessageService } from './message.service';

@Module({
  imports: [
    ElasticsearchModule.register({
      node: process.env.ELASTIC_SEARCH_URI,
    }),
  ],
  providers: [JwtStrategy, MessageService],
  controllers: [MessageController],
  exports: [],
})
export class MessageModule {}
