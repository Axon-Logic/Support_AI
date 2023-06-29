import { Component, OnInit, ViewChild, inject } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ApiService } from '../shared/services/api.service';
import { DataShare } from '../shared/services/data-share.service';
import { WebsocketService } from '../shared/services/websocket.service';
import { SimpleService } from '../shared/services/simple.service';
import { Filesystem, Directory } from '@capacitor/filesystem';
import * as uuid from 'uuid';
import { environment } from 'src/environments/environment';

@Component({
  selector: 'app-chats',
  templateUrl: './chats.page.html',
  styleUrls: ['./chats.page.scss'],
})
export class ChatsPage implements OnInit {
  public folder!: string;
  private activatedRoute = inject(ActivatedRoute);
  currentUser: any;
  botCondition = false;
  isOpenAdjustImage = false;
  idAdjustImage = '';
  environments = environment;
  @ViewChild('content') private content: any;
  messagesList: any = [];
  message = '';

  constructor(
    private router: Router,
    private apiService: ApiService,
    private dataShare: DataShare,
    public simpleService: SimpleService,
    private webSocketService: WebsocketService
  ) {}

  ngOnInit() {
    this.simpleService.changes.subscribe(() => {
      setTimeout(() => {
        this.content.scrollToBottom(1);
      }, 100);
    });
  }
  ionViewWillEnter() {
    this.currentUser = this.dataShare.selectedUser;
    if (!this.currentUser.client_id) {
      this.currentUser = JSON.parse(
        sessionStorage.getItem('currentUser') || ''
      );
    }
    this.folder = this.activatedRoute.snapshot.paramMap.get('id') as string;
    this.dataShare.selectedUserId = this.currentUser.client_id;
    if (!this.folder) {
      this.closeChat();
    }
    this.apiService.getBotCondtion(this.folder).subscribe((res: any) => {
      this.botCondition = res.switch.is_support;
    });
    this.apiService.getMessages(this.folder).subscribe(
      (res: any) => {
        this.runWebsocket();
        this.simpleService.messages = res;
        setTimeout(() => {
          this.content.scrollToBottom(1);
        }, 100);
      },
      (err: any) => {
        this.closeChat();
      }
    );
  }

  ionViewDidEnter() {
    this.content.scrollToBottom(1);
    this.apiService.readMessage(this.folder).subscribe(() => {});
  }

  runWebsocket() {
    this.simpleService.selectChat(this.folder);
  }

  closeChat() {
    this.router.navigate(['/control']);
  }

  changeBotCondition(e: any) {
    console.log(e.detail.checked);
    this.apiService
      .changeBotCondtion({
        id: this.folder,
        is_support: !e.detail.checked,
      })
      .subscribe(() => {});
  }

  sendMessage() {
    if (!this.message.length) {
      return;
    }
    const myId = uuid.v4();
    this.simpleService.pendingList.push(myId);
    this.simpleService.messages.push({
      message: this.message,
      MessageId: myId,
      status: 'pending',
      type: 'text',
      owner: true,
    });
    this.apiService
      .sendMessage({
        Sender: this.currentUser.client_id,
        Message: this.message,
        MessageType: 'text',
        ClientName: 'telegram',
        Caption: '',
        MessageId: myId,
      })
      .subscribe(
        (res) => {
          this.apiService.getBotCondtion(this.folder).subscribe((res: any) => {
            this.botCondition = res.switch.is_support;
          });
        },
        (err) => {}
      );
    this.message = '';
  }

  async saveAudio(data: any) {
    const result: any = await Filesystem.writeFile({
      path: `${data.client_id}-${data.message_id}` + '.wav',
      data: data.base64,
      directory: Directory.Data,
    });
    const pathe: any = result['uri'];
    return pathe;
  }

  adjustImage(uId: string) {
    this.idAdjustImage = uId;
    this.isOpenAdjustImage = true;
  }

  closeAdjustImage() {
    this.idAdjustImage = '';
    this.isOpenAdjustImage = false;
  }
}
