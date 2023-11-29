import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'button-shared',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './button.component.html',
})
export class ButtonComponent {
  @Input('primary') primaryButton: boolean = false;
  @Input('danger') dangerButton: boolean = false;
  @Input('type') type: string = 'text';
}
