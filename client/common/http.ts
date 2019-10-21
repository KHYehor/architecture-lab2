import fetch from 'node-fetch';

export default {
  sendData: (URL, name, battery, currentVideo, deviceTime) => (
    fetch(`${URL}/sendData`, {
      method: 'POST',
      body: JSON.stringify({ name, battery, currentVideo, deviceTime }),
      headers: { 'Content-Type': 'application/json' }
    }).then(res => res.json())
  ),
  getData: (URL, name) => (
    fetch(`${URL}/getData`, {
      method: 'POST',
      body: JSON.stringify({name}),
      headers: { 'Content-Type': 'application/json' }
    }).then(res => res.json())
  )
};