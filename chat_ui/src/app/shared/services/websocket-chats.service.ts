import { EventEmitter, Injectable } from '@angular/core';
import { webSocket } from 'rxjs/webSocket';
import { environment } from 'src/environments/environment';
import { AuthService } from './auth.service';
import { DataShare } from './data-share.service';

@Injectable({
  providedIn: 'root',
})
export class WebsocketChatsService {
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
      url: environment.wsurl + 'getChats/',
      deserializer: ({ data }) => {
        if (data) return JSON.parse(data);
      },
    });
    this.subject.subscribe(
      async (msg: any) => {
        console.log(msg);
        this.getDataFromWS.emit(msg);
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
    if (this.error) {
      this.connect();
    }
    return this.subject.next();
  }

  public async unsubscribe() {
    this.subject.next();

    this.subscribed = false;
  }
}
