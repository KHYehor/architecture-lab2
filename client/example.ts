'use strict';

import Tablet from './device/tablet';

const URL = 'http://localhost:8080';
const Tablet1 = new Tablet(URL, 'class1-tablet2');

Tablet1.getData()
  .then(console.log, console.error);
Tablet1.sendData("89%", "currentVideo", new Date())
  .then(console.log, console.error);

