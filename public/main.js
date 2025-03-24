function getPos() {
  navigator.geolocation.getCurrentPosition((pos) =>{
    localForecasts(pos.coords.latitude,pos.coords.longitude)
  })
}

async function localForecasts(lat,lng) {
  const f = await fetch(`nfc?lat=${lat}&lng=${lng}`)
  const txt = await f.text()
  const s = document.getElementById("localsec")
  s.innerHTML = txt
  console.log(`lat: ${lat}, lng: ${lng}`)
}


document.addEventListener("DOMContentLoaded", getPos)
