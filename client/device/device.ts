export default abstract class Device {
  protected name: string = '';
  protected URL: string = '';
  constructor(URL: string, name: string) {
    this.URL = URL;
    this.name = name;
  };
  abstract getData(): Promise<any>;
  abstract sendData(battery: string, currentVideo: string, date: Date): Promise<any>;
}

