import {
  Component,
  EventEmitter,
  Input,
  OnChanges,
  Output,
} from '@angular/core';
import { CommonModule } from '@angular/common';
import { IPost } from '../../types';
import { RouterLinkWithHref, Router } from '@angular/router';
import { convertTimeStampToDate } from '../../utils';
import { AuthService } from 'src/app/services/auth.service';
import { ButtonComponent } from '../button/button.component';

@Component({
  selector: 'app-card',
  standalone: true,
  imports: [CommonModule, RouterLinkWithHref, ButtonComponent],
  templateUrl: './card.component.html',
})
export class CardComponent implements OnChanges {
  @Input() data!: IPost;
  date!: string;
  constructor(private router: Router, private authService: AuthService) {}

  isAuthorized: boolean = this.authService.state().isAuthorized;

  @Output() deletePost: EventEmitter<string> = new EventEmitter();

  handleDelete(id: string) {
    this.deletePost.emit(id);
  }

  ngOnChanges(changes: {
    [propertyName: string]: import('@angular/core').SimpleChange;
  }) {
    if (changes['data']) {
      this.date = convertTimeStampToDate(this.data.CreateAt);
    }
  }

  clickCard(id: string) {
    this.router.navigate([`post/${id}`]);
  }
}
