export default abstract class Device {
  protected name: string = '';
  constructor(name: string) {
    this.name = name;
  };
  abstract getData(count: number);
  abstract sendData(battery, currentVideo, date);
}

