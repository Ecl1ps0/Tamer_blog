import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  template: `<header-component></header-component>
    <main class="md:container  mx-auto">
      <router-outlet></router-outlet>
    </main>
    <footer-component></footer-component>`,
})
export class AppComponent {}
