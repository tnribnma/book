const API_BASE = 'http://localhost:8080';

async function register() {
  const email = document.getElementById('regEmail').value;
  const password = document.getElementById('regPassword').value;
  const messageEl = document.getElementById('regMessage');

  try {
    const res = await fetch(API_BASE + '/users/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password })
    });

    const data = await res.json();

    if (res.ok) {
      messageEl.innerHTML = `<span class="text-green-600">Registration successful! <a href="login.html">Login now</a></span>`;
    } else {
      messageEl.innerHTML = `<span class="text-red-600">${data.error || 'Registration failed'}</span>`;
    }
  } catch (err) {
    messageEl.innerHTML = `<span class="text-red-600">Server error</span>`;
  }
}

async function login() {
  const email = document.getElementById('email').value;
  const password = document.getElementById('password').value;
  const messageEl = document.getElementById('message');

  try {
    const res = await fetch(API_BASE + '/auth/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password })
    });

    const data = await res.json();

    if (res.ok) {
      localStorage.setItem('token', data.token);
      window.location.href = 'books.html';
    } else {
      messageEl.innerHTML = `<span class="text-red-600">${data.error || 'Login failed'}</span>`;
    }
  } catch (err) {
    messageEl.innerHTML = `<span class="text-red-600">Server error</span>`;
  }
}