import Device from './device';
import requests from '../common/http';

export default class Tablet extends Device {
  constructor(name) {
    super(name);
  };
  public getData() {
    return requests.getData(this.name);
  }
  public sendData = (battery, currentVideo, date) => {
    return requests.sendData(this.name, battery, currentVideo, date);
  }
}
