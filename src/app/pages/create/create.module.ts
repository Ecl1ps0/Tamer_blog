import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CreateComponent } from './create.component';
import { ReactiveFormsModule } from '@angular/forms';
import { ButtonComponent } from 'src/app/shared/ui/button/button.component';
@NgModule({
  declarations: [CreateComponent],
  imports: [CommonModule, ReactiveFormsModule, ButtonComponent],
})
export class CreateModule {}
