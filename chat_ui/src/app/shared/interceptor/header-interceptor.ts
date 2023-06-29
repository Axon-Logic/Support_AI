import { LoaderService } from './../services/loader.service';
import {
  HttpEvent,
  HttpHandler,
  HttpHeaders,
  HttpInterceptor,
  HttpRequest,
  HttpResponse,
} from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { throwError as observableThrowError, catchError } from 'rxjs';
import { environment } from 'src/environments/environment';
import { AuthService } from '../services/auth.service';
import { map } from 'rxjs/operators';
import { Preferences } from '@capacitor/preferences';
import { DataShare } from '../services/data-share.service';

@Injectable()
export class HeaderInterceptor implements HttpInterceptor {
  constructor(
    private authService: AuthService,
    private router: Router,
    private loaderService: LoaderService,
    private dataShare: DataShare
  ) {}

  intercept(req: HttpRequest<any>, next: HttpHandler) {
    let token: any = this.authService.getToken();
    this.loaderService.loader.emit(true);

    let header = new HttpHeaders({
      'Content-Type': 'application/json',
      'Accept-Language': 'ru',
      type: 'BROWSER',
      Accept: 'application/json',
      'device-type': 'BROWSER',
    });

    if (
      token &&
      req.url != environment.host + '/api/v1/auth/login' &&
      req.url != environment.host + '/api/v1/auth/register'
    ) {
      header = header.append(
        'Authorization',
        'Bearer ' + this.authService.getToken()
      );
    }

    if (token && req.url.includes('/api/v1/orders/cancel')) {
      header = header.append('Host', 'http://corporate.railway.uz');
      header = header.append('Content-Length', '56');
    } else if (!token && req.url.includes('/api/v1/orders/cancel')) {
      (async () => {
        await Preferences.get({ key: 'token' }).then((res) => {
          token = res.value;
          header = header.append('Content-Length', '56');
        });
      })();
    }

    let authReq = req.clone({
      headers: header,
    });

    if (req.url.includes('/pdf')) {
      header = header.append('Accept', 'application/text');
      authReq = req.clone({
        responseType: 'text' as 'json',
        headers: header,
      });
    }

    return next
      .handle(authReq)
      .pipe(
        catchError((err) => {
          this.loaderService.loader.emit(false);
          if (err.status === 401 || err.status === 0) {
            let url = this.router.url;
            if (url != '/login') {
              Preferences.remove({ key: 'token' });
              Preferences.remove({ key: 'userId' });
              this.authService.notLoggedIn();
              // window.location.reload();
            }
          } else if (
            err.status === 500 &&
            err.error?.message == 'Invalid key exception!'
          ) {
            this.router.navigate(['/login']).then();
          }
          return observableThrowError(() => err);
        })
      )
      .pipe(
        map((event: HttpEvent<any>) => {
          if (event instanceof HttpResponse) {
            this.loaderService.loader.emit(false);
          }
          return event;
        })
      );
  }
}
