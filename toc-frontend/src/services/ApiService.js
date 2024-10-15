import axios from 'axios';

// Utility functions
function buildRestAPIEndPoint(path) {
  let host
  const version = 1

  if (process.env.REACT_APP_REST_API_HOST && process.env.REACT_APP_REST_API_HOST !== ""){
    // If environment variable is set, override other assumed settings.
    host = process.env.REACT_APP_REST_API_HOST
  } else if(process.env.NODE_ENV === "production") {
    // Assumes same host and server serving frontend app.
    host = "" 
  } else if(process.env.NODE_ENV === "development") {
    // Assumes same host but different development servers serving frontend 
    // app and REST API.
    host = "localhost:8888"  
  }

  if (path.charAt(0) === '/'){
    path = path.slice(1)
  }

  return `http://${host}/v${version}/${path}`
}

function curl(method, url, data) {
  let request
  
  let m = String(method).toUpperCase()
  let u = buildRestAPIEndPoint(url)

  if(m === "GET"){
    request = axios.get(u)
  } else if(m === "POST"){
    if(data != null){
      request = axios.post(u, data)
    } else {
      request = axios.post(u)
    }
  } else {
    console.error(`Unknown method: ${method}`)
    return
  }
  
  // Now execute the request
  request
    .then(onResponse)
    .catch(onError);
}

function onResponse(res){
  console.log(res);
  console.log(res.data);
}

function onError(error){
  // Error
  if (error.response) {
    // The request was made and the server responded with a status code
    // that falls out of the range of 2xx
    console.error(error.response.data);
    console.error(error.response.status);
    console.error(error.response.headers);
  } else if (error.request) {
      // The request was made but no response was received
      // `error.request` is an instance of XMLHttpRequest in the 
      // browser and an instance of
      // http.ClientRequest in node.js
      console.error(error.request);
  } else {
      // Something happened in setting up the request that triggered an Error
      console.error('Error', error.message);
  }
  console.error(error.config);
}

function newBackup() {
  curl("POST", "/backups", null)
}

function getBackups() {
  curl("GET", "/backups", null)
}

function getBackupByID(id) {
  curl("GET", `/backups/${id}`, null)
}

function newRestoration(backup) {
  curl("POST", "/restorations", backup)
}

function getRestorations() {
  curl("GET", "/restorations", null)
}

function getRestorationByID(id) {
  curl("GET", `/restorations/${id}`, null)
}

function newSynchronization() {
  curl("POST", "/synchronizations", null)
}

function getQuotas() {
  curl("GET", "/quotas", null)
}

function newLogs() {
  curl("POST", "/logs", null)
}

function newArchive() {
  curl("POST", "/archives", null)
}

export {
  newBackup,
  getBackups,
  getBackupByID,
  newRestoration,
  getRestorations,
  getRestorationByID,
  newSynchronization,
  getQuotas,
  newLogs,
  newArchive
};
