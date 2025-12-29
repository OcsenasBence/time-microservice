export default {
  label: "Idő Lekérése Microservice-ből",
  description: "Meghívja a Go szervert és elmenti az időt a TimeRecord modulba",

  triggers({ on }) {
    return on('manual')
      .for('compose:module')
      .uiProp('app', 'compose')
      .where('module', 'timerecord')
      .where('namespace', 'timemicroservice')
  },

  async exec(args, { Compose }) {
    const endpointUrl = 'http://time-microservice-service/time';

    console.log(`[TimeWorkflow] Lekérés indítása ide: ${endpointUrl}`);

    let responseData;

    try {
      const response = await fetch(endpointUrl, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' }
      });

      if (!response.ok) {
        throw new Error(`Hiba a szerver válaszában: ${response.status}`);
      }

      responseData = await response.json();
      console.log(`[TimeWorkflow] Kapott idő: ${responseData.current_time}`);

    } catch (e) {
      console.error('[TimeWorkflow] Hiba a lekérésnél:', e);
      throw e;
    }

    try {
        const newRecord = await Compose.makeRecord({
            time: responseData.current_time
        }, args.module);

        await Compose.saveRecord(newRecord);
        console.log('[TimeWorkflow] Rekord sikeresen mentve!');
        return;

    } catch (e) {
        console.error('[TimeWorkflow] Hiba a mentésnél:', e);
        throw e;
    }
  }
};