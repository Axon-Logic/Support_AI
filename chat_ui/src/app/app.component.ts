import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { MenuController, NavController } from '@ionic/angular';
import { ApiService } from './shared/services/api.service';
import { DataShare } from './shared/services/data-share.service';
import { SimpleService } from './shared/services/simple.service';
import { Preferences } from '@capacitor/preferences';
import { AuthService } from './shared/services/auth.service';
@Component({
  selector: 'app-root',
  templateUrl: 'app.component.html',
  styleUrls: ['app.component.scss'],
})
export class AppComponent implements OnInit {
  isVisible = true;
  public chats: any = [];
  constructor(
    private menuCtrl: MenuController,
    private router: Router,
    private authService: AuthService,
    public dataShare: DataShare,
    public simpleService: SimpleService
  ) {}
  ngOnInit(): void {
    if (window.innerWidth < 992) {
      this.dataShare.showSideMenu = false;
    } else {
      this.dataShare.showSideMenu = true;
    }
    console.log(window.innerWidth);

    console.log(this.dataShare.showSideMenu);

    addEventListener('resize', (event) => {
      if (window.innerWidth < 992) {
        this.dataShare.showSideMenu = false;
      } else {
        this.dataShare.showSideMenu = true;
      }
    });
    this.router.events.subscribe((res: any) => {
      const url = this.router.url;
      if (url === '/control/login') {
        this.isVisible = false;
      } else this.isVisible = true;
    });
    this.simpleService.connectWebsocket();
    this.simpleService.connectWebsocketChats();
    this.simpleService.getChats();

    async () => {
      await Preferences.get({ key: 'token' }).then((res) => {
        this.authService.token = res.value;
      });
    };
  }

  openChat(chat: any) {
    this.menuCtrl.close('menu');
    this.simpleService.messages = [];
    sessionStorage.setItem('currentUser', JSON.stringify(chat));
    this.router.navigate(['/control/chats/', chat.client_id]);
    this.dataShare.selectedUser = chat;
  }
}
