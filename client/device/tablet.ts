import Device from './device';
import requests from '../common/http';

export default class Tablet extends Device {
  constructor(name) {
    super(name);
  };
  public getData(count) {
    requests.getData(count);
  }
  public sendData = (battery, currentVideo, date) => {
    requests.sendData(battery, currentVideo, date);
  }
}
