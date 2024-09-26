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

function newBackup() {
  console.log("POST " + buildRestAPIEndPoint("/backups"))
}

function getBackups() {
  console.log("GET " + buildRestAPIEndPoint("/backups"))
}

function getBackupByID(id) {
  console.log("GET " + buildRestAPIEndPoint(`/backups/${id}`))
}

function newRestoration(backup) {
  console.log("POST " + buildRestAPIEndPoint("/restorations"))
  console.log(JSON.stringify(backup))
}

function getRestorations() {
  console.log("GET " + buildRestAPIEndPoint("/restorations"))
}

function getRestorationByID(id) {
  console.log("GET " + buildRestAPIEndPoint(`/restorations/${id}`))
}

function newSynchronization() {
  console.log("POST " + buildRestAPIEndPoint("/synchronizations"))
}

function getQuotas() {
  console.log("GET " + buildRestAPIEndPoint("/quotas"))
}

function newLogs() {
  console.log("POST " + buildRestAPIEndPoint("/logs"))
}

function newArchive() {
  console.log("POST " + buildRestAPIEndPoint("/archives"))
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
