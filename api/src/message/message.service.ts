import { Injectable } from '@nestjs/common';
import { ElasticsearchService } from '@nestjs/elasticsearch';

@Injectable()
export class MessageService {
  constructor(private readonly elasticsearchService: ElasticsearchService) {}

  filter(search: string, clanId: number) {
    return this.elasticsearchService.search(
      {
        index: `clan_${clanId}`,
        type: 'message',
        body: {
          size: 200,
          from: 0,
          query: {
            match: {
              message: search,
            },
          },
        },
      },
      {},
    );
  }
}
