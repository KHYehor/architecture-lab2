export default abstract class Device {
  protected name: string = '';
  protected URL: string = '';
  constructor(URL: string, name: string) {
    this.URL = URL;
    this.name = name;
  };
  abstract getData();
  abstract sendData(battery, currentVideo, date);
}

