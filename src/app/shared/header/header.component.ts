import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ButtonComponent } from '../ui/button/button.component';
import { Router, RouterLinkWithHref } from '@angular/router';
import { ModalComponent } from '../ui/modal/modal.component';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'header-component',
  standalone: true,
  imports: [
    CommonModule,
    ButtonComponent,
    RouterLinkWithHref,
    ModalComponent,
    ReactiveFormsModule,
  ],
  templateUrl: './header.component.html',
})
export class HeaderComponent {
  isOpen: boolean = false;
  isOpenModal: boolean = false;

  constructor(
    private readonly authService: AuthService,
    private readonly router: Router
  ) {}

  isAuthorized: boolean = this.authService.state().isAuthorized;

  password = new FormControl('');

  handleAuthorization() {
    this.authService.signIn(this.password.value!);
    console.log(this.authService.state());
    this.isOpenModal = false;
    this.isAuthorized = this.authService.state().isAuthorized;
  }

  handleOpenModal() {
    this.isOpenModal = !this.isOpenModal;
  }

  handleOpen() {
    this.isOpen = !this.isOpen;
  }

  handlePostCreate() {
    this.router.navigate(['/create']);
  }
}
