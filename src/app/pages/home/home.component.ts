import { Component } from '@angular/core';
import { Subject, switchMap } from 'rxjs';
import { BaseService } from 'src/app/services/base.service';
import { IPost } from 'src/app/shared/types';
@Component({
  selector: 'home-page',
  templateUrl: './home.component.html',
})
export class HomeComponent {
  posts$: Subject<IPost[]> = new Subject();
  isLoading: boolean = true;

  constructor(private base: BaseService) {
    this.base.getAllPosts().subscribe((res) => {
      this.posts$.next(res.posts);
      this.isLoading = false;
    });
  }

  handleDeletePost(id: string) {
    this.isLoading = true;
    this.base
      .deletePostById(id)
      .pipe(switchMap(() => this.base.getAllPosts()))
      .subscribe({
        next: (res) => {
          this.posts$.next(res.posts);
        },
        error(err) {
          console.log(err);
        },
        complete: () => {
          this.isLoading = false;
        },
      });
  }
}
