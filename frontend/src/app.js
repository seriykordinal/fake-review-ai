document.addEventListener("DOMContentLoaded", () => {
    const buttons = document.querySelectorAll('.mode-btn');
    const contents = document.querySelectorAll('.mode-content');

    buttons.forEach(btn => {
        btn.addEventListener('click', () => {
        
            buttons.forEach(b => b.classList.remove('active'));
            contents.forEach(c => c.classList.remove('active'));

            btn.classList.add('active');
            document.getElementById(btn.dataset.mode).classList.add('active');
        });
    });

    
    const button = document.querySelector(".primary-btn");
    const textarea = document.querySelector("textarea");
    const resultsContainer = document.getElementById("results");

    button.addEventListener("click", async () => {
    const text = textarea.value.trim();

    if (!text) {
        alert("Введите текст отзыва");
        return;
    }

    const response = await fetch("/reviews", {
        method: "POST",
        headers: {
        "Content-Type": "application/json",
        },
        body: JSON.stringify({ text }),
    });

    if (!response.ok) {
        alert("Ошибка сервера");
        return;
    }

    const data = await response.json();

    renderResult(text, data);
    textarea.value = "";
    });
    
    function renderResult(text, data) {
        const resultBlock = document.createElement("div");
        resultBlock.className = "result-block";
    
        const percent = Math.round(data.score * 100);
        
        resultBlock.innerHTML = `
            <div class="review-card">
                <div class="review-header">
                <div class="score">
                    <span>Вероятность фейка</span>
                    <strong>${percent}%</strong>
                </div>

                <div class="badge ${data.is_fake ? "fake" : "real"}">
                    <i class="fa-solid ${
                    data.is_fake ? "fa-triangle-exclamation" : "fa-circle-check"
                    }"></i>
                    ${data.is_fake ? "Подозрительный отзыв" : "Похоже на реальный отзыв"}
                </div>
                </div>

                <div class="review-text">
                ${escapeHtml(text)}
                </div>
            </div>
            `;

            
        resultsContainer.prepend(resultBlock);
    }
    
    // защита от XSS
    function escapeHtml(text) {
        const div = document.createElement("div");
        div.textContent = text;
        return div.innerHTML;
    }
    
});

