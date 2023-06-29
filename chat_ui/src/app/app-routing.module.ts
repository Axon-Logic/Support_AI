import { NgModule } from '@angular/core';
import { PreloadAllModules, RouterModule, Routes } from '@angular/router';
import { CabinetGuard } from './shared/guards/cabinet.guard';

const routes: Routes = [
  {
    path: '',
    redirectTo: 'control/login',
    pathMatch: 'full',
  },
  {
    path: 'control',
    redirectTo: 'control/landing',
  },
  {
    path: 'control/chats/:id',
    loadChildren: () =>
      import('./chats/chats.module').then((m) => m.ChatsPageModule),
    canActivate: [CabinetGuard],
  },
  {
    path: 'control/settings',
    loadChildren: () =>
      import('./settings/settings.module').then((m) => m.SettingsPageModule),
    canActivate: [CabinetGuard],
  },
  {
    path: 'control/analytics',
    loadChildren: () =>
      import('./analytics/analytics.module').then((m) => m.AnalyticsPageModule),
    canActivate: [CabinetGuard],
  },
  {
    path: 'control/login',
    loadChildren: () =>
      import('./login/login.module').then((m) => m.LoginPageModule),
  },
  {
    path: 'control/landing',
    loadChildren: () =>
      import('./landing/landing.module').then((m) => m.LandingPageModule),
    canActivate: [CabinetGuard],
  },
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes, { preloadingStrategy: PreloadAllModules }),
  ],
  exports: [RouterModule],
})
export class AppRoutingModule {}
