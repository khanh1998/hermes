import { Channel, Prisma } from '.prisma/client';
import { Injectable } from '@nestjs/common';
import { PrismaService } from 'src/myprisma/prisma.service';

@Injectable()
export class ChannelService {
  constructor(private prisma: PrismaService) {}

  async create(data: Prisma.ChannelCreateInput): Promise<Channel> {
    return await this.prisma.channel.create({ data });
  }

  async findAll(params: {
    skip?: number;
    take?: number;
    cursor?: Prisma.ChannelWhereUniqueInput;
    where?: Prisma.ChannelWhereInput;
    orderBy?: Prisma.ChannelOrderByWithRelationInput;
  }): Promise<Array<Channel>> {
    const { cursor, where, orderBy, skip, take } = params;
    return await this.prisma.channel.findMany({
      cursor,
      where,
      orderBy,
      skip,
      take,
    });
  }
}
