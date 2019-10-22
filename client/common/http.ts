import fetch from 'node-fetch';

export default {
  sendData: (URL: string, name: string, battery: string, currentVideo: string, deviceTime: Date) => (
    fetch(`${URL}/sendData`, {
      method: 'POST',
      body: JSON.stringify({ name, battery, currentVideo, deviceTime }),
      headers: { 'Content-Type': 'application/json' }
    }).then((res: any) => res.json())
  ),
  getData: (URL: string, name: string) => (
    fetch(`${URL}/getData`, {
      method: 'POST',
      body: JSON.stringify({name}),
      headers: { 'Content-Type': 'application/json' }
    }).then((res: any) => res.json())
  )
};