import { Entity, Column, PrimaryGeneratedColumn } from 'typeorm';

export enum TokenType {
  MAIN_APP = 'MAIN_ACCESS_TOKEN',
  WEBSOCKET = 'WEBSOCKET_AUTH_TOKEN',
}

@Entity()
export class Token {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ unique: true })
  token: string;

  @Column()
  type: TokenType;
}
