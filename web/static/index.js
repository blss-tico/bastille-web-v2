window.onload = () => {
  console.log("[info]: index window onload");

  sessionStorage.removeItem('id');
  sessionStorage.removeItem('jails');

  const uuidVal = uuidv7Generate();
  const uuidStr = Array.from(uuidVal).map((b) => b.toString(16).padStart(2, "0")).join("");
  const uuidFmt = uuidv7StringFormat(uuidStr);
  sessionStorage.setItem('id', uuidFmt);

  setTimeout(async () => { 
    await loadAllJailsIndex(); 
    document.dispatchEvent(notifEventIndex("Jails loaded"));
  }, 500);
}

function notifEventIndex(message) {
  console.log("[info]: notifEvent - ", message);
  return new CustomEvent('notification', { detail: { message }});
}

async function loadAllJailsIndex() {
  const formJson = JSON.stringify({ "action":"all", "options":"-j" });
  const jails = await sendHttpRequest("list", formJson);  
  if (jails.data.msg) { 
    sessionStorage.setItem('jails', jails.data.msg);
    console.log('[info]: jails loaded');
  }
}

function getCookie(name) { 
  const value = `; ${document.cookie}`; 
  const parts = value.split(`; ${name}=`);
  console.log("getCookie", value, parts);

  if (parts.length === 2) { return parts.pop().split(';').shift(); }
}

async function sendHttpRequest(url, bodyData) {
  console.log("[info]: sendHttpRequest ", url);

  const ipaddr = sessionStorage.getItem('ipaddr'); 
  const port = sessionStorage.getItem('port');

  if (!ipaddr) { return { code: 500, error: 'ip address node missing' }; }
  if (!port) { return { code: 500, error: 'port node missing' }; }
  
  console.log('sendHttpRequest: ', ipaddr, port);
  const response = await fetch(`http://${ipaddr}:${port}/` + url, {
    method: 'POST',
    credentials: "include",
    headers: {
      'Access-Control-Allow-Origin': '*',
      'accept': 'application/json',
      'Content-Type': 'application/json',
      'X-Request-ID': sessionStorage.getItem('id'),
      'Authorization': `Bearer ${sessionStorage.getItem('bw_actk')}`
    },
    body: bodyData,
    credentials: "include"
  });

  if (!response.ok) {
    const error = await response.text();
    console.log("[error]: ", error);
    document.dispatchEvent(notifEventIndex(`Error ${response.status} - ${error.message || error}`));
    return { code: response.status, error: error };
  }
  
  const responseData = await response.json();
  return { code: response.status, data: responseData };
}

function uuidv7Generate() {
  console.log("[info]: uuidv7Generate");

  const value = new Uint8Array(16);
  crypto.getRandomValues(value);

  const timestamp = BigInt(Date.now());
  value[0] = Number((timestamp >> 40n) & 0xffn);
  value[1] = Number((timestamp >> 32n) & 0xffn);
  value[2] = Number((timestamp >> 24n) & 0xffn);
  value[3] = Number((timestamp >> 16n) & 0xffn);
  value[4] = Number((timestamp >> 8n) & 0xffn);
  value[5] = Number(timestamp & 0xffn);
  value[6] = (value[6] & 0x0f) | 0x70;
  value[8] = (value[8] & 0x3f) | 0x80;

  return value;
}

function uuidv7StringFormat(uuid) {
  console.log("[info]: uuidv7StringFormat");

  const p1 = uuid.slice(0, 8);
  const p2 = uuid.slice(8, 12);
  const p3 = uuid.slice(12, 16);
  const p4 = uuid.slice(16, 20);
  const p5 = uuid.slice(20);
  const uuidFmt = p1 + "-" + p2 + "-" + p3 + "-" + p4 + "-" + p5;
  
  return uuidFmt;
}
