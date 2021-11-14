import { Injectable } from '@nestjs/common';
import { ElasticsearchService } from '@nestjs/elasticsearch';

@Injectable()
export class MessageService {
  constructor(private readonly elasticsearchService: ElasticsearchService) {}

  filter(clanId: number, channelId: number, search: string) {
    return this.elasticsearchService.search(
      {
        index: `clan_${clanId}`,
        type: 'message',
        body: {
          size: 10,
          from: 0,
          query: {
            bool: {
              must: [
                {
                  match: {
                    message: search,
                  },
                },
                {
                  match: {
                    channelId: channelId,
                  },
                },
              ],
            },
          },
        },
      },
      {},
    );
  }

  getLastest(clanId: number, channelId: number, limit: number) {
    return this.elasticsearchService.search({
      index: `clan_${clanId}`,
      type: 'message',
      body: {
        size: limit,
        sort: { time: 'desc' },
        query: {
          bool: {
            must: { match: { channelId: channelId } },
          },
        },
      },
    });
  }

  getBefore(clanId: number, channelId: number, time: number, limit: number) {
    return this.elasticsearchService.search({
      index: `clan_${clanId}`,
      type: 'message',
      body: {
        size: limit,
        query: {
          bool: {
            must: [
              {
                range: {
                  time: {
                    lt: time,
                  },
                },
              },
              { match: { channelId: channelId } },
            ],
          },
        },
      },
    });
  }
}
