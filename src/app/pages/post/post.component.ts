import { Component, OnInit, WritableSignal, signal } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { BaseService } from 'src/app/services/base.service';
import { IPost } from 'src/app/shared/types';
import { convertTimeStampToDate } from 'src/app/shared/utils';

@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
})
export class PostComponent implements OnInit {
  postDate: WritableSignal<string> = signal('');
  post: WritableSignal<IPost> = signal({
    CreateAt: 0,
    ID: '',
    ImageContent: '',
    TextContent: '',
    Title: '',
    UpdatedAt: 0,
  });
  constructor(private route: ActivatedRoute, private base: BaseService) {}

  ngOnInit(): void {
    this.route.params.subscribe((params) => {
      this.base.getPostById(params['id']).subscribe((res) => {
        this.post.set(res.post);
        this.postDate.set(convertTimeStampToDate(res.post.CreateAt));
      });
    });
  }
}
