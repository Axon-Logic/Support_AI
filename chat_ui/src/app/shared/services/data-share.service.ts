import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class DataShare {
  constructor() {}
  platform: any = '';
  orderId: any = '';
  selectedUser: any = {};
  selectedUserId = '';
  showSideMenu = false;
}
