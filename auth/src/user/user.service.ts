import { HttpService } from '@nestjs/axios';
import { Injectable } from '@nestjs/common';
import { AxiosResponse } from 'axios';
import { map, Observable } from 'rxjs';
import { User } from './user.entity';

@Injectable()
export class UserService {
  constructor(private httpService: HttpService) {}

  findOne(id: string): Observable<any> {
    const res = this.httpService.get(`http://localhost:5000/user/${id}`);
    return res.pipe(map((res: AxiosResponse) => res.data));
  }

  findOneByUsername(username: string): Observable<User> {
    const res = this.httpService.get(`http://localhost:5000/user/${username}`);
    return res.pipe(map((res: AxiosResponse) => res.data as User));
  }
}
