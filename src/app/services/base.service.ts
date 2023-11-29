import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { API_URL } from './constants';
import { IPost } from '../shared/types';

@Injectable({
  providedIn: 'root',
})
export class BaseService {
  constructor(private readonly http: HttpClient) {}

  getAllPosts() {
    return this.http.get<{ posts: IPost[] }>(`${API_URL}/`);
  }

  getPostById(id: string) {
    return this.http.get<{ post: IPost }>(`${API_URL}/get/${id}`);
  }

  addPost(data: any) {
    const formData = new FormData();

    formData.append('title', data.Title!);
    formData.append('textContent', data.TextContent!);
    formData.append('imageContent', data.ImageContent!);

    return this.http.post(`${API_URL}/create`, formData, {
      withCredentials: true,
    });
  }

  deletePostById(id: string) {
    return this.http.delete(`${API_URL}/delete/${id}`, {
      withCredentials: true,
    });
  }
}
