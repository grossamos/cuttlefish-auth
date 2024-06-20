function login() {
  const username = document.getElementById('username').value
  const password = document.getElementById('password').value

  const credentials = btoa(`${username}:${password}`);
    
  fetch('/login', {
    method: 'POST',
    headers: {
        'Authorization': `Basic ${credentials}`
    },
  })
  .then(response => {
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    return response.json();
  })
  .then(data => {
    const token = data.token;
    localStorage.setItem('authToken', token);
    window.location.href = '/';
  })
  .catch(error => {
    console.error('Error:', error);
  });
}
