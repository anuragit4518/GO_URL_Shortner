
const BASE_URL = "http://localhost:8080";

async function signup() {
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
  
    const res = await fetch("http://localhost:8080/signup", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password })
    });
  
    document.getElementById("msg").innerText =
      res.ok ? "Signup successful. Go to login." : "Signup failed";
}


async function login() {
  const email = document.getElementById("email").value;
  const password = document.getElementById("password").value;

  const res = await fetch(`${BASE_URL}/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password })
  });

  const data = await res.json();

  if (res.ok) {
    localStorage.setItem("token", data.token);
    window.location.href = "dashboard.html";
  } else {
    document.getElementById("msg").innerText = "Login failed";
  }

}

async function sendOTP() {
    const email = document.getElementById("email").value;
  
    const res = await fetch(`${BASE_URL}/forgot-password`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email })
    });
  
    document.getElementById("msg").innerText =
      res.ok ? "OTP sent to email" : "Failed to send OTP";
  
    if (res.ok) {
      setTimeout(() => {
        window.location.href = "reset-password.html";
      }, 1500);
    }
  }
  
  async function resetPassword() {
    const email = document.getElementById("email").value;
    const otp = document.getElementById("otp").value;
    const newPassword = document.getElementById("newPassword").value;
  
    const res = await fetch(`${BASE_URL}/reset-password`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        email,
        otp,
        new_password: newPassword
      })
    });
  
    document.getElementById("msg").innerText =
      res.ok ? "Password reset successful. Go to login." : "Reset failed";
  
    if (res.ok) {
      setTimeout(() => {
        window.location.href = "login.html";
      }, 2000);
    }
  }
  

