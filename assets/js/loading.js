function updateProgress() {
    fetch('/progress')
        .then(response => response.json())
        .then(data => {
            const progressBar = document.getElementById('progress-bar');
            progressBar.value = data.progress;
            if (data.progress >= 100) {
                window.location.href = '/home';
            } else {
                setTimeout(updateProgress, 500);
            }
        });
}
window.onload = updateProgress;