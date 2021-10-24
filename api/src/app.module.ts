import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ClanModule } from './clan/clan.module';

@Module({
  imports: [ClanModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
