import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { AuthService } from '../shared/services/auth.service';
import { Preferences } from '@capacitor/preferences';
import * as uuid from 'uuid';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.page.html',
  styleUrls: ['./login.page.scss'],
})
export class LoginPage implements OnInit {
  loginForm: any = new FormGroup({
    username: new FormControl('', [
      Validators.required,
      Validators.minLength(4),
    ]),
    password: new FormControl('', [
      Validators.required,
      Validators.minLength(4),
    ]),
  });

  constructor(private authService: AuthService, private router: Router) {}

  ngOnInit() {}

  login() {
    if (
      this.loginForm.value.username.length < 4 ||
      this.loginForm.value.password.length < 4
    ) {
      return;
    }
    this.authService
      .login(this.loginForm.value.username, this.loginForm.value.password)
      .subscribe(
        (res) => {
          const myId = uuid.v4();
          (async () => {
            await Preferences.set({ key: 'token', value: myId });
          })();
          this.authService.token = myId;
          this.router.navigate(['/control/landing']);
        },
        (err) => {}
      );
  }
}
