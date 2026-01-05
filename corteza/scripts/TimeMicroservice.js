export default {
  label: "Idő és Feldolgozás",
  triggers({ on }) {
    return on('manual').for('compose:module').uiProp('app', 'compose');
  },

  async exec({ module }, { Compose }) {
    const timeRes = await fetch('http://time-microservice-service/time');
    const { current_time } = await timeRes.json();

    const procRes = await fetch('http://time-microservice-service/process', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ message: "teszt adat" })
    });
    const { processed } = await procRes.json();

    await Compose.saveRecord(await Compose.makeRecord({
      time: current_time,
      result: processed
    }, module));
  }
};