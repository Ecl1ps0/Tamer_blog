import { Injectable, WritableSignal, signal } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { API_URL } from './constants';

interface initialState {
  isAuthorized: boolean;
}

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  state: WritableSignal<initialState> = signal({
    isAuthorized: localStorage.getItem('isTamer') ? true : false,
  });

  constructor(private readonly http: HttpClient) {}

  signIn(password: string) {
    const formData = new FormData();
    formData.append('password', password);

    const isAuth = this.http
      .post(`${API_URL}/authorize`, formData, {
        withCredentials: true,
      })
      .subscribe({
        next(val) {
          return val;
        },
        error() {
          return null;
        },
      });

    if (isAuth) {
      localStorage.setItem('isTamer', 'true');
      this.state.update(() => ({
        isAuthorized: true,
      }));
    }

    return 'Authorized';
  }
}
