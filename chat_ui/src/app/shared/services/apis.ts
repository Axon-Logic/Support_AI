import { isPlatform } from '@ionic/angular';
import { Preferences } from '@capacitor/preferences';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { JwtHelperService } from '@auth0/angular-jwt';
import { Storage } from '@ionic/storage';
import { HTTP } from '@ionic-native/http/ngx';
import { from, map } from 'rxjs';
import { LoaderService } from './loader.service';
import { DataShare } from './data-share.service';
import { environment } from 'src/environments/environment';

const helper = new JwtHelperService();

export class ApiServices {
  checking = false;
  token: string | any = '';
  decoded: object | any = null;
  userName = '';
  userId = '';
  loginDialog: any;
  isLoggedIn = false;
  constructor(
    public router: Router,
    private http: HttpClient,
    public storage: Storage,
    private httpCap: HTTP,
    private loaderService: LoaderService,
    private dataShare: DataShare
  ) {
    this.setTokenToAuth();
  }

  checkPlatform() {
    if (isPlatform('capacitor')) {
      return 'capacitor';
    } else {
      return 'web';
    }
  }

  getGetRequestLink(url: string): any {
    if (isPlatform('capacitor')) {
      if (!this.token) {
        this.setTokenToAuth();
      }
      this.loaderService.loader.emit(true);
      this.httpCap.setDataSerializer('json');
      const headers: any = {
        'Content-Type': 'application/json',
        'Accept-Language': 'ru',
        Accept: 'application/json',
        'X-FORWARDED-FOR': '192.168.1.39',
        type: this.dataShare.platform,
        'device-type': 'MOBILE',
        Authorization: `Bearer ${this.token}`,
      };
      // if (url.includes('pdf')) {
      //   headers['Accept'] = 'application/octet-stream';
      // }
      console.log(url);

      return from(
        this.httpCap.get(`${environment.url}${url}`, {}, headers)
      ).pipe(
        map((res) => {
          console.log(res);
          this.loaderService.loader.emit(false);
          if (url.includes('pdf')) return res?.data;
          else {
            if (res?.data) {
              return JSON.parse(res?.data);
            }
          }
        })
      );
    } else {
      return this.http.get(url);
    }
  }
  getPutRequestLink(url: string): any {
    if (isPlatform('capacitor')) {
      if (!this.token) {
        this.setTokenToAuth();
      }
      this.loaderService.loader.emit(true);
      this.httpCap.setDataSerializer('json');
      const headers: any = {
        'Content-Type': 'application/json',
        'Accept-Language': 'ru',
        Accept: 'application/json',
        'X-FORWARDED-FOR': '192.168.1.39',
        type: this.dataShare.platform,
        'device-type': 'MOBILE',
        Authorization: `Bearer ${this.token}`,
      };

      return from(
        this.httpCap.put(`${environment.url}${url}`, {}, headers)
      ).pipe(
        map((res) => {
          console.log(res);
          this.loaderService.loader.emit(false);
          if (url.includes('pdf')) return res?.data;
          else {
            if (res?.data) {
              return JSON.parse(res?.data);
            }
          }
        })
      );
    } else {
      return this.http.put(url, {});
    }
  }

  getPostRequestLink(url: string, data: any) {
    if (isPlatform('capacitor')) {
      if (!this.token) {
        this.setTokenToAuth();
      }
      this.loaderService.loader.emit(true);
      this.httpCap.setDataSerializer('json');

      const headers: any = {
        'Content-Type': 'application/json',
        'Accept-Language': 'ru',
        Accept: 'application/json',
        type: this.dataShare.platform,
        'X-FORWARDED-FOR': '192.168.1.39',
        'device-type': 'MOBILE',
      };

      if (this.isLoggedIn) {
        headers['Authorization'] = `Bearer ${this.token}`;
      }
      return from(
        this.httpCap.post(`${environment.url}${url}`, data, headers)
      ).pipe(
        map((res) => {
          console.log(res);
          this.loaderService.loader.emit(false);
          if (url.includes('pdf')) return res?.data;
          else {
            if (res?.data) {
              return JSON.parse(res?.data);
            }
          }
        })
      );
    } else {
      return this.http.post(url, data);
    }
  }

  getDeleteRequestLink(url: string) {
    if (isPlatform('capacitor')) {
      if (!this.token) {
        this.setTokenToAuth();
      }
      this.loaderService.loader.emit(true);
      this.httpCap.setDataSerializer('json');

      const headers = {
        'Content-Type': 'application/json',
        'Accept-Language': 'ru',
        Accept: 'application/json',
        type: this.dataShare.platform,
        'X-FORWARDED-FOR': '192.168.1.39',
        'device-type': 'MOBILE',
        Authorization: `Bearer ${this.token}`,
      };

      return from(
        this.httpCap.delete(`${environment.url}${url}`, {}, headers)
      ).pipe(
        map((res) => {
          console.log(res);
          this.loaderService.loader.emit(false);
          if (url.includes('pdf')) return res?.data;
          else return JSON.parse(res?.data);
        })
      );
    } else {
      return this.http.delete(`${url}`);
    }
  }

  setToken(token: string, check: boolean = false): void {
    this.token = token;
    try {
      let decode = helper.decodeToken(this.token);
      if (decode) {
        this.userId = decode.id;
        this.userName = decode.sub;
      } else {
        this.userId = '';
        this.userName = '';
      }
    } catch (e) {
      console.error(e);
    }
    (async () => {
      await Preferences.set({ key: 'token', value: token }).then((res) => {
        if (token?.length) {
          this.isLoggedIn = true;
        }
      });
    })();
    (async () => {
      await Preferences.set({ key: 'userName', value: this.userName });
    })();
    (async () => {
      await Preferences.set({ key: 'userId', value: this.userId });
    })();
  }

  notLoggedIn() {
    this.token = '';
    this.isLoggedIn = false;
    this.userId = '';
    this.userName = '';
    Preferences.remove({ key: 'token' });
    Preferences.remove({ key: 'userId' });
    Preferences.remove({ key: 'userName' });
  }

  getToken() {
    if (this.token) {
      return this.token;
    } else {
      return '';
    }
  }

  async setTokenToAuth() {
    await Preferences.get({ key: 'token' }).then((res: any) => {
      if (res.value && res.value !== '[object Object]') {
        this.isLoggedIn = true;
        this.token = res.value;
        console.log(res);
      }
    });
  }

  async setUserName() {
    await Preferences.get({ key: 'userName' }).then((res) => {
      if (res.value) {
        this.userName = res.value;
      }
    });
  }

  async setUserID() {
    await Preferences.get({ key: 'userId' }).then((res) => {
      if (res.value) {
        this.userId = res.value;
      }
    });
  }
}
