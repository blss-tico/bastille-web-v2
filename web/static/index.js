// bastille-web-v2 global interface logic
function notifEventIndex(message) {
  console.log("[info]: notifEvent - ", message);
  return new CustomEvent('notification', { detail: { message }});
}

function getCookie(name) { 
  const value = `; ${document.cookie}`; 
  const parts = value.split(`; ${name}=`);
  console.log("[info]: getCookie", value, parts);

  if (parts.length === 2) { return parts.pop().split(';').shift(); }
}

async function sendHttpRequest(url, bodyData) {
  console.log("[info]: sendHttpRequest ", url);

  const ipaddr = sessionStorage.getItem('ipaddr'); 
  const port = sessionStorage.getItem('port');

  if (!ipaddr) { return { code: 500, error: 'ip address node missing' }; }
  if (!port) { return { code: 500, error: 'port node missing' }; }
  
  let response;
  //response = await fetch(`http://${ipaddr}:${port}/` + url, {
  response = await fetch("/" + url, {
    method: 'POST',
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

  if (response.status === 401) { 
    const refreshRes = await fetch(`/refresh`, { 
      method: "POST",
      body: JSON.stringify({ bw_rftk: sessionStorage.getItem('bw_rftk') }),
      credentials: "include" 
    }); 
    
    if (refreshRes.ok) { 
      const data = await refreshRes.json();
      sessionStorage.setItem("bw_actk", data.bw_actk); 
       
      response = await fetch(`/` + url, {
        method: 'POST',
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
    }
  }

  if (!response.ok) {
    const error = await response.text();
    console.log("[error]: ", error);
    document.dispatchEvent(notifEventIndex(`Error ${response.status} - ${error.message || error}`));
    return { code: response.status, error: error };
  }
  
  const responseData = await response.json();
  return { code: response.status, data: responseData };
}
