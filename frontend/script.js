document.getElementById('url-form').addEventListener('submit', async function (e) {
    e.preventDefault();

    const originalUrl = document.getElementById('original-url').value;
    const response = await fetch('http://localhost:8080/shorten-and-redirect', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ original_url: originalUrl }),
    });

    if (response.ok) {
        const data = await response.json();
        document.getElementById('short-url-result').innerText = `http://localhost:8080/${data.shortURL}`;
        document.getElementById('short-url').classList.remove('hidden');
    } else {
        alert('Error shortening URL!');
    }
});
