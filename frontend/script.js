document.addEventListener("DOMContentLoaded", function () {
    document.getElementById("url-form").addEventListener("submit", async function (event) {
        event.preventDefault();

        const originalUrl = document.getElementById("original-url").value;
        const shortUrlDiv = document.getElementById("short-url");
        const shortUrlResult = document.getElementById("short-url-result");
        const errorMessage = document.getElementById("error-message");
        const copyBtn = document.getElementById("copy-btn");

        shortUrlDiv.classList.add("hidden");
        errorMessage.classList.add("hidden");

        try {
            const response = await fetch("/shorten", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ original_url: originalUrl }),
            });

            if (!response.ok) throw new Error("Failed to shorten URL.");

            const data = await response.json();
            shortUrlResult.innerHTML = `<a href="${data.shortened_url}" target="_blank">${data.shortened_url}</a>`;
            shortUrlDiv.classList.remove("hidden");

            // Ensure only one event listener is bound for copying
            copyBtn.onclick = function () {
                navigator.clipboard.writeText(data.shortened_url).then(() => {
                    alert("Copied to clipboard!");
                }).catch(err => {
                    console.error("Clipboard copy failed:", err);
                    alert("Copy failed, please try manually.");
                });
            };

        } catch (error) {
            errorMessage.textContent = error.message;
            errorMessage.classList.remove("hidden");
        }
    });
});
