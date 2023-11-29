import { Component, EventEmitter, Output } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'modal-ui',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './modal.component.html',
})
export class ModalComponent {
  @Output() setIsOpen: EventEmitter<any> = new EventEmitter();

  handleClose() {
    this.setIsOpen.emit(false);
  }
}
