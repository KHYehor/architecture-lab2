import Device from './device';
import requests from '../common/http';

export default class Tablet extends Device {
  constructor(URL: string, name: string) {
    super(URL, name);
  };
  public getData() {
    return requests.getData(this.URL, this.name);
  }
  public sendData = (battery: string, currentVideo: string, date: Date) => {
    return requests.sendData(this.URL, this.name, battery, currentVideo, date);
  }
}
