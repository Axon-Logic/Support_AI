import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { ApiServices } from './apis';
import { AuthService } from './auth.service';
import { Observable } from 'rxjs';
import { Storage } from '@ionic/storage';
// import { HTTP } from '@awesome-cordova-plugins/http/ngx';
import { HTTP } from '@ionic-native/http/ngx';
import { LoaderService } from './loader.service';
import { DataShare } from './data-share.service';

@Injectable({
  providedIn: 'root',
})
export class ApiService extends ApiServices {
  constructor(
    public _router: Router,
    private _http: HttpClient,
    private _storge: Storage,
    private authService: AuthService,
    private _httpCap: HTTP,
    private _loaderService: LoaderService,
    private _dataShare: DataShare
  ) {
    super(_router, _http, _storge, _httpCap, _loaderService, _dataShare);
  }

  getChats() {
    return this.getGetRequestLink('/history/getChats');
  }

  getMessages(id: string) {
    return this.getGetRequestLink('/history/getMessages/' + id);
  }

  readMessage(id: string) {
    return this.getPostRequestLink(
      '/firesocketsupportapi/temp/update/' + id,
      {}
    );
  }

  changeBotCondtion(data: { id: string; is_support: boolean }) {
    return this.getPostRequestLink('/history/switch', data);
  }

  getBotCondtion(id: string) {
    return this.getGetRequestLink(`/history/switch_position/${id}`);
  }

  sendMessage(message: {
    Sender: string;
    Message: string;
    MessageType: string;
    MessageId: string;
    ClientName: string;
    Caption: string;
  }) {
    return this.getPostRequestLink(
      '/firesocketsupportapi/temp/newMessage/' + message.Sender,
      message
    );
  }

  getMedia(uId: string) {
    return this.getGetRequestLink('/history/media/' + uId);
  }
}
