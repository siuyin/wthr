
async function getPos() {
  const pos = await new Promise((success)=>{
    navigator.geolocation.getCurrentPosition(success)
  })
  const lat = pos.coords.latitude
  const lng = pos.coords.longitude
  const f = await fetch(`nfc?lat=${lat}&lng=${lng}`)
  const s = document.getElementById("localsec")
  s.innerHTML = await f.text()
  console.log(`${lat},${lng}`)
}

document.addEventListener("DOMContentLoaded", getPos)
