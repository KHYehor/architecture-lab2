import Device from './device';
import requests from '../common/http';

export default class Tablet extends Device {
  constructor(URL, name) {
    super(URL, name);
  };
  public getData() {
    return requests.getData(this.URL, this.name);
  }
  public sendData = (battery, currentVideo, date) => {
    return requests.sendData(this.URL, this.name, battery, currentVideo, date);
  }
}
