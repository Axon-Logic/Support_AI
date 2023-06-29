import { EventEmitter, Injectable } from '@angular/core';
import { WebsocketService } from './websocket.service';
import { WebsocketChatsService } from './websocket-chats.service';
import { ApiService } from './api.service';

@Injectable({
  providedIn: 'root',
})
export class SimpleService {
  chats: any = {};
  chatsValues: any = [];
  messages: any = [];
  changes: EventEmitter<any>;
  changesChats: EventEmitter<any>;
  pendingList: any = [];

  constructor(
    private webSocketService: WebsocketService,
    private apiService: ApiService,
    private webSocketChatsService: WebsocketChatsService
  ) {
    this.changes = new EventEmitter();
    this.changesChats = new EventEmitter();
  }

  uniq(a: any[]) {
    return Array.from(new Set(a));
  }

  connectWebsocket() {
    this.webSocketService.subscribe();
    this.webSocketService.getDataFromWS.subscribe((res: any) => {
      this.messages.push(res);
      this.deletePendingMessage();
      this.changes.emit();
    });
  }

  getChats() {
    this.apiService.getChats().subscribe(
      (res: any) => {
        this.chats = res;
        res.map((chat: any) => {
          this.chats[chat.client_id] = chat;
          this.chatsValues.push(chat.client_id);
        });
        this.chatsValues = this.uniq(this.chatsValues);
      },
      (err: any) => {}
    );
  }

  connectWebsocketChats() {
    this.webSocketChatsService.subscribe();
    this.webSocketChatsService.getDataFromWS.subscribe((res: any) => {
      const foundValue = this.chatsValues.find(
        (chat: any) => chat === res.client_id
      );
      if (!foundValue) {
        this.chatsValues.push(res.client_id);
      }
      this.chats[res.client_id] = res;
      this.chatsValues = this.uniq(this.chatsValues);
      this.changesChats.emit();
    });
  }

  selectChat(chatId: string) {
    this.messages = [];
    this.webSocketService.selectChat(chatId);
  }

  deletePendingMessage() {
    this.messages = this.messages.filter(
      (message: any) =>
        !(
          message.status == 'pending' &&
          this.pendingList.includes(message.MessageId)
        )
    );
  }
}
