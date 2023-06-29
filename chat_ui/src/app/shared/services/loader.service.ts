import { Injectable, EventEmitter } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class LoaderService {
  loader = new EventEmitter();
  timeOut: any = undefined;
  constructor() {
    this.loader.subscribe((res) => {
      if (res) {
        this.setTimer();
      } else {
        this.stopTimer();
      }
    });
  }

  setTimer() {
    this.timeOut = setTimeout(() => {
      this.loader.emit(false);
    }, 20000);
  }
  stopTimer() {
    clearTimeout(this.timeOut);
  }
}
