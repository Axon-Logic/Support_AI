import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { ApiServices } from './apis';
import { JwtHelperService } from '@auth0/angular-jwt';
import { Storage } from '@ionic/storage';
import { Preferences } from '@capacitor/preferences';
// import { HTTP } from '@awesome-cordova-plugins/http/ngx';
import { HTTP } from '@ionic-native/http/ngx';
import { LoaderService } from './loader.service';
import { DataShare } from './data-share.service';

const helper = new JwtHelperService();

@Injectable({
  providedIn: 'root',
})
export class AuthService extends ApiServices {
  constructor(
    public _router: Router,
    private _http: HttpClient,
    private _storage: Storage,
    private _httpCap: HTTP,
    private _loaderService: LoaderService,
    private _dataShare: DataShare
  ) {
    super(_router, _http, _storage, _httpCap, _loaderService, _dataShare);

    async () => {
      await Preferences.get({ key: 'token' }).then((res) => {
        this.token = res.value;
      });
    };
  }

  login(username: string | any, password: string | any) {
    return this.getPostRequestLink('/auth/login', {
      username,
      password,
    });
  }

  getUserNameFromSession(): string | any {
    async () => {
      await this._storage.get('username').then((res: any) => {
        return res;
      });
    };
  }

  logOut(reload: boolean = false): void {
    this.token = null;
    this.decoded = null;
    Preferences.remove({ key: 'token' });
    Preferences.remove({ key: 'userId' });
    if (reload) {
      window.location.reload();
    }
  }
}
