import { Clan, Prisma } from '.prisma/client';
import { Injectable } from '@nestjs/common';
import { PrismaService } from 'src/myprisma/prisma.service';

@Injectable()
export class ClanService {
  constructor(private prisma: PrismaService) {}

  async findOne(
    clanWhereUniqueInput: Prisma.ClanWhereUniqueInput,
  ): Promise<Clan | null> {
    return await this.prisma.clan.findUnique({
      where: clanWhereUniqueInput,
      include: {
        channels: {},
        chief: {},
        members: {},
      },
    });
  }

  async findAll(params: {
    skip?: number;
    take?: number;
    cursor?: Prisma.ClanWhereUniqueInput;
    where?: Prisma.ClanWhereInput;
    orderBy?: Prisma.ClanOrderByWithRelationInput;
  }): Promise<Clan[]> {
    const { skip, take, cursor, where, orderBy } = params;
    return this.prisma.clan.findMany({
      skip,
      take,
      cursor,
      where,
      orderBy,
    });
  }

  async create(clan: Prisma.ClanCreateInput): Promise<Clan | null> {
    return await this.prisma.clan.create({ data: clan });
  }

  async createClanAndDefaultChannel(
    clan: Prisma.ClanCreateInput,
  ): Promise<Clan | null> {
    return await this.prisma.clan.create({
      data: {
        ...clan,
        channels: {
          create: {
            name: 'Origin',
          },
        },
      },
    });
  }

  async update(params: {
    where: Prisma.ClanWhereUniqueInput;
    data: Prisma.ClanUpdateInput;
  }): Promise<Clan> {
    const { where, data } = params;
    return this.prisma.clan.update({
      data,
      where,
    });
  }

  async delete(where: Prisma.ClanWhereUniqueInput): Promise<Clan> {
    return this.prisma.clan.delete({
      where,
    });
  }
}
