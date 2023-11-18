import { RouterTestingModule } from '@angular/router/testing';
import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ButtonComponent } from '../ui/button/button.component';

@Component({
  selector: 'header-component',
  standalone: true,
  imports: [CommonModule, ButtonComponent, RouterTestingModule],
  templateUrl: './header.component.html',
})
export class HeaderComponent {
  isOpen: boolean = false;

  handleOpen() {
    this.isOpen = !this.isOpen;
  }
}
