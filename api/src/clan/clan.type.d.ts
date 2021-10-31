import { Type } from 'class-transformer';
import { IsInt } from 'class-validator';

export class ClanQuery {
  @IsInt()
  @Type(() => Number)
  clanId: number;

  @IsInt()
  @Type(() => Number)
  userId: number;
}
