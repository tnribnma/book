const API_BASE = 'http://localhost:8080';

function getToken() {
  return localStorage.getItem('token');
}

function authHeader() {
  return {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${getToken()}`
  };
}

function showMessage(elementId, message, isError = false) {
  const el = document.getElementById(elementId);
  if (el) {
    el.textContent = message;
    el.className = `message ${isError ? 'error' : 'success'}`;
  }
}

async function apiRequest(url, method = 'GET', body = null) {
  const options = {
    method,
    headers: authHeader()
  };

  if (body) {
    options.body = JSON.stringify(body);
  }

  const res = await fetch(API_BASE + url, options);
  
  if (res.status === 401) {
    alert("Session expired. Please login again.");
    localStorage.removeItem('token');
    window.location.href = 'login.html';
    return null;
  }

  return res;
}