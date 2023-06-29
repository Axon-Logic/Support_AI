import { EventEmitter, Injectable } from '@angular/core';
import { webSocket } from 'rxjs/webSocket';
import { environment } from 'src/environments/environment';
import { AuthService } from './auth.service';
import { DataShare } from './data-share.service';

@Injectable({
  providedIn: 'root',
})
export class WebsocketService {
  private subject: any;
  public getDataFromWS: EventEmitter<any>;

  private error: boolean = false;
  subscribed = false;
  connectCount = 0;

  constructor(private dataShare: DataShare, private authService: AuthService) {
    this.getDataFromWS = new EventEmitter();
    this.connect();
  }

  public connect() {
    this.subject = webSocket({
      url: environment.wsurl + 'getMessages/',
      deserializer: ({ data }) => {
        if (data) return JSON.parse(data);
      },
    });
    this.subject.subscribe(
      async (msg: any) => {
        let id = this.dataShare.selectedUserId;
        if (msg) {
          if (msg.client_id == id) {
            this.getDataFromWS.emit(msg);
          }
        }
      },
      (err: any) => {
        this.error = true;
      },
      () => {}
    );
  }

  selectChat(chatId: string) {
    if (this.error) {
      this.connect();
    }
    if (chatId) {
      return this.subject.next({
        client_id: chatId,
      });
    }
  }

  public subscribe() {
    let id = this.dataShare.selectedUserId;
    let userId = +id;
    if (this.error) {
      this.connect();
    }
    if (userId) {
      return this.subject.next(userId);
    }
  }

  public async unsubscribe() {
    this.subject.next();
    this.subscribed = false;
  }
}
