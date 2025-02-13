document.addEventListener('DOMContentLoaded', function() {
    const tabButtons = document.querySelectorAll('.tab-button');
    const materialLists = document.querySelectorAll('.materials-list');

    // Set default active tab
    if (tabButtons.length > 0) {
        tabButtons[0].classList.add('active');
        materialLists[0].classList.add('active');
    }

    tabButtons.forEach(button => {
        button.addEventListener('click', () => {
            // Remove active class from all buttons and lists
            tabButtons.forEach(btn => btn.classList.remove('active'));
            materialLists.forEach(list => list.classList.remove('active'));

            // Add active class to clicked button and corresponding list
            button.classList.add('active');
            const level = button.dataset.level;
            document.querySelector(`.materials-list[data-level="${level}"]`).classList.add('active');
        });
    });
}); 