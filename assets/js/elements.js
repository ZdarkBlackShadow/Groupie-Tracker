document.addEventListener('DOMContentLoaded', function() {
    const carousel = document.querySelector('.carousel');
    const items = document.querySelectorAll('.element-item');
    const centerTitle = document.querySelector('.center-title h1');
    
    const totalElements = items.length;
    const angleStep = 360 / totalElements;
    const radius = 225;

    function initializeCarousel() {
        items.forEach((item, index) => {
            const angle = angleStep * index;
            const radian = (angle - 90) * Math.PI / 180;
            
            const x = Math.cos(radian) * radius;
            const y = Math.sin(radian) * radius;
            
            // Positionner l'élément
            item.style.transform = `translate(-50%, -50%)`;
            item.style.left = `calc(50% + ${x}px)`;
            item.style.top = `calc(50% + ${y}px)`;

            // Contre-rotation pour maintenir l'élément droit
            const content = item.querySelector('.element-item-content');
            if (content) {
                // Animation de contre-rotation
                content.style.animation = `counter-rotate 20s infinite linear`;
            }
        });
    }

    initializeCarousel();

    // Effet de hover pour le titre
    items.forEach(item => {
        item.addEventListener('mouseenter', () => {
            const elementName = item.querySelector('img').alt;
            centerTitle.textContent = elementName;
            centerTitle.style.opacity = '1';
        });

        item.addEventListener('mouseleave', () => {
            centerTitle.textContent = 'Elements';
            centerTitle.style.opacity = '0.8';
        });
    });
}); 