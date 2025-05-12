<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Message Board</title>
  <style>
    body {
      font-family: sans-serif;
      max-width: 600px;
      margin: 40px auto;
      padding: 0 1rem;
    }
    h1 {
      text-align: center;
    }
    form {
      display: flex;
      gap: 0.5rem;
      margin-bottom: 1rem;
    }
    input[type="text"] {
      flex: 1;
      padding: 0.5rem;
      font-size: 1rem;
    }
    button {
      padding: 0.5rem 1rem;
      font-size: 1rem;
      cursor: pointer;
    }
    ul {
      list-style: none;
      padding: 0;
    }
    li {
      background: #f1f1f1;
      padding: 0.5rem;
      margin-bottom: 0.5rem;
      border-radius: 4px;
    }
  </style>
</head>
<body>
  <h1>Message Board</h1>
  <form id="message-form">
    <input type="text" id="message-input" placeholder="Type your message..." required />
    <button type="submit">Send</button>
  </form>
  <ul id="message-list"></ul>

  <script>
    const form = document.getElementById('message-form');
    const input = document.getElementById('message-input');
    const list = document.getElementById('message-list');

    async function fetchMessages() {
      const res = await fetch("${api_url}");
      const data = await res.json();
      list.innerHTML = '';
      for (let i = data.length - 1; i >= 0; i--) {
        const li = document.createElement('li');
        li.textContent = data[i].text;
        list.appendChild(li);
      }
    }

    form.addEventListener('submit', async (e) => {
      e.preventDefault();
      const text = input.value.trim();
      if (!text) return;
      await fetch("${api_url}", {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ text }),
      });
      input.value = '';
      fetchMessages();
    });

    fetchMessages();
  </script>
</body>
</html>
