import Tablet from './device/tablet';
const URL = 'http://localhost:8080';
const Tablet1 = new Tablet(URL, 'class1-tablet2');
Tablet1.getData()
  .then(console.log, console.error)
Tablet1.sendData("89%", "currentVudeo", "2019-10-21T09:39:22.599909Z")
  .then(console.log, console.error)
