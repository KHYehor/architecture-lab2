import fetch from 'node-fetch';
const URL = 'http://localhost:8080';

export default {
  sendData: (name, battery, currentVideo, deviceTime) => (
    fetch(`${URL}/sendData`, {
      method: 'POST',
      body: JSON.stringify({ name, battery, currentVideo, deviceTime }),
      headers: { 'Content-Type': 'application/json' }
    })
  ),
  getData: name => (
    fetch(`${URL}/getData`, {
      method: 'POST',
      body: JSON.stringify({name}),
      headers: { 'Content-Type': 'application/json' }
    })
  )
};