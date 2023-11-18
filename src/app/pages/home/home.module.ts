import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomeComponent } from './home.component';
import { CardComponent } from 'src/app/shared/ui/card/card.component';

@NgModule({
  declarations: [HomeComponent],
  imports: [CommonModule, CardComponent],
  exports: [CardComponent],
})
export class HomeModule {}
