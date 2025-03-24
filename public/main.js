async function localForecasts() {
  const f = await fetch("/nfc")
  const txt = await f.text()
  const s = document.getElementById("localsec")
  s.innerHTML = txt
}

document.addEventListener("DOMContentLoaded", localForecasts)

