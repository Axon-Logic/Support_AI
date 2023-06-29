import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MenuController } from '@ionic/angular';
import { ApiService } from '../shared/services/api.service';
import { DataShare } from '../shared/services/data-share.service';
import { SimpleService } from '../shared/services/simple.service';

@Component({
  selector: 'app-landing',
  templateUrl: './landing.page.html',
  styleUrls: ['./landing.page.scss'],
})
export class LandingPage implements OnInit {
  isVisible = true;

  constructor(
    private menuCtrl: MenuController,
    private router: Router,
    private apiService: ApiService,
    public simpleService: SimpleService,
    public dataShare: DataShare
  ) {}

  ngOnInit(): void {}

  ionViewWillEnter() {}

  openChat(chat: any) {
    this.menuCtrl.close('menu');
    this.simpleService.messages = [];
    sessionStorage.setItem('currentUser', JSON.stringify(chat));
    this.router.navigate(['/control/chats/', chat.client_id]);
    this.dataShare.selectedUser = chat;
  }
}
