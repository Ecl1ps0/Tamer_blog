import { NgModule } from '@angular/core';
import {
  RouterModule,
  Routes,
  provideRouter,
  withComponentInputBinding,
} from '@angular/router';
import { HomeComponent } from './pages/home/home.component';
import { AboutComponent } from './pages/about/about.component';
import { PostComponent } from './pages/post/post.component';
import { CreateComponent } from './pages/create/create.component';

export const routes: Routes = [
  {
    path: '',
    component: HomeComponent,
  },
  {
    path: 'about',
    component: AboutComponent,
  },
  {
    path: 'post/:id',
    component: PostComponent,
  },
  {
    path: 'create',
    component: CreateComponent,
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { bindToComponentInputs: true })],
  providers: [provideRouter(routes, withComponentInputBinding())],
  exports: [RouterModule],
})
export class AppRoutingModule {}
