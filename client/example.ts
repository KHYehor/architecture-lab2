import Tablet from './device/tablet';

const Tablet1 = new Tablet('class1-tablet2');
Tablet1.getData()
  .then(res => res.json())
  .then(console.log)
  .catch(console.error)
