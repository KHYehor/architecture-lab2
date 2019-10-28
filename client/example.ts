'use strict';

import Tablet from './device/tablet';

const URL = 'http://localhost:8080';
const Tablet1 = new Tablet(URL, 'class1-tablet2');

Tablet1.getData()
  .then(console.log, console.error);

Tablet1.sendData("89%", "currentVideo", new Date())
  .then(console.log, console.error);

Tablet1.sendData("49%", "currentVideo", new Date())
  .then(console.log, console.error);

Tablet1.sendData("39%", "currentVideo", new Date())
  .then(console.log, console.error);

const Tablet2 = new Tablet(URL, 'class1-tablet1');

Tablet2.sendData("89%", "currentVideo", new Date())
  .then(console.log, console.error);

Tablet2.sendData("49%", "currentVideo", new Date())
  .then(console.log, console.error);

Tablet2.sendData("39%", "currentVideo", new Date())
  .then(console.log, console.error);

Tablet2.getData()
  .then(console.log, console.error);

