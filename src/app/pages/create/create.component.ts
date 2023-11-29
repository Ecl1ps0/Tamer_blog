import { Component } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { BaseService } from 'src/app/services/base.service';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
})
export class CreateComponent {
  imageArrayBuffer: string | ArrayBuffer | null = null;

  postFormData = new FormGroup({
    imageContent: new FormControl<ArrayBuffer | null>(null),
    textContent: new FormControl<string>(''),
    title: new FormControl<string>(''),
  });

  onFileSelected(event: any): void {
    const file = event.target.files[0];
    this.postFormData.patchValue({
      imageContent: event.target.files[0],
    });
    if (file) {
      const reader = new FileReader();
      reader.onload = (e) => {
        this.imageArrayBuffer = e.target?.result!;
      };

      reader.readAsDataURL(file);
    }
  }

  constructor(
    private readonly base: BaseService,
    private readonly router: Router
  ) {}

  addPost() {
    const postData = {
      ImageContent: this.postFormData.get('imageContent')?.value as ArrayBuffer,
      TextContent: this.postFormData.get('textContent')?.value as string,
      Title: this.postFormData.get('title')?.value as string,
    };
    console.log(postData);
    this.base.addPost(postData).subscribe({
      next: (value) => {
        console.log(value);
        this.router.navigate(['/']);
      },
      error: (err) => {
        console.log(err);
      },
    });
  }
}
