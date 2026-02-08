const BASE_URL = "http://localhost:8080";
const token = localStorage.getItem("token");

if (!token) {
  window.location.href = "login.html";
}

// Fetch user URLs on load
window.onload = fetchMyUrls;

async function fetchMyUrls() {
  const res = await fetch(`${BASE_URL}/my-urls`, {
    headers: {
      "Authorization": "Bearer " + token
    }
  });

  const data = await res.json();
  const list = document.getElementById("urlList");
  list.innerHTML = "";

  data.forEach(item => {
    const li = document.createElement("li");
    li.innerHTML = `
      <a href="${item.short_url}" target="_blank">${item.short_url}</a>
      (${item.clicks} clicks)
    `;
    list.appendChild(li);
  });
}

async function shortenUrl() {
  const url = document.getElementById("urlInput").value;

  const res = await fetch(`${BASE_URL}/shorten`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Authorization": "Bearer " + token
    },
    body: JSON.stringify({ original_url: url })
  });

  await res.json();
  fetchMyUrls();
}
