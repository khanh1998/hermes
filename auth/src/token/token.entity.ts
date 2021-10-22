import { Entity, Column, PrimaryGeneratedColumn } from 'typeorm';

export enum TokenType {
  MAIN_APP,
  WEBSOCKET,
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
